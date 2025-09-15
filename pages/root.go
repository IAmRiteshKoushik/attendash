package pages

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/IAmRiteshKoushik/attendash/components"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	sidebarView = iota
	workspaceView
)

type root struct {
	pane      int
	sidebar   list.Model
	workspace tea.Model

	showEventForm  bool
	eventForm      *huh.Form
	eventFormState api.Event

	showParticipantForm  bool
	participantForm      *huh.Form
	participantFormState api.Participant
}

func (r root) Init() tea.Cmd {
	return nil
}

func (r root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return r, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := components.SidebarStyle.GetFrameSize()
		r.sidebar.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	r.sidebar, cmd = r.sidebar.Update(msg)
	return r, cmd
}

func (r root) View() string {
	return components.SidebarStyle.Render(r.sidebar.View())
}

func NewRoot() root {
	r := root{
		pane:      sidebarView,
		sidebar:   components.NewSidebar(),
		workspace: components.NewWorkspace(),
	}

	return r
}
