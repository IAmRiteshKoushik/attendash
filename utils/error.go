package utils

import "github.com/charmbracelet/lipgloss"

func ErrorString(msg string) string {
	redStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	return redStyle.Render(msg)
}
