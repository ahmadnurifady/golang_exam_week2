package repository

import (
	"database/sql"
	"excercise2/internal/domain"
)

type CompositeEventTicket interface {
	CreateComposite(tx *sql.Tx, request domain.EventTicket) (domain.EventTicket, error)
	FindByIdComposite(eventTicketId string) (domain.EventTicket, error)
}

type compositeEventTicket struct {
	db *sql.DB
}

// CreateComposite implements CompositeEventTicket.
func (repo *compositeEventTicket) CreateComposite(tx *sql.Tx, request domain.EventTicket) (domain.EventTicket, error) {
	var composite domain.EventTicket

	err := tx.QueryRow("INSERT INTO composite_event_ticket(id, event_id, ticket_id) VALUES ($1, $2, $3) RETURNING id, event_id, ticket_id", request.Id, request.EventId, request.TicketId).Scan(&composite.Id, &composite.EventId, &composite.TicketId)
	if err != nil {
		return domain.EventTicket{}, err
	}

	return composite, nil
}

// FindByIdComposite implements CompositeEventTicket.
func (repo *compositeEventTicket) FindByIdComposite(eventTicketId string) (domain.EventTicket, error) {
	var composite domain.EventTicket

	err := repo.db.QueryRow("SELECT id, event_id, ticket_id FROM composite_event_ticket RETURNING id, event_id, ticket_id").Scan(&composite.Id, &composite.EventId, &composite.TicketId)
	if err != nil {
		return domain.EventTicket{}, err
	}

	return composite, nil
}

func NewRepositoryCompositeEventTicket(db *sql.DB) CompositeEventTicket {
	return &compositeEventTicket{db: db}
}
