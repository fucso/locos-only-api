package usecase

import (
	"github.com/fucso/locos-only-api/src/domain/event"
	"github.com/fucso/locos-only-api/src/infrastructure"
	"github.com/fucso/locos-only-api/src/repository"
)

type EventUsecase struct {
	repo *repository.EventRepository
}

func NewEventUsecase(Database *infrastructure.Database) *EventUsecase {
	return &EventUsecase{
		repo: repository.NewEventRepository(Database),
	}
}

func (usecase *EventUsecase) FindAll() ([]*event.Event, error) {
	return usecase.repo.FindAll()
}
