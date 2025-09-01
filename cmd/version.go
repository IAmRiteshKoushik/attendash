package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version of the tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Attendash v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
