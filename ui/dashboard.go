package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusIndex uint

const (
	Main focusIndex = iota
	Preview
)

var (
	activeStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(0, 1)

	noStyle = lipgloss.NewStyle().
		Padding(0, 1)
)

type rootModel struct {
	screenWidth  int
	paneSelected focusIndex
	modelsMap    map[focusIndex]tea.Model
}

func (r *rootModel) Init() tea.Cmd { return nil }

func (r *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.screenWidth = msg.Width
		cmds = append(cmds, tea.ClearScreen)
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			r.paneSelected = (r.paneSelected + 1) % 2
		case "ctrl+c", "esc":
			return r, tea.Quit
		}
	}

	for _, m := range r.modelsMap {
		var cmd tea.Cmd
		m, cmd = m.Update(msg)
		cmds = append(cmds, cmd)
	}

	return r, tea.Batch(cmds...)
}

func (r *rootModel) View() string {
	windowSize := r.screenWidth / 2

	leftView := r.modelsMap[Main].View()
	rightView := r.modelsMap[Preview].View()

	if r.paneSelected == Main {
		leftView = activeStyle.Width(windowSize).Render(leftView)
	} else {
		leftView = noStyle.Width(windowSize).Render(leftView)
	}

	if r.paneSelected == Preview {
		leftView = activeStyle.Width(windowSize).Render(leftView)
	} else {
		leftView = noStyle.Width(windowSize).Render(leftView)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, leftView, rightView)
}

func DashboardInit() error {
	// modelsMap := make(map[focusIndex]tea.Model)
	return nil
}
