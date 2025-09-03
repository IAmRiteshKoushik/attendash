package cmd

import (
	"errors"
	"fmt"

	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/IAmRiteshKoushik/attendash/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Action for login to Attendash",
	Long: `login requires an internet connectivity to communicate with Firebase 
where all the admin's details are stored. There is no way to add new admins
other than manually adding them to the database. An admin entry includes an 
email and a SHA-256 hashed password.`,
	RunE: loginRunner,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginRunner(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return errors.New(utils.ErrorString("login does not take any positional arguments"))
	}

	var email string
	var pwd string

	form := huh.NewForm(
		huh.NewGroup(

			huh.NewNote().
				Title("Attendash | Admin Portal for Attendex"),

			huh.NewInput().
				Title("Email").
				Prompt("> ").
				Placeholder("electro@boom.com").
				Validate(utils.ValidateEmail).
				Value(&email),

			huh.NewInput().
				Title("Password").
				Prompt("> ").
				Placeholder("Super secret password!").
				EchoMode(huh.EchoModePassword).
				Validate(utils.ValidatePwd).
				Value(&pwd),
		),
	).WithShowErrors(true)

	err := form.Run()
	if err != nil {
		return err
	}

	_ = spinner.New().
		Title("Verifying your details...").
		Action(func() {
			ok, _ := api.SubmitLoginForm(email, pwd)
			if ok {
				fmt.Println(
					lipgloss.NewStyle().
						Width(50).
						BorderStyle(lipgloss.RoundedBorder()).
						BorderForeground(lipgloss.Color("1")).
						Padding(1, 1).
						Render(utils.ErrorString("Login Failed! Error occured.")),
				)
				return
			}
			fmt.Println(
				lipgloss.NewStyle().
					Width(50).
					BorderStyle(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("63")).
					Padding(1, 1).
					Render("Login Successful!\nPlease type `attendash` to visit your dashboard"),
			)
		}).
		Run()

	return nil
}
