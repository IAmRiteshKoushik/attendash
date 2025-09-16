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
	TeamSize  int

	// Datetime is deserialized into this for simpler handling
	Day    string
	Month  string
	Year   string
	Hour   string
	Minute string
}

// The Title() method is used by the sidebar component to render the title
func (e Event) Title() string {
	return e.Name
}

// The description method is used by the sidebar to render the description
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

// The filter value method is used by the sidebar to fuzzy find events
func (e Event) FilterValue() string {
	return e.Name
}

// The following methods involve invoking the AppWrite SDK for pushing into DB
func (e *Event) New() (*Event, error) {
	return e, nil
}

func (e *Event) Edit() (*Event, error) {
	return e, nil
}

func (e *Event) Delete() error {
	return nil
}
