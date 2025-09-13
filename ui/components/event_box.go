package components

import (
	"github.com/charmbracelet/lipgloss"
)

type Event struct {
	Id        string
	Name      string
	Location  string
	IsOffline bool
	Datetime  string
	Label     string // "Solo" | "Team"
}

func (e *Event) CheckOffline() string {
	if e.IsOffline == true {
		return "Offline"
	} else {
		return "Online"
	}
}

type EventCard struct {
	Event   Event
	Focused bool
}

var (
	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1).
			Margin(0, 0, 1, 0).
			Width(60)

	focusedCardStyle = cardStyle.
				BorderForeground(lipgloss.Color("62")).
				Bold(true)

	labelColors = map[string]lipgloss.Color{
		"Solo": lipgloss.Color("196"),
		"Team": lipgloss.Color("214"),
	}
)

func NewEventCard(e Event) *EventCard {
	return &EventCard{Event: e}
}

func (c *EventCard) View() string {
	style := cardStyle
	if c.Focused {
		style = focusedCardStyle
	}

	labelStyle := lipgloss.NewStyle().
		Foreground(labelColors[c.Event.Label]).
		Bold(true)

	row1 := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Bold(true).Render(c.Event.Datetime),
		lipgloss.NewStyle().Faint(true).MarginLeft(4).Render(c.Event.CheckOffline()),
		lipgloss.NewStyle().MarginLeft(4).Render(labelStyle.Render(c.Event.Label)),
	)
	row2 := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("229")).
		MarginTop(1).
		Render(c.Event.Name)
	row3 := lipgloss.NewStyle().
		Italic(true).
		Faint(true).
		MarginTop(1).
		Render(c.Event.Location)

	return style.Render(lipgloss.JoinVertical(lipgloss.Left, row1, row2, row3))
}
