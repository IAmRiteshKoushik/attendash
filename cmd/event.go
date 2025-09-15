package cmd

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/IAmRiteshKoushik/attendash/forms"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Create a new event",
	Long:  `Create a new event right from your CLI`,
	Run:   launchEventForm,
}

func init() {
	rootCmd.AddCommand(eventCmd)
}

func launchEventForm(cmd *cobra.Command, args []string) {
	newEvent := api.Event{}
	form := forms.NewEventForm(newEvent)
	if err := form.Run(); err != nil {
		log.Error(err)
	}
}
