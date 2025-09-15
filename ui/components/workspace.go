package ui

import tea "github.com/charmbracelet/bubbletea"

type workspace struct {
}

func (c workspace) Init() tea.Cmd {
	return nil
}

func (c workspace) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c workspace) View() string {
	return ""
}

func NewWorkspace() workspace {
	return workspace{}
}
