package utils

import (
	"errors"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func ErrorString(msg string) error {
	redStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	return errors.New(redStyle.Render(msg))
}

type ClearErrorMsg struct{}

func ClearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}
