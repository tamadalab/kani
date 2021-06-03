package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/tamada/kani/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "kani init",
	Long:  "initialize kani",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0] == "-" {
			printShellInitializer()
		} else {
			runInitializeKani(initializeKani)
		}
	},
}
var deinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "kani deinit",
	Long:  "deinitialize kani",
	Run: func(cmd *cobra.Command, args []string) {
		runInitializeKani(deinitializeKani)
	},
}

var deinitOpts = struct {
	deleteAll bool
}{deleteAll: false}

func runInitializeKani(initializer func(dir string) error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	projectDir, err2 := utils.FindProjectDir(pwd)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	if err := initializer(projectDir); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s: set as the target project of kani", projectDir)
}

func createKaniDirectory(projectDir string) error {
	kaniDir := filepath.Join(projectDir, ".kani")
	if utils.ExistDir(kaniDir) {
		return nil
	}
	if err := os.Mkdir(kaniDir, 0755); err != nil {
		return err
	}
	files := []string{
		"analyses/dummy.sh",
	}
	kaniHome, err := utils.KaniHome()
	if err != nil {
		return err
	}
	for _, from := range files {
		utils.CopyFile(kaniHome, from, kaniDir)
	}
	return nil
}

func findLine(line, fileName string) bool {
	reader, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		entry := scanner.Text()
		if entry == line {
			return true
		}
	}
	return false
}

func addProjectName(name, fileName string) error {
	writer, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer writer.Close()
	writer.Write([]byte(name))
	writer.Write([]byte("\r\n"))
	return nil
}

func addProjectNameToProjectsList(projectDir string) error {
	projectsList := filepath.Join(utils.FindConfDir(), "projects")
	if !findLine(projectDir, projectsList) {
		return addProjectName(projectDir, projectsList)
	}
	return nil
}

func initializeKani(projectDir string) error {
	if err := createKaniDirectory(projectDir); err != nil {
		return err
	}
	if err := addProjectNameToProjectsList(projectDir); err != nil {
		return err
	}
	return nil
}

func removeFromProjectList(projectDir, projectsList string) error {
	wholeProjects, err := ioutil.ReadFile(projectsList)
	if err != nil {
		return err
	}
	writer, err := os.OpenFile(projectsList, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer writer.Close()

	projects := strings.Split(string(wholeProjects), "\r\n")
	for _, project := range projects {
		writer.Write([]byte(project))
		writer.Write([]byte("\r\n"))
	}
	return nil
}

func deinitializeKani(projectDir string) error {
	projectsList := filepath.Join(utils.FindConfDir(), "projects")
	if findLine(projectDir, projectsList) {
		removeFromProjectList(projectDir, projectsList)
	}
	return nil
}

func findKaniHome() (string, error) {
	entries := []string{
		os.Getenv("KANI_HOME"),
		"/usr/local/opt/kani",
		"$(HOME)/go/src/github.com/tamadalab/kani",
		"$(HOME)/go/src/github.com/tamada/kani",
	}
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		kaniHome := strings.ReplaceAll(entry, "$(HOME)", homeDir)
		stat, err := os.Stat(kaniHome)
		if err == nil && stat.Mode().IsDir() {
			return kaniHome, nil
		}
	}
	return "", errors.New("KANI_HOME did not found")
}

func printShellInitializer() {
	shell := os.Getenv("SHELL")
	kaniHome, err := findKaniHome()
	if err != nil {
		fmt.Printf("kani: %s", err.Error())
	}
	if strings.HasSuffix(shell, "zsh") {
		printZshInitializer(kaniHome)
	} else if strings.HasSuffix(shell, "bash") || strings.HasSuffix(shell, "/sh") {
		printBashInitializer(kaniHome)
	}
}

func messageInstallingBashPreexec() string {
	return `kani on bash requires rcaloras/bash-preexec (https://github.com/rcaloras/bash-preexec)
Please run 'curl https://raw.githubusercontent.com/rcaloras/bash-preexec/master/bash-preexec.sh -o ~/.bash-preexec.sh'
or 'brew install bash-preexec'.`
}

func installedBashPreexec() (bool, string) {
	locations := []string{ "~/.bash-preexec.sh", "/usr/local/etc/profile.d/bash-preexec.sh", }
	for _, location := range locations {
		if isRegularFile(location) {
			return true, location
		}
	}
	return false, ""
}

func isRegularFile(location string) bool {
	if strings.HasPrefix(location, "~") {
		home, err := homedir.Dir()
		if err != nil {
			return false
		}
		return isRegularFile(strings.ReplaceAll(location, "~", home))
	}
	stat, err := os.Stat(location)
	return err == nil && stat.Mode().IsRegular()
}

func bashPreexec(location string) string {
	if location == "/usr/local/etc/profile.d/bash-preexec.sh" {
		return ""
	}
	return fmt.Sprintf("source %s", location)
}

func printBashInitializer(kaniHome string) {
	ok, location := installedBashPreexec()
	if !ok {
		fmt.Printf(`echo "%s"
`, messageInstallingBashPreexec())
		return
	}
	fmt.Printf(`%s
__kani_preexec_hook() {
  if [[ ! -e %s ]]; then
    echo "%s"
    return
  else
    %s/scripts/preexec_hook.sh "$1"
  fi
}
__kani_precmd_hook() {
  statusCode=($?)
  if [[ ! -e %s ]]; then
    echo "%s"
    return
  else
    %s/scripts/precmd_hook.sh $statusCode
  fi
}
preexec_functions+=(__kani_preexec_hook)
precmd_functions+=(__kani_precmd_hook)
`, bashPreexec(location), location, messageInstallingBashPreexec(), kaniHome, location, messageInstallingBashPreexec(), kaniHome)
}

func printZshInitializer(kaniHome string) {
	fmt.Printf(`function __kani_preexec_hook() {
  %s/scripts/preexec_hook.sh "$1"
}
function __kani_precmd_hook() {
  %s/scripts/precmd_hook.sh $? # gives the status code
}

autoload -Uz add-zsh-hook
PERIOD=60
add-zsh-hook preexec  __kani_preexec_hook
add-zsh-hook precmd   __kani_precmd_hook
`, kaniHome, kaniHome)
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(deinitCmd)
	deinitCmd.Flags().BoolVar(&deinitOpts.deleteAll, "--delete-all", false, "deletes .kani directory on the project root")
}
