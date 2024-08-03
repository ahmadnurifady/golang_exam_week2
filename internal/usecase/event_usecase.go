package usecase

import (
	"context"
	"database/sql"
	"excercise2/internal/domain"
	"excercise2/internal/repository"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type UsecaseEvent interface {
	UsecaseCreateEvent
	UsecaseFindAllEvent
	UsecaseFindByIdEvent
}

type UsecaseCreateEvent interface {
	Create(ctx context.Context, request domain.Event) (domain.Event, error)
}

type UsecaseFindAllEvent interface {
	FindAll(ctx context.Context) ([]domain.Event, error)
}

type UsecaseFindByIdEvent interface {
	FindById(ctx context.Context, eventId string) (domain.Event, error)
}

type usecaseEvent struct {
	repo            repository.RepositoryEvent
	repoTicket      repository.RepositoryTicket
	repoEventTicket repository.CompositeEventTicket
	database        *sql.DB
}

// Create implements UsecaseEvent.
func (uc *usecaseEvent) Create(ctx context.Context, request domain.Event) (domain.Event, error) {

	tx, err := uc.database.Begin()
	if err != nil {
		return domain.Event{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var idAllEvent []string

	for _, tkt := range request.Ticket {
		uuidGenerate, _ := uuid.NewV4()
		tkt.Id = uuidGenerate.String()
		tkt.EventName = request.EventName
		_, err := uc.repoTicket.CreateTicketWithTx(tx, tkt)
		if err != nil {
			log.Info().Any("Error In [USECASE] start create TICKET [Create]", err.Error()).Msg("")
			return domain.Event{}, err
		}
		idAllEvent = append(idAllEvent, tkt.Id)
	}

	result, err := uc.repo.Create(tx, ctx, domain.Event{
		Id:        request.Id,
		EventName: request.EventName,
		Location:  request.Location,
		Date:      request.Date,
	})
	if err != nil {
		log.Info().Any("Error In [USECASE] start create EVENT [Create]", err.Error()).Msg("")
		return domain.Event{}, err
	}

	for _, idEvent := range idAllEvent {
		uuidGenerate, _ := uuid.NewV4()
		_, err = uc.repoEventTicket.CreateComposite(tx, domain.EventTicket{
			Id:       uuidGenerate.String(),
			EventId:  result.Id,
			TicketId: idEvent,
		})
		if err != nil {
			log.Info().Any("Error In [USECASE] start composite TICKER [Create]", err.Error()).Msg("")
			return domain.Event{}, err
		}
	}
	result.Ticket = request.Ticket

	err = tx.Commit()
	if err != nil {
		return domain.Event{}, err
	}

	return result, nil
}

// FindAll implements UsecaseEvent.
func (uc *usecaseEvent) FindAll(ctx context.Context) ([]domain.Event, error) {

	var allEventWithTicket []domain.Event

	tx, err := uc.database.Begin()
	if err != nil {
		log.Info().Any("ERROR at [USECASE] - [EVENT] - [FindAll] - [tx begin]", err).Msg("")
		return []domain.Event{}, err
	}

	defer func() {
		if err != nil {
			log.Info().Any("ERROR at [USECASE] - [EVENT] - [FindAll] - [ROLLBACK]", err).Msg("")
			tx.Rollback()
		}
	}()

	allEvent, err := uc.repo.FindAll(tx, ctx)
	if err != nil {
		log.Info().Any("ERROR at [USECASE] - [EVENT] - [FindAll] - [get data from repo event find all]", err).Msg("")
		return []domain.Event{}, err
	}

	for _, event := range allEvent {
		allTicket, err := uc.repoTicket.FindByEventName(tx, event.EventName)
		if err != nil {
			log.Info().Any("ERROR at [USECASE] - [EVENT] - [FindAll] - [get data from repo ticket findByEventName]", err).Msg("")
			return []domain.Event{}, err
		}
		event.Ticket = allTicket
		allEventWithTicket = append(allEventWithTicket, event)
	}

	err = tx.Commit()
	if err != nil {
		return []domain.Event{}, err
	}

	return allEventWithTicket, nil

}

// FindById implements UsecaseEvent.
func (uc *usecaseEvent) FindById(ctx context.Context, eventId string) (domain.Event, error) {

	return domain.Event{}, nil

}

func NewUsecaseEvent(repo repository.RepositoryEvent, repoTicket repository.RepositoryTicket, repoEventTicket repository.CompositeEventTicket, database *sql.DB) UsecaseEvent {
	return &usecaseEvent{
		repo:            repo,
		repoTicket:      repoTicket,
		repoEventTicket: repoEventTicket,
		database:        database,
	}
}
