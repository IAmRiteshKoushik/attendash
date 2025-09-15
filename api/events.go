package api

import (
	"fmt"
	"strings"
)

type Event struct {
	Id        string
	Name      string
	Location  string
	IsOffline bool
	Datetime  string
	Label     string // "Solo" | "Team"
}

func (e Event) Title() string {
	return e.Name
}

func (e Event) Description() string {
	var sb strings.Builder

	// row1: Add DateTime && Online | Offline
	sb.WriteString(fmt.Sprintf("%s | ", e.Datetime))
	if e.IsOffline {
		sb.WriteString("Offline\n")
	} else {
		sb.WriteString("Online\n")
	}

	// row3: Solo | Team && Location
	sb.WriteString(fmt.Sprintf("%s | ", e.Label))
	sb.WriteString(fmt.Sprintf("%s\n", e.Location))

	return sb.String()
}

func (e Event) FilterValue() string {
	return e.Name
}
