package repository

import (
	"database/sql"
	"errors"

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

// Select
func (repo *EventRepository) FindAll() ([]*event.Event, error) {
	ds := goqu.Select("id", "name").From("events")
	rows, err := repo.Database.Select(ds)
	if err != nil {
		return nil, err
	}

	var events []*event.Event
	for rows.Next() {
		event, err := buildModel(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (repo *EventRepository) Find(id int) (*event.Event, error) {
	ds := goqu.Select("id", "name").From("events").Where(goqu.Ex{"id": id})
	rows, err := repo.Database.Select(ds)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		return buildModel(rows)
	} else {
		return nil, errors.New("not found")
	}
}

func buildModel(rows *sql.Rows) (*event.Event, error) {
	var (
		id   int
		name string
	)

	if err := rows.Scan(&id, &name); err != nil {
		return nil, err
	}

	return event.NewEventBuilder().ID(id).Name(name).Build()
}

// Insert
func (repo *EventRepository) Create(e *event.Event) (*event.Event, error) {
	value := [][]interface{}{{interface{}(e.Name)}}
	id, err := repo.Database.Insert("events", []interface{}{"name"}, value)
	if err != nil {
		return nil, err
	}

	return repo.Find(*id)
}
