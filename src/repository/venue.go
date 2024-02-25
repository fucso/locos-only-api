package repository

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/fucso/locos-only-api/src/domain/venue"
	"github.com/fucso/locos-only-api/src/infrastructure"
)

type VenueRepository struct {
	Database *infrastructure.Database
}

func NewVenueRepository(db *infrastructure.Database) *VenueRepository {
	return &VenueRepository{
		Database: db,
	}
}

// Select
func (repo *VenueRepository) Find(id int) (*venue.Venue, error) {
	ds := goqu.Select("id", "name").From("venues").Where(goqu.Ex{"id": id})
	rows, err := repo.Database.Select(ds)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var (
			id   int
			name string
		)

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		return venue.NewVenueBuilder().ID(id).Name(name).Build()
	} else {
		return nil, errors.New("not found")
	}
}

// Insert
func (repo *VenueRepository) Create(v *venue.Venue) (*venue.Venue, error) {
	value := [][]interface{}{{interface{}(v.Name)}}
	id, err := repo.Database.Insert("venues", []interface{}{"name"}, value)
	if err != nil {
		return nil, err
	}

	return repo.Find(*id)
}
