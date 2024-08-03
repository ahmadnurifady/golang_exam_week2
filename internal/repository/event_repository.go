package repository

import (
	"context"
	"database/sql"
	"excercise2/internal/domain"
	"sync"

	"github.com/rs/zerolog/log"
)

type RepositoryEvent interface {
	CreateEvent
	FindAllEvent
	FindByIdEvent
}

type CreateEvent interface {
	Create(tx *sql.Tx, ctx context.Context, request domain.Event) (domain.Event, error)
}

type FindAllEvent interface {
	FindAll(tx *sql.Tx, ctx context.Context) ([]domain.Event, error)
}

type FindByIdEvent interface {
	FindById(tx *sql.Tx, ctx context.Context, eventId string) (domain.Event, error)
}

type repositoryEvent struct {
	database *sql.DB
	sync.WaitGroup
}

// Create implements RepositoryEvent.
func (repo *repositoryEvent) Create(tx *sql.Tx, ctx context.Context, request domain.Event) (domain.Event, error) {
	var event domain.Event

	err := tx.QueryRowContext(ctx, "INSERT INTO event(id, event_name, location, date) VALUES($1, $2, $3, $4) RETURNING id, event_name, location, date", request.Id, request.EventName, request.Location, request.Date).Scan(&event.Id, &event.EventName, &event.Location, &event.Date)
	if err != nil {
		log.Info().Any("Error In [REPOSITORY] start create EVENT [Create]", err.Error()).Msg("")
		return domain.Event{}, err
	}

	return event, nil

}

// FindAll implements RepositoryEvent.
func (repo *repositoryEvent) FindAll(tx *sql.Tx, ctx context.Context) ([]domain.Event, error) {
	var allEvent []domain.Event

	rows, err := tx.QueryContext(ctx, "SELECT id, event_name, location, date FROM event")
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [EVENT] - [FindAll] - [get data from database]", err).Msg("")
		return []domain.Event{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var data domain.Event
		err := rows.Scan(&data.Id, &data.EventName, &data.Location, &data.Date)
		if err != nil {
			log.Info().Any("ERROR at [REPOSITORY] - [EVENT] - [FindAll] - [scan data event]", err).Msg("")
			return []domain.Event{}, err
		}
		allEvent = append(allEvent, data)
	}

	return allEvent, nil
}

// FindById implements RepositoryEvent.
func (repo *repositoryEvent) FindById(tx *sql.Tx, ctx context.Context, eventId string) (domain.Event, error) {

	var event domain.Event

	err := tx.QueryRowContext(ctx, "SELECT id, event_name, location, date FROM event WHERE id = $1", eventId).Scan(
		&event.Id,
		&event.EventName,
		&event.Location,
		&event.Date,
	)
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [EVENT] - [FindById] - [scan data from query]", err.Error()).Msg("")
		return domain.Event{}, err
	}

	return event, nil

}

func NewRepositoryEvent(database *sql.DB) RepositoryEvent {
	return &repositoryEvent{
		database: database,
	}
}
