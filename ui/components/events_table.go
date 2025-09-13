package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type EventSelectedMsg struct {
	Event Event
}

type EventTable struct {
	events     []Event
	eventCards []EventCard
	cursor     int
}

func NewEventTable(events []Event) *EventTable {
	cards := make([]EventCard, len(events))
	for i, e := range events {
		cards[i] = *NewEventCard(e)
	}
	return &EventTable{
		events:     events,
		eventCards: cards,
	}
}

func (m *EventTable) Init() tea.Cmd {
	return nil
}

func (m *EventTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.events)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.events) > 0 {
				selected := m.events[m.cursor]
				// Send message upwards
				return m, func() tea.Msg {
					return EventSelectedMsg{Event: selected}
				}
			}
		}
	}
	return m, nil
}

func (m *EventTable) View() string {
	var rows []string
	for i := range m.eventCards {
		m.eventCards[i].Focused = (i == m.cursor)
		rows = append(rows, m.eventCards[i].View())
	}
	return strings.Join(rows, "\n")
}
