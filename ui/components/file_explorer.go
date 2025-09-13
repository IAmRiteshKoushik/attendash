package components

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/IAmRiteshKoushik/attendash/utils"
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type FileSelectedMsg struct {
	Path string
}

type FilePickerModel struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
}

func NewFilePicker() FilePickerModel {
	fp := filepicker.New()
	fp.SetHeight(20)

	fp.AllowedTypes = []string{".csv"}
	fp.ShowHidden = false
	fp.DirAllowed = true
	fp.CurrentDirectory, _ = os.UserHomeDir()
	fp.CurrentDirectory += "/Downloads" // setting default

	return FilePickerModel{
		filepicker: fp,
	}
}

func NewFilePickerModel() FilePickerModel {
	return NewFilePicker()
}

func (m FilePickerModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m FilePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case utils.ClearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// If a selection has happened, then get the path
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
		selectedCmd := func() tea.Msg {
			return FileSelectedMsg{Path: path}
		}

		return m, tea.Batch(cmd, selectedCmd)
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, utils.ClearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m FilePickerModel) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a file:")
	} else {
		s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}
