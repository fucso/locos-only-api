package usecase

import (
	"github.com/fucso/locos-only-api/src/domain"
	"github.com/fucso/locos-only-api/src/repository"
)

type EventUsecase struct {
	EventRepository *repository.EventRepository
}

func NewEventUsecase(repo *repository.EventRepository) *EventUsecase {
	return &EventUsecase{
		EventRepository: repo,
	}
}

func (usecase *EventUsecase) FindAll() ([]*domain.Event, error) {
	return usecase.EventRepository.FindAll()
}
