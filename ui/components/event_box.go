package components

import (
	"github.com/charmbracelet/lipgloss"
)

type EventType int
type EventLabel int

const (
	Online EventType = iota
	Offline
)

const (
	Workshop EventLabel = iota
	Hackathon
	Talks
)

type Event struct {
	ID       int
	Name     string
	Date     string
	Time     string
	Location string
	Type     EventType
	Label    EventLabel
}

func (t EventType) String() string {
	return [...]string{"Offline", "Online"}[t]
}

func (l EventLabel) String() string {
	return [...]string{"Workshop", "Hackathon", "Talks"}[l]
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

	focusedCardStyle = cardStyle.Copy().
				BorderForeground(lipgloss.Color("62")).
				Bold(true)

	labelColors = map[EventLabel]lipgloss.Color{
		Workshop:  lipgloss.Color("196"),
		Hackathon: lipgloss.Color("214"),
		Talks:     lipgloss.Color("34"),
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
		lipgloss.NewStyle().Bold(true).Render(c.Event.Date),
		lipgloss.NewStyle().MarginLeft(2).Render(c.Event.Time),
		lipgloss.NewStyle().Faint(true).MarginLeft(4).Render(c.Event.Type.String()),
		lipgloss.NewStyle().MarginLeft(4).Render(labelStyle.Render(c.Event.Label.String())),
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
