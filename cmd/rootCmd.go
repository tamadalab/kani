package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helpMessage() string {
	return `git kani init       initialize kani for the current project.
git kani deinit     deinitialize kani of the project.
git kani enable     enable kani.
git kani disable    disable kani.`
}

var RootCmd = &cobra.Command{
	Use:   "kani",
	Short: "kani",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(helpMessage())
	},
}

func init() {
	cobra.OnInitialize()
}

func Execute() error {
	return RootCmd.Execute()
}
