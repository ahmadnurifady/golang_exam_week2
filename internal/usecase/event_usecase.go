package usecase

import (
	"excercise2/internal/domain"
	"excercise2/internal/repository"
)

type UsecaseEvent interface {
	UsecaseCreateEvent
	UsecaseFindAllEvent
	UsecaseFindByIdEvent
}

type UsecaseCreateEvent interface {
	Create(request *domain.Event) (*domain.Event, error)
}

type UsecaseFindAllEvent interface {
	FindAll() ([]domain.Event, error)
}

type UsecaseFindByIdEvent interface {
	FindById(eventId string) (domain.Event, error)
}

type usecaseEvent struct {
	repo repository.RepositoryEvent
}

// Create implements UsecaseEvent.
func (uc *usecaseEvent) Create(request *domain.Event) (*domain.Event, error) {
	result, err := uc.repo.Create(request)
	if err != nil {
		return &domain.Event{}, err
	}

	return result, nil
}

// FindAll implements UsecaseEvent.
func (uc *usecaseEvent) FindAll() ([]domain.Event, error) {
	result, err := uc.repo.FindAll()
	if err != nil {
		return []domain.Event{}, err
	}

	return result, nil
}

// FindById implements UsecaseEvent.
func (uc *usecaseEvent) FindById(eventId string) (domain.Event, error) {
	result, err := uc.repo.FindById(eventId)
	if err != nil {
		return domain.Event{}, err
	}
	return result, nil

}

func NewUsecaseEvent(repo repository.RepositoryEvent) UsecaseEvent {
	return &usecaseEvent{
		repo: repo,
	}
}
