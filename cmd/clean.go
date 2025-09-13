package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up all the seed used in development mode",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: cleanFunc,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func cleanFunc(cmd *cobra.Command, args []string) error {
	if _, err := Orm.Delete(dbName); err != nil {
		return err
	}

	log.Info("DB removed")
	return nil
}
