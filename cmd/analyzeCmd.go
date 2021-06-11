package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/tamada/kani/utils"
)

const formatter string = "2006-01-02 15:04:05"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

type options struct {
	dbfile string
}

var opts = &options{}

type historyAnalyzer struct {
	h *history
	m map[string]*analyzer
	a map[string]bool
}

type history struct {
	Name       string
	Experiment string
	ID         int
	Datetime   time.Time
	Command    string
	Status     int
	Branch     string
	Revision   string
}

type analyzer struct {
	Name         string
	Experiment   string
	ID           int
	AnalyzerName string
	PerformCode  int
}

var analyzeCmd = &cobra.Command{
	Use:    "analyze",
	Short:  "kani analyze [OPTIONS]",
	Long:   "analyze data of the database",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return analyzeDatabase(cmd, args)
	},
}

func readAllAnalyzers(db *sql.DB, projectName, userName string) ([]*analyzer, error) {
	search := `SELECT * FROM analyzers ORDER BY history_id, analyzer_name`
	rows, err := db.Query(search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	analyzers := []*analyzer{}
	for rows.Next() {
		a := &analyzer{Experiment: projectName, Name: userName}
		rows.Scan(&a.ID, &a.AnalyzerName, &a.PerformCode)
		analyzers = append(analyzers, a)
	}
	return analyzers, nil
}

func readAllHistories(db *sql.DB, projectName, userName string) ([]*history, error) {
	search := `SELECT id, datetime, command, status_code, branch, revision FROM histories ORDER BY id`
	rows, err := db.Query(search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	histories := []*history{}
	for rows.Next() {
		h := createHistory(rows, projectName, userName)
		histories = append(histories, h)
	}
	return histories, nil
}

func createHistory(rows *sql.Rows, projectName, userName string) *history {
	h := &history{Experiment: projectName, Name: userName}
	var timeString string
	rows.Scan(&h.ID, &timeString, &h.Command, &h.Status, &h.Branch, &h.Revision)
	dt, _ := time.Parse(formatter, timeString)
	h.Datetime = dt.In(jst)
	return h
}

func extractProjectNameFromPath(path string) string {
	base := filepath.Base(path)
	if base != "kani.sqlite" {
		return strings.TrimSuffix(base, ".sqlite")
	}
	dir := filepath.Dir(path)
	if filepath.Base(dir) != ".kani" {
		return "unknown"
	}
	name := filepath.Base(filepath.Dir(dir))
	if name == "." {
		pwd, _ := os.Getwd()
		p, _ := utils.FindProjectDir(pwd)
		return filepath.Base(p)
	}
	return name
}

func findUserName() string {
	u, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return u.Username
}

func readAllData(path string) ([]*history, []*analyzer, error) {
	db, err := openDatabase(path)
	if err != nil {
		return nil, nil, fmt.Errorf("openDatabase failed: %w", err)
	}
	project := extractProjectNameFromPath(path)
	userName := findUserName()
	histories, err := readAllHistories(db, project, userName)
	if err != nil {
		return nil, nil, fmt.Errorf("readAllHistories failed: %w", err)
	}
	analyzers, err := readAllAnalyzers(db, project, userName)
	if err != nil {
		return nil, nil, fmt.Errorf("readAllAnalyzers failed: %w", err)
	}
	return histories, analyzers, nil
}

func findExitCode(ha *historyAnalyzer, name string) int {
	a, ok := ha.m[name]
	if ok {
		return a.PerformCode
	}
	return 0
}

func findNextAction(ha, next *historyAnalyzer, name string) bool {
	if next == nil {
		return false
	}
	nextCommand := strings.TrimSpace(next.h.Command)
	switch name {
	case "recommends_add_by_editted_lines.py", "recommends_add_by_updating_program.py":
		return strings.HasPrefix(nextCommand, "git add")
	case "recommends_commit_by_all_files_staged.sh":
		return strings.HasPrefix(nextCommand, "git commit")
	case "recommends_push_after_commit.sh":
		return strings.HasPrefix(nextCommand, "git push")
	}
	return false
}

func escapeCommand(command string) string {
	return strings.ReplaceAll(command, "\"", "\"\"")
}

func printAll(ha, next *historyAnalyzer) {
	fmt.Printf("%s,%s,%d,%s,\"%s\",%d,%s,%s", ha.h.Name, ha.h.Experiment, ha.h.ID, ha.h.Datetime.Format(formatter), escapeCommand(ha.h.Command), ha.h.Status, ha.h.Branch, ha.h.Revision)
	names := []string{"recommends_add_by_editted_lines.py", "recommends_add_by_updating_program.py", "recommends_commit_by_all_files_staged.sh", "recommends_push_after_commit.sh"}
	for _, name := range names {
		value := findExitCode(ha, name)
		nextAction := findNextAction(ha, next, name)
		fmt.Printf(",%d,%v", value, nextAction)
		ha.a[name] = nextAction
	}
	fmt.Println()
}

func toHistoryAnalyzer(h *history, analyzers []*analyzer) *historyAnalyzer {
	ha := &historyAnalyzer{h: h, m: map[string]*analyzer{}, a: map[string]bool{}}
	for _, a := range analyzers {
		if h.ID == a.ID && h.Name == a.Name && h.Experiment == a.Experiment {
			ha.m[filepath.Base(a.AnalyzerName)] = a
		}
	}
	return ha
}

func toHistoryAnalyzers(histories []*history, analyzers []*analyzer) []*historyAnalyzer {
	has := []*historyAnalyzer{}
	for _, history := range histories {
		ha := toHistoryAnalyzer(history, analyzers)
		has = append(has, ha)
	}
	return has
}

func printStats(has []*historyAnalyzer) {
	names := []string{"recommends_add_by_editted_lines.py", "recommends_add_by_updating_program.py", "recommends_commit_by_all_files_staged.sh", "recommends_push_after_commit.sh"}
	// recommendation rates
	printRecommendationRates(has, names)
	printRecommendationFollowingRates(has, names)
}

func printRecommendationFollowingRates(has []*historyAnalyzer, names []string) {
	countAll := 0
	counts := map[string]int{}
	denominator := 0
	denominators := map[string]int{}
	for _, ha := range has {
		anyOk := false
		anyOkDenominator := false
		for _, name := range names {
			_, ok1 := ha.m[name]
			ok2 := ha.a[name]
			if ok1 && ok2 {
				counts[name] = counts[name] + 1
				anyOk = true
			}
			if ok1 {
				denominators[name] = denominators[name] + 1
				anyOkDenominator = true
			}
		}
		if anyOk {
			countAll++
		}
		if anyOkDenominator {
			denominator++
		}
	}
	fmt.Printf("recommendation following rates,%f", float64(countAll)/float64(len(has)))
	for _, name := range names {
		fmt.Printf(",%f", float64(counts[name])/float64(len(has)))
	}
	fmt.Println()
	fmt.Printf("recommendation following rates (recommended),%f", float64(countAll)/float64(denominator))
	for _, name := range names {
		fmt.Printf(",%f", float64(counts[name])/float64(denominators[name]))
	}
	fmt.Println()
}

func printRecommendationRates(has []*historyAnalyzer, names []string) {
	countAll := 0
	counts := map[string]int{}
	for _, ha := range has {
		anyOk := false
		for _, name := range names {
			_, ok := ha.m[name]
			if ok {
				counts[name] = counts[name] + 1
				anyOk = true
			}
		}
		if anyOk {
			countAll++
		}
	}
	fmt.Printf("recommendation rates,%f", float64(countAll)/float64(len(has)))
	for _, name := range names {
		fmt.Printf(",%f", float64(counts[name])/float64(len(has)))
	}
	fmt.Println()
}

func analyzeDatabase(cmd *cobra.Command, args []string) error {
	path := findDatabasePath(opts)
	fmt.Printf("findDatabasePath(): %s\n", path)
	histories, analyzers, err := readAllData(path)
	if err != nil {
		return fmt.Errorf("readAllData(%s) failed: %w", path, err)
	}
	has := toHistoryAnalyzers(histories, analyzers)
	for i, ha := range has {
		var next *historyAnalyzer = nil
		if (i + 1) < len(has) {
			next = has[i+1]
		}
		printAll(ha, next)
	}
	printStats(has)
	return nil
}

func findDatabasePath(opts *options) string {
	if opts.dbfile == "" {
		return findDBPath()
	}
	return opts.dbfile
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
	flags := analyzeCmd.Flags()
	flags.StringVarP(&opts.dbfile, "file", "f", "", "specifies the sqlite database file")
}
