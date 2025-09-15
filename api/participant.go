package api

type Participant struct {
	Id        string
	TeamName  string
	Name      string
	Email     string
	IsPresent bool
}

func (p *Participant) New() (*Participant, error) {
	return p, nil
}

func (p *Participant) Edit() (*Participant, error) {
	return p, nil
}

func (p *Participant) Delete() (*Participant, error) {
	return p, nil
}

