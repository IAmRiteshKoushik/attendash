package components

import tea "github.com/charmbracelet/bubbletea"

type Workspace struct{}

func (c Workspace) Init() tea.Cmd {
	return nil
}

func (c Workspace) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c Workspace) View() string {
	return ""
}

func NewWorkspace() Workspace {
	return Workspace{}
}
