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
		return runAnalyzers()
	},
}

func runAnalyzers() error {
	kaniHome := os.Getenv("KANI_HOME")
	analyzersDir := filepath.Join(kaniHome, "analyses")
	entries, err := ioutil.ReadDir(analyzersDir)
	if err != nil {
		return err
	}
	return runAnalyzersImpl(analyzersDir, entries)
}

func execFile(fp string) int {
	command := exec.Command(fp)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	command.Run()
	return command.ProcessState.ExitCode()
}

func execFileIfNeeded(wg *sync.WaitGroup, ch chan<- int, fp string, entry fs.FileInfo) {
	if entry.Mode()&1 != 1 {
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- execFile(fp)
	}()
}

func runAnalyzersImpl(dir string, entries []fs.FileInfo) error {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	for _, entry := range entries {
		fp := filepath.Join(dir, entry.Name())
		execFileIfNeeded(wg, ch, fp, entry)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	statuses := receive(ch)
	printGuidesIfNeeded(statuses, dir)
	return nil
}

func printGuides(resourceDir string) {
	guideFile := filepath.Join(resourceDir, "commit_guide.txt")
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

func printGuidesIfNeeded(statuses []int, analysesDir string) {
	printFlag := false
	for _, status := range statuses {
		if status == 1 {
			printFlag = true
			break
		}
	}
	if printFlag {
		printGuides(filepath.Join(filepath.Dir(analysesDir), "resources"))
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
