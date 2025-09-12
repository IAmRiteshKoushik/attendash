package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version of the tool",
	Long: `Displays the current version of the Attendash CLI tool.

Use this command to quickly check which version of Attendash you are running.
This can help diagnose issues or confirm you have the latest features and fixes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Attendash v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
