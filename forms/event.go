package forms

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/huh"
)

func NewEventForm(e api.Event) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Event Name
			huh.NewInput().Title("Event Name").
				Placeholder("ACM Winter of Code").
				Value(&e.Name).
				Validate(func(s string) error {
					return nil
				}),
			// Event Location
			huh.NewInput().Title("Event Location").
				Value(&e.Location).
				Placeholder("Venue (offline) / Platform (online)"),
			// Event IsOffline
			huh.NewSelect[bool]().Title("Event Mode").
				Options(
					huh.NewOption("Offline", true).Selected(e.IsOffline),
					huh.NewOption("Online", false).Selected(!e.IsOffline)).
				Value(&e.IsOffline),
		).WithShowErrors(true),

		huh.NewGroup(
			// Event date
			huh.NewInput().Title("Event date").Value(&e.Day).Placeholder("DD"),
			huh.NewInput().
				Title("Event month").
				Value(&e.Month).
				Placeholder("MM"),
			huh.NewInput().
				Title("Event year").
				Value(&e.Year).
				Placeholder("YYYY"),
		).WithShowErrors(true),
		huh.NewGroup(
			// Event time
			huh.NewInput().
				Title("Hours (24h)").
				Value(&e.Hour).
				Placeholder("HH"),
			huh.NewInput().
				Title("Minutes").
				Value(&e.Minute).
				Placeholder("MM"),
		),
		huh.NewGroup(
			// Event Label
			huh.NewSelect[string]().Title("Label").
				Options(
					huh.NewOption("Solo", "Solo").Selected(e.Label == "Solo"),
					huh.NewOption("Team", "Team").Selected(e.Label == "Team")).
				Value(&e.Label),
			// Submit button
			huh.NewConfirm().
				Title("Submit the form ?").
				Affirmative("Yes").
				Negative("No."),
		),
	).WithLayout(huh.LayoutGrid(2, 2)).WithWidth(75)
	return form
}
