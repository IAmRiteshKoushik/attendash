package api

type Event struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type EventWithID struct {
	ID string `json:"id"`
	Event
}

type Participant struct {
	Name string
	Roll string
}

type ParticipantWithID struct {
	ID string `json:"id"`
	Participant
}

func FetchAllEvents() []EventWithID {
	return nil
}

func (e *Event) FetchEventParticipants() []ParticipantWithID {
	return nil
}

func (e *Event) CreateEvent() EventWithID {
	return EventWithID{}
}

func (e *Event) EditEvent() EventWithID {
	return EventWithID{}
}

func (e *EventWithID) PopulateEvent(p []Participant) {
}
