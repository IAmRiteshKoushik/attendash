package utils

import (
	"errors"

	"github.com/charmbracelet/lipgloss"
)

func ErrorString(msg string) error {
	redStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	return errors.New(redStyle.Render(msg))
}
