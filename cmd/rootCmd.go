package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helpMessage() string {
	return `kani init       initialize kani for the current project.
kani deinit     deinitialize kani of the project.
kani enable     enable kani.
kani disable    disable kani.`
}

// RootCmd shows root command for kani.
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

// Execute executes the command.
func Execute() error {
	return RootCmd.Execute()
}
