package repository

import (
	"github.com/doug-martin/goqu/v9"

	"github.com/fucso/locos-only-api/src/domain/event"
	"github.com/fucso/locos-only-api/src/infrastructure"
)

type EventRepository struct {
	Database *infrastructure.Database
}

func NewEventRepository(db *infrastructure.Database) *EventRepository {
	return &EventRepository{
		Database: db,
	}
}

func (repo *EventRepository) FindAll() ([]*event.Event, error) {
	ds := goqu.Select("id", "name").From("events")
	rows, err := repo.Database.Select(ds)
	if err != nil {
		return nil, err
	}

	var events []*event.Event
	for rows.Next() {
		var (
			id   int
			name string
		)

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		event, err := event.NewEventBuilder().ID(id).Name(name).Build()
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
