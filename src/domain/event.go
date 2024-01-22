package domain

type Event struct {
	ID   int
	Name string
}

func NewEvent(id int, name string) (*Event, error) {
	return &Event{
		ID:   id,
		Name: name,
	}, nil
}
