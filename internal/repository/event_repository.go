package repository

import (
	"excercise2/internal/domain"
	"fmt"
	"sync"
)

type RepositoryEvent interface {
	CreateEvent
	FindAllEvent
	FindByIdEvent
	UpdateEventTicketStock
}

type CreateEvent interface {
	Create(request *domain.Event) (*domain.Event, error)
}

type FindAllEvent interface {
	FindAll() ([]domain.Event, error)
}

type FindByIdEvent interface {
	FindById(eventId string) (domain.Event, error)
}

type UpdateEventTicketStock interface {
	UpdateEventTicketStock(eventId string, ticketType string) (*domain.Event, error)
}

type repositoryEvent struct {
	db map[string]domain.Event
	sync.Mutex
	countRace int
}

// UpdateEventTicketStock implements RepositoryEvent.
func (repo *repositoryEvent) UpdateEventTicketStock(eventId string, ticketType string) (*domain.Event, error) {
	var indexTicket int
	existingEvent, exists := repo.db[eventId]
	if !exists {
		return &domain.Event{}, fmt.Errorf("event dengan id: %s tidak ditemukan", eventId)
	}

	for i, ticket := range existingEvent.Ticket {
		if ticket.Type == ticketType {
			indexTicket = i
			break
		}
	}
	repo.Mutex.Lock()
	existingEvent.Ticket[indexTicket].Stock -= 1
	repo.countRace++

	repo.Mutex.Unlock()

	// repo.WaitGroup.Wait()
	fmt.Println("indexticket == ", indexTicket)
	fmt.Println(repo.countRace)

	repo.db[eventId] = existingEvent
	return &existingEvent, nil
}

// Create implements RepositoryEvent.
func (repo *repositoryEvent) Create(request *domain.Event) (*domain.Event, error) {
	// var event domain.Event

	// uuidGenerate, _ := uuid.NewV4()
	// event.Id = uuidGenerate.String()
	// event.EventName = request.EventName

	repo.db[request.Id] = *request
	return request, nil

}

// FindAll implements RepositoryEvent.
func (repo *repositoryEvent) FindAll() ([]domain.Event, error) {
	var allEvent []domain.Event

	for _, event := range repo.db {
		allEvent = append(allEvent, event)
	}

	return allEvent, nil
}

// FindById implements RepositoryEvent.
func (repo *repositoryEvent) FindById(eventId string) (domain.Event, error) {

	var bucketEvent domain.Event

	if _, exist := repo.db[eventId]; !exist {
		return domain.Event{}, fmt.Errorf("event dengan id: %s tidak ditemukan", eventId)
	}

	for _, event := range repo.db {
		if event.Id == eventId {
			bucketEvent = event
		}
	}

	return bucketEvent, nil

}

func NewRepositoryEvent() RepositoryEvent {
	return &repositoryEvent{
		db: make(map[string]domain.Event),
	}
}
