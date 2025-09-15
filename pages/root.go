package pages

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	sidebarView = iota
	workspaceView
)

type root struct {
	showEventForm        bool
	showParticipantForm  bool
	pane                 int
	sidebar              list.Model
	workspace            tea.Model
	eventForm            *huh.Form
	eventFormState       api.Event
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
		h, v := sidebarStyle.GetFrameSize()
		r.sidebar.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	r.sidebar, cmd = r.sidebar.Update(msg)
	return r, cmd
}

func (r root) View() string {
	return sidebarStyle.Render(r.sidebar.View())
}

func NewRoot() root {
	r := root{
		pane:      sidebarView,
		sidebar:   NewSidebar(),
		workspace: NewWorkspace(),
	}

	return r
}
