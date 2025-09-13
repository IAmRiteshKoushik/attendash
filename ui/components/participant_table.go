package components

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Participant struct {
	Name      string
	TeamName  string
	Email     string
	IsPresent bool
}

// table model
type Model struct {
	table table.Model
}

// test data
var sampleParticipants = []Participant{
	{Name: "Ada Lovelace", TeamName: "Pioneers", Email: "ada@example.com", IsPresent: true},
	{Name: "Grace Hopper", TeamName: "Pioneers", Email: "grace@example.com", IsPresent: false},
	{Name: "Alan Turing", TeamName: "Enigmas", Email: "alan@example.com", IsPresent: true},
	{Name: "Linus Torvalds", TeamName: "Kernel Crew", Email: "linus@example.com", IsPresent: false},
	{Name: "Margaret Hamilton", TeamName: "Apollo", Email: "margaret@example.com", IsPresent: true},
	{Name: "Vint Cerf", TeamName: "Internet Fathers", Email: "vint@example.com", IsPresent: true},
}

var (
	// present -> green
	presentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	// absent -> red
	absentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)

func ParticipantsTable() Model {
	totalWidth := 100
	colWidth := totalWidth / 3

	columns := []table.Column{
		{Title: "Name", Width: colWidth},
		{Title: "Team Name", Width: colWidth},
		{Title: "Email ID", Width: colWidth},
	}

	// making rows
	rows := make([]table.Row, len(sampleParticipants))
	for i, p := range sampleParticipants {
		style := absentStyle
		if p.IsPresent {
			style = presentStyle
		}
		rows[i] = table.Row{
			style.Render(p.Name),
			style.Render(p.TeamName),
			style.Render(p.Email),
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		// table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return Model{table: t}
}

// Init
func (m Model) Init() tea.Cmd {
	return nil
}

// Update
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View
func (m Model) View() string {
	return m.table.View()
}
