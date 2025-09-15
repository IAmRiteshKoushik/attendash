package components

import (
	"github.com/IAmRiteshKoushik/attendash/api"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var events = []list.Item{
	api.Event{
		Id:        "e1",
		Name:      "Raspberry Pi’s Special",
		Location:  "Anugraha Hall",
		IsOffline: true,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Team",
	},
	api.Event{
		Id:        "e2",
		Name:      "Winter of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e3",
		Name:      "Summer of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e1",
		Name:      "Raspberry Pi’s Special",
		Location:  "Anugraha Hall",
		IsOffline: true,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Team",
	},
	api.Event{
		Id:        "e2",
		Name:      "Winter of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e3",
		Name:      "Summer of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e1",
		Name:      "Raspberry Pi’s Special",
		Location:  "Anugraha Hall",
		IsOffline: true,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Team",
	},
	api.Event{
		Id:       "e2",
		Name:     "Winter of Code",
		Location: "None", IsOffline: false,
		Datetime: "Mon 20th Jan, 2025 10:00PM",
		Label:    "Solo",
	},
	api.Event{
		Id:        "e3",
		Name:      "Summer of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e1",
		Name:      "Raspberry Pi’s Special",
		Location:  "Anugraha Hall",
		IsOffline: true,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Team",
	},
	api.Event{
		Id:        "e2",
		Name:      "Winter of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
	api.Event{
		Id:        "e3",
		Name:      "Summer of Code",
		Location:  "None",
		IsOffline: false,
		Datetime:  "Mon 20th Jan, 2025 10:00PM",
		Label:     "Solo",
	},
}

var SidebarStyle = lipgloss.NewStyle().Margin(1, 0).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")).
	BorderTop(false).
	BorderLeft(false).
	BorderBottom(false).
	BorderRight(true).
	Padding(0, 2, 0, 0)

func NewSidebar() list.Model {
	defDelegate := list.NewDefaultDelegate()
	defDelegate.SetHeight(3)

	sidebar := list.New([]list.Item{}, defDelegate, 10, 50)

	sidebar.Title = "Attendash: TUI Admin Platform"
	sidebar.SetShowHelp(false)
	sidebar.SetStatusBarItemName("event", "events")
	sidebar.FilterInput.CharLimit = 15
	sidebar.FilterInput.Width = 15
	sidebar.FilterInput.Prompt = "> "
	sidebar.FilterInput.Placeholder = "Search Events"
	sidebar.InfiniteScrolling = true

	return sidebar
}
