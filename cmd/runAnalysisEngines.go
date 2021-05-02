package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var analysesEngineCmd = &cobra.Command{
	Use:    "run-analyzers",
	Short:  "kani run-analyzers",
	Long:   "run analyzers of kani",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		updateEnvs(args)
		return runAnalyzers()
	},
}

func updateEnvs(args []string) {
	os.Setenv("KANI_CURRENT_BRANCH", args[0])
	os.Setenv("KANI_CURRENT_REVISION", args[1])
}

func collectAnalyzers(path string) ([]string, error) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	return pickExecutables(path, entries)
}

func pickExecutables(dir string, entries []fs.FileInfo) ([]string, error) {
	paths := []string{}
	for _, entry := range entries {
		if entry.Mode()&1 == 1 {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}
	return paths, nil
}

func analyzersPaths() ([]string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return []string{}, err
	}
	return []string{
		filepath.Join(os.Getenv("KANI_HOME"), "analyses"),
		filepath.Join(config, "kani", "analyses"),
	}, nil
}

func runAnalyzers() error {
	analyzersDirs, err := analyzersPaths()
	if err != nil {
		return err
	}
	analyzers := []string{}
	for _, path := range analyzersDirs {
		paths, err := collectAnalyzers(path)
		if err != nil {
			return err
		}
		analyzers = append(analyzers, paths...)
	}
	return runAnalyzersImpl(analyzers)
}

func execFile(fp string) int {
	command := exec.Command(fp)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	command.Run()
	return command.ProcessState.ExitCode()
}

func runAnalyzersImpl(analyzers []string) error {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	for _, analyzer := range analyzers {
		wg.Add(1)
		analyzer := analyzer
		go func() {
			defer wg.Done()
			status := execFile(analyzer)
			ch <- status
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	statuses := receive(ch)
	printGuidesIfNeeded(statuses)
	return nil
}

func findResourceDir() string {
	return filepath.Join(os.Getenv("KANI_HOME"), "resources")
}

func printGuides() {
	guideFile := filepath.Join(findResourceDir(), "commit_guide.txt")
	file, err := os.Open(guideFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	defer file.Close()
	data := make([]byte, 1024)
	for {
		len, err := file.Read(data)
		if len == 0 {
			break
		}
		if err == io.EOF {
			fmt.Print(string(data))
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		fmt.Print(string(data))
	}
}

func printGuidesIfNeeded(statuses []int) {
	printFlag := false
	for _, status := range statuses {
		if status == 1 {
			printFlag = true
			break
		}
	}
	if printFlag {
		printGuides()
	}
}

func receive(ch <-chan int) []int {
	statuses := []int{}
	for status := range ch {
		statuses = append(statuses, status)
	}
	return statuses
}

func init() {
	RootCmd.AddCommand(analysesEngineCmd)
}
