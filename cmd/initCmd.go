package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tamada/kani/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "git kani init",
	Long:  "initialize kani",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0] == "-" {
			printZshInitializer()
		} else {
			runInitializeKani(initializeKani)
		}
	},
}
var deinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "git kani deinit",
	Long:  "deinitialize kani",
	Run: func(cmd *cobra.Command, args []string) {
		runInitializeKani(deinitializeKani)
	},
}

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

func printZshInitializer() {
	fmt.Println(`
function __kani_chpwd_hook() {
  /usr/local/opt/kani/scripts/chpwd_hook.sh
}
function __kani_periodic_hook() {
  /usr/local/opt/kani/scripts/periodic_hook.sh
}
function preexec_test() {
  /usr/local/opt/kani/scripts/preexec_hook.sh
}

autoload -Uz add-zsh-hook
PERIOD=60
add-zsh-hook chpwd    __kani_chpwd_hook
add-zsh-hook periodic __kani_periodic_hook
add-zsh-hook preexec  __kani_preexec_hook`)
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(deinitCmd)
}
