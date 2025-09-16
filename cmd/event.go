package cmd

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/IAmRiteshKoushik/attendash/forms"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	"fmt"
	"time"
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
	form := forms.NewEventForm(&newEvent)
	if err := form.Run(); err != nil {
		log.Error(err)
	}

	dateTimeString := fmt.Sprintf("%s-%s-%sT%s:%s:00Z", newEvent.Year, newEvent.Month, newEvent.Day, newEvent.Hour, newEvent.Minute)

	// This block is essential to prevent crashes and ensure correct format.
	parsedTime, err := time.Parse("2006-01-02T15:04:05Z", dateTimeString)
	if err != nil {
		log.Error("Invalid date or time entered", "err", err, "\n", dateTimeString)
		fmt.Println(dateTimeString)
		return
	}
	newEvent.Datetime = parsedTime.Format(time.RFC3339)

	if _, err := CreateEvent(&newEvent); err != nil {
		log.Error("Failed to create event in Appwrite", "err", err)
		return
	}

	log.Infof("Successfully created event: '%s'", newEvent.Name)

}

func CreateEvent(e *api.Event) (*models.Row, error) {
	eventDate := map[string]interface{}{
		"eventName":     e.Name,
		"isOffline":     e.IsOffline,
		"eventLocation": e.Location,
		"dateAndTime":   e.Datetime,
		"label":         e.Label,
	}

	log.Infof("dbId: %s", dbName)
	log.Infof("tableId: %s", eventsTable)
	log.Infof("event data: %+v", eventDate)
	log.Infof("Orm is nil? %v", Orm == nil)

	doc, err := Orm.CreateRow(
		dbName,
		eventsTable,
		id.Unique(),
		eventDate,
	)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
