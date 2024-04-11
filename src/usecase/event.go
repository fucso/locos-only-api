package usecase

import (
	params "github.com/fucso/locos-only-api/src/controller/params/event"
	"github.com/fucso/locos-only-api/src/domain/event"
	"github.com/fucso/locos-only-api/src/infrastructure"
	"github.com/fucso/locos-only-api/src/repository"
)

type EventUsecase struct {
	repo *repository.EventRepository
}

func NewEventUsecase(db *infrastructure.Database) *EventUsecase {
	return &EventUsecase{
		repo: repository.NewEventRepository(db),
	}
}

func (usecase *EventUsecase) FindAll() ([]*event.Event, error) {
	return usecase.repo.FindAll()
}

func (usecase *EventUsecase) Create(p *params.CreateRequest) (*event.Event, error) {
	e, err := event.NewEventBuilder().Name(p.Name).Build()
	if err != nil {
		return nil, err
	}

	e, err = usecase.repo.Create(e)
	if err != nil {
		return nil, err
	}

	return e, nil
}
