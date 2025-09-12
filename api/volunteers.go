package api

import "time"

type Volunteer struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Year  int       `json:"year"`
	Added time.Time `json:"added"`
}

type VolunteerWithID struct {
	ID string `json:"id"`
	Volunteer
}

func FetchAllVolunteers() []VolunteerWithID {
	return nil
}

func (e *EventWithID) FetchVolunteersByEvent() []VolunteerWithID {
	return nil
}

func (v *Volunteer) CreateVolunteers() VolunteerWithID {
	return VolunteerWithID{}
}

func (v *VolunteerWithID) EditVolunteer() VolunteerWithID {
	return VolunteerWithID{}
}

func (e *EventWithID) AddVolunteersToEvent() bool {
	return true
}

func (e *EventWithID) RemoveVolunteerFromEvent() bool {
	return true
}
