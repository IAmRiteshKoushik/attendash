package pages

import (
	"github.com/IAmRiteshKoushik/attendash/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func DashboardInit() error {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return utils.ErrorString("Error running the program")
	}
	return nil
}
