package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tamada/kani/utils"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "git kani enable",
	Long:  "enable kani",
	Run: func(cmd *cobra.Command, args []string) {
		disableKani(false)
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "git kani disable",
	Long:  "disable kani",
	Run: func(cmd *cobra.Command, args []string) {
		disableKani(true)
	},
}

func findProjectKaniDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	projectDir, err2 := utils.FindProjectDir(pwd)
	if err2 != nil {
		return "", err2
	}
	return filepath.Join(projectDir, ".kani"), nil
}

func disableKani(flag bool) {
	kaniDir, err := findProjectKaniDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	disableFile := filepath.Join(kaniDir, "disable")
	if flag {
		if err := utils.Touch(disableFile); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("kani: disable")
	} else {
		os.Remove(disableFile)
		fmt.Println("kani: enable")
	}
}

func init() {
	RootCmd.AddCommand(enableCmd)
	RootCmd.AddCommand(disableCmd)
}
