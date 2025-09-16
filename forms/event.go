package forms

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/huh"
)

func NewEventForm(e *api.Event) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			// Event Name
			huh.NewInput().Title("Event Name").
				Placeholder("ACM Winter of Code").
				Value(&e.Name),
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
		),

		huh.NewGroup(
			// Event date
			huh.NewInput().
				Title("Event date").
				Value(&e.Day).
				Placeholder("DD").
				Validate(validateDay),
			huh.NewInput().
				Title("Event month").
				Value(&e.Month).
				Placeholder("MM").
				Validate(validateMonth),
			huh.NewInput().
				Title("Event year").
				Value(&e.Year).
				Placeholder("YYYY").Validate(validateYear),
		),
		huh.NewGroup(
			// Event time
			huh.NewInput().
				Title("Hours (24h)").
				Value(&e.Hour).
				Placeholder("HH").Validate(validateHour),
			huh.NewInput().
				Title("Minutes").
				Value(&e.Minute).
				Placeholder("MM").Validate(validateMinute),
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
	).WithLayout(huh.LayoutGrid(1, 4)).WithWidth(120).WithShowErrors(true)

	return form
}

func validateDay(s string) error {
	if s == "" || s != strings.TrimSpace(s) {
		return errors.New("day is either empty or has leading/trailing spaces")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if num < 1 || num > 31 {
		return errors.New("date must be between 1 to 31")
	}

	return nil
}

func validateMonth(s string) error {
	if s == "" || s != strings.TrimSpace(s) {
		return errors.New(
			"month is either empty or has leading/trailing spaces",
		)
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if num < 1 || num > 12 {
		return errors.New("month must be between 1 to 12")
	}

	return nil
}

func validateYear(s string) error {
	if s == "" || s != strings.TrimSpace(s) {
		return errors.New("year is either empty or has leading/trailing spaces")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if num < time.Now().Year() {
		return errors.New("year must be greater or equal to current year")
	}

	return nil
}

func validateHour(s string) error {
	if s == "" || s != strings.TrimSpace(s) {
		return errors.New("hour is either empty or has leading/trailing spaces")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if num < 0 || num > 23 {
		return errors.New("hour must be between 0 and 23")
	}

	return nil
}

func validateMinute(s string) error {
	if s == "" || s != strings.TrimSpace(s) {
		return errors.New(
			"minutes is either empty or has leading/trailing spaces",
		)
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if num < 0 || num > 59 {
		return errors.New("minutes must be between 0 to 59")
	}

	return nil
}
