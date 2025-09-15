package forms

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/huh"
)

func NewParticipantForm(isTeam bool, p api.Participant) *huh.Form {
	fields := []huh.Field{
		huh.NewNote().
			Title("Participant").
			Description("Create or edit participant details"),
		// Full Name
		huh.NewInput().
			Title("Full name").
			Placeholder("IAmRiteshKoushik").
			Value(&p.Name).
			Validate(validateName),
		// Email
		huh.NewInput().
			Title("Student email").
			Placeholder("roll@cb.students.amrita.edu").
			Validate(validateEmail).Value(&p.Email),
		// Present validation
		huh.NewSelect[bool]().Title("IsPresent").
			Value(&p.IsPresent).
			Options(
				huh.NewOption("Present", true),
				huh.NewOption("Absent", false),
			),
		// Form Submission
		huh.NewConfirm().
			Title("Confirm your changes").
			Affirmative("Yes").
			Negative("No"),
	}

	// Insert team field conditionally in index 1 if it is a 'team' form
	if isTeam {
		teamField := huh.NewInput().Title("Team name").Value(&p.TeamName)
		fields = append(
			fields[:1],
			append([]huh.Field{teamField}, fields[1:]...)...)
	}
	form := huh.NewForm(huh.NewGroup(fields...))

	return form
}

func validateName(input string) error {
	for _, i := range input {
		if !(unicode.IsLetter(i) || i == ' ') {
			return errors.New(
				"full name can only contain letters and spaces",
			)
		}
	}
	return nil
}

func validateEmail(input string) error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(input) {
		return errors.New("invalid email address")
	}
	return nil
}
