package ui

import (
	"fmt"

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

	for k, m := range r.modelsMap {
		var cmd tea.Cmd
		m, cmd = m.Update(msg)
		r.modelsMap[k] = m
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
		rightView := activeStyle.Width(windowSize).Render(rightView)
		return lipgloss.JoinHorizontal(lipgloss.Left, leftView, rightView)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, leftView, rightView)
}
//test data for now
type simpleModel struct {
	title string
}

func NewMainModel() tea.Model {
	return &simpleModel{title: "Main Pane"}
}

func NewPreviewModel() tea.Model {
	return &simpleModel{title: "Preview Pane"}
}

func (s *simpleModel) Init() tea.Cmd { return nil }

func (s *simpleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *simpleModel) View() string {
	return fmt.Sprintf("\n  %s\n\n  (Press TAB to switch panes, ESC/Ctrl+C to quit)", s.title)
}

func DashboardInit() error {
	modelsMap := map[focusIndex]tea.Model{
		Main:    NewMainModel(),
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
