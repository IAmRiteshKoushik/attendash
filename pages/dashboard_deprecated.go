package pages

import (
	"fmt"
	"strings"

	"github.com/IAmRiteshKoushik/attendash/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusIndex int

const (
	Main focusIndex = iota
	Preview
)

var (
	borderColorActive   = lipgloss.Color("62")
	borderColorInactive = lipgloss.Color("240")

	paneTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Bold(true).
			Padding(0, 1)

	paneBaseStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(1, 2).
			Margin(0, 1)

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")).
			MarginTop(1).
			Italic(true)
)

type rootModel struct {
	screenWidth       int
	paneSelected      focusIndex
	modelsMap         map[focusIndex]tea.Model
	selectedEvent     *components.Event
	showingFilePicker bool
	filePicker        components.FilePickerModel
}

func (r *rootModel) Init() tea.Cmd { return nil }

func (r *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.screenWidth = msg.Width
		cmds = append(cmds, tea.ClearScreen)
	case components.FileSelectedMsg:
		r.showingFilePicker = false
		fmt.Println("File picked:", msg.Path)
		return r, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			r.paneSelected = (r.paneSelected + 1) % 2
		case "ctrl+c", "esc":
			return r, tea.Quit

		case "p":
			r.showingFilePicker = true
			r.filePicker = components.NewFilePickerModel()
			return r, r.filePicker.Init()
		}
	case components.EventSelectedMsg:
		r.selectedEvent = &msg.Event
		if pm, ok := r.modelsMap[Preview]; ok {
			var cmd tea.Cmd
			pm, cmd = pm.Update(msg)
			r.modelsMap[Preview] = pm
			cmds = append(cmds, cmd)
		}
	}

	if r.showingFilePicker {
		model, cmd := r.filePicker.Update(msg)
		if _, ok := model.(components.FilePickerModel); ok {
			r.filePicker = model.(components.FilePickerModel)
		}
		return r, cmd
	}

	for k, m := range r.modelsMap {
		var cmd tea.Cmd
		m, cmd = m.Update(msg)
		r.modelsMap[k] = m
		cmds = append(cmds, cmd)
	}

	return r, tea.Batch(cmds...)
}

func (r *rootModel) View() string {
	if r.showingFilePicker {
		return r.filePicker.View()
	}

	windowSize := r.screenWidth / 2

	leftView := r.modelsMap[Main].View()
	rightView := r.modelsMap[Preview].View()

	leftPane := renderPane(
		"Events",
		leftView,
		windowSize,
		r.paneSelected == Main,
	)
	rightPane := renderPane(
		"Details",
		rightView,
		windowSize,
		r.paneSelected == Preview,
	)

	ui := lipgloss.JoinHorizontal(lipgloss.Top, leftPane, rightPane)

	footer := footerStyle.Render(
		"↑/↓ to navigate • ENTER to select • TAB to switch panes • ESC to quit",
	)
	return lipgloss.JoinVertical(lipgloss.Left, ui, footer)
}

func renderPane(title, content string, width int, active bool) string {
	style := paneBaseStyle.Width(width)

	if active {
		style = style.BorderForeground(borderColorActive)
	} else {
		style = style.BorderForeground(borderColorInactive)
	}

	header := paneTitleStyle.Render(title)
	body := lipgloss.NewStyle().MarginTop(1).Render(content)

	return style.Render(lipgloss.JoinVertical(lipgloss.Left, header, body))
}

// test data for now
type simpleModel struct {
	title string
}

type previewModel struct {
	event        *components.Event
	participants map[string][]string
}

func NewPreviewModel() tea.Model {
	return &previewModel{
		participants: map[string][]string{
			"1": {"Alice", "Bob", "Charlie"},
			"2": {"Diana", "Ethan"},
			"3": {"Frank", "Grace", "Hannah", "Ivan"},
		},
	}
}

func (p *previewModel) Init() tea.Cmd { return nil }

func (p *previewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case components.EventSelectedMsg:
		p.event = &msg.Event
	}
	return p, nil
}

func (p *previewModel) View() string {
	if p.event == nil {
		return "Right Pane (Placeholder)\n\n(Select an event to see participants here)"
	}
	list := p.participants[p.event.Id]
	if len(list) == 0 {
		return fmt.Sprintf(
			"Participants for %s\n\n(No participants yet)",
			p.event.Name,
		)
	}
	rows := strings.Join(list, "\n- ")
	return fmt.Sprintf("Participants for %s\n\n- %s", p.event.Name, rows)
}

func (s *simpleModel) View() string {
	return fmt.Sprintf(
		"\n  %s\n\n  (Press TAB to switch panes, ESC/Ctrl+C to quit, p for file picker)",
		s.title,
	)
}

func DbInit() error {
	events := []components.Event{
		{
			Id:        "1",
			Name:      "GoConf",
			Datetime:  "12 Sep 10:00 AM",
			Location:  "Hall A",
			IsOffline: true,
			Label:     "Solo",
		},
		{
			Id:        "2",
			Name:      "Workshop 101",
			Datetime:  "12 Sep 11:30AM",
			Location:  "Hall B",
			IsOffline: true,
			Label:     "Solo",
		},
		{
			Id:        "3",
			Name:      "Tech Meetup",
			Datetime:  "13 Sep 2:00 PM",
			Location:  "Lounge",
			IsOffline: true,
			Label:     "Team",
		},
	}

	modelsMap := map[focusIndex]tea.Model{
		Main:    components.NewEventTable(events),
		Preview: NewPreviewModel(),
	}

	root := &rootModel{
		screenWidth:  80,
		paneSelected: Main,
		modelsMap:    modelsMap,
	}

	p := tea.NewProgram(root, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		return err
	}

	return nil
}
