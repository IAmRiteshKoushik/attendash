package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database for testing during development",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: seedFunc,
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seedFunc(cmd *cobra.Command, args []string) error {
	fmt.Println("seed called")
	return nil
}
