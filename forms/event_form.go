package forms

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/huh"
)

func NewEventForm(e api.Event) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Event").
				Description("Create or edit your event"),
			// Event Name
			huh.NewInput().Title("").
				Placeholder("ACM Winter of Code").
				Value(&e.Name).
				Validate(func(s string) error {
					return nil
				}),
			// Event Location
			huh.NewInput().Title("Location").
				Value(&e.Location).
				Placeholder("Venue (offline) / Platform (online)"),
			// Event IsOffline
			huh.NewSelect[bool]().Title("Event Mode").
				Options(huh.NewOption("Offline", true), huh.NewOption("Online", false)).
				Value(&e.IsOffline),
			// Event Date and Time
			// Event Label
			huh.NewSelect[string]().Title("Label").
				Options(huh.NewOptions("Solo", "Team")...).
				Value(&e.Label),
			// Submit button
			huh.NewConfirm().
				Title("Submit the form ?").
				Affirmative("Yes").
				Negative("No."),
		),
	)

	return form
}
