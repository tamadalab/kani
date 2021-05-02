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
	"github.com/tamada/kani/utils"
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
	if !utils.ExistDir(path) {
		return []string{}, nil
	}
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
	home, err := os.UserHomeDir()
	if err != nil {
		return []string{}, err
	}
	return []string{
		filepath.Join(os.Getenv("KANI_HOME"), "analyses"),
		filepath.Join(home, ".config", "kani", "analyses"),
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

type analyzerResult struct {
	name   string
	status int
}

func runAnalyzersImpl(analyzers []string) error {
	wg := new(sync.WaitGroup)
	ch := make(chan *analyzerResult)
	for _, analyzer := range analyzers {
		wg.Add(1)
		analyzer := analyzer
		go func() {
			defer wg.Done()
			status := execFile(analyzer)
			ch <- &analyzerResult{name: analyzer, status: status}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	statuses := receive(ch)
	err := storeResult(statuses)
	if err == nil {
		printGuidesIfNeeded(statuses)
	}
	return err
}

func storeResult(results []*analyzerResult) error {
	db, err := openDatabase(findDBPath())
	if err != nil {
		return err
	}
	id, err := findLastID(db)
	if err != nil {
		return err
	}
	tx, _ := db.Begin()
	defer tx.Rollback() // rollback will be ignored after commit.
	stmt, err := tx.Prepare("INSERT INTO analyzers VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, result := range results {
		_, err := stmt.Exec(id, result.name, result.status)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
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
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func printGuidesIfNeeded(statuses []*analyzerResult) {
	printFlag := false
	for _, ar := range statuses {
		if ar.status != 0 {
			printFlag = true
			break
		}
	}
	if printFlag {
		printGuides()
	}
}

func receive(ch <-chan *analyzerResult) []*analyzerResult {
	statuses := []*analyzerResult{}
	for status := range ch {
		statuses = append(statuses, status)
	}
	return statuses
}

func init() {
	RootCmd.AddCommand(analysesEngineCmd)
}
