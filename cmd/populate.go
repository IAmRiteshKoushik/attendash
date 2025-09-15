package cmd

import (
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var populateCmd = &cobra.Command{
	Use:   "csv",
	Short: "Clean up all the seed used in development mode",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: populateFunc,
}

func init() {
	rootCmd.AddCommand(populateCmd)
}

func populateFunc(cmd *cobra.Command, args []string) error {
	return nil
}
