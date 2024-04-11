package usecase

import (
	params "github.com/fucso/locos-only-api/src/controller/params/venue"
	"github.com/fucso/locos-only-api/src/domain/venue"
	"github.com/fucso/locos-only-api/src/infrastructure"
	"github.com/fucso/locos-only-api/src/repository"
)

type VenueUsecase struct {
	repo *repository.VenueRepository
}

func NewVenueUsecase(db *infrastructure.Database) *VenueUsecase {
	return &VenueUsecase{
		repo: repository.NewVenueRepository(db),
	}
}

func (usecase *VenueUsecase) Create(params params.CreateRequest) (*venue.Venue, error) {
	v, err := venue.NewVenueBuilder().Name("test").Build()
	if err != nil {
		return nil, err
	}

	return usecase.repo.Create(v)
}
