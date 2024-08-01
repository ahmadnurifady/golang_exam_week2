package repository

import (
	"database/sql"
	"excercise2/internal/domain"
	"fmt"

	"github.com/rs/zerolog/log"
)

type RepositoryTicket interface {
	Create(request domain.Ticket) (domain.Ticket, error)
	CreateTicketWithTx(tx *sql.Tx, request domain.Ticket) (domain.Ticket, error)
	FindAll()
	FindById(tx *sql.Tx, ticketId string) (domain.Ticket, error)
	FindByEventName(tx *sql.Tx, ticketName string) ([]domain.Ticket, error)
	UpdateStock(tx *sql.Tx, ticketId string) (domain.Ticket, error)
}

type repositoryTicket struct {
	db       map[int]domain.Ticket
	database *sql.DB
}

// CreateTicketWithTx implements RepositoryTicket.
func (repo *repositoryTicket) CreateTicketWithTx(tx *sql.Tx, request domain.Ticket) (domain.Ticket, error) {
	var ticket domain.Ticket

	err := tx.QueryRow("INSERT INTO ticket (id, type, price, stock, event_name) VALUES ($1, $2, $3, $4, $5) RETURNING id, type, price, stock, event_name", request.Id, request.Type, request.Price, request.Stock, request.EventName).Scan(
		&ticket.Id,
		&ticket.Type,
		&ticket.Price,
		&ticket.Stock,
		&ticket.EventName,
	)
	if err != nil {
		return domain.Ticket{}, err
	}
	return ticket, nil
}

// Create implements RepositoryTicket.
func (repo *repositoryTicket) Create(request domain.Ticket) (domain.Ticket, error) {
	var ticket domain.Ticket

	err := repo.database.QueryRow("INSERT INTO ticket (id, type, price, stock, event_name) VALUES ($1, $2, $3, $4, $5) RETURNING id, type, price, stock, event_name", request.Id, request.Type, request.Price, request.Stock, request.EventName).Scan(
		&ticket.Id,
		&ticket.Type,
		&ticket.Price,
		&ticket.Stock,
		&ticket.EventName,
	)
	if err != nil {
		return domain.Ticket{}, err
	}
	return ticket, nil
}

// UpdateStock implements RepositoryTicket.
func (repo *repositoryTicket) UpdateStock(tx *sql.Tx, ticketId string) (domain.Ticket, error) {

	result, err := tx.Exec("UPDATE ticket SET stock = stock - 1 WHERE id = $1", ticketId)
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]", err).Msg("")
		return domain.Ticket{}, err
	}

	log.Info().Any("SUCCESS at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]", result).Msg("")

	return domain.Ticket{}, nil

}

// FindByName implements RepositoryTicket.
func (repo *repositoryTicket) FindByEventName(tx *sql.Tx, ticketName string) ([]domain.Ticket, error) {
	var allTicket []domain.Ticket

	fmt.Println("eventName  ==", ticketName)

	rows, err := tx.Query("SELECT id, type, price, stock, event_name FROM ticket WHERE event_name = $1", ticketName)
	if err != nil {
		return []domain.Ticket{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket domain.Ticket

		err := rows.Scan(&ticket.Id, &ticket.Type, &ticket.Price, &ticket.Stock, &ticket.EventName)
		if err != nil {
			return []domain.Ticket{}, err
		}
		allTicket = append(allTicket, ticket)
	}

	fmt.Println("repo ticket all ticket == ", allTicket)

	return allTicket, nil

}

// FindAll implements RepositoryTicket.
func (repo *repositoryTicket) FindAll() {
	panic("unimplemented")
}

// FindById implements RepositoryTicket.
func (repo *repositoryTicket) FindById(tx *sql.Tx, ticketId string) (domain.Ticket, error) {
	var ticket domain.Ticket

	err := tx.QueryRow("SELECT id, type, price, stock, event_name FROM ticket WHERE id = $1", ticketId).Scan(&ticket.Id, &ticket.Type, &ticket.Price, &ticket.Stock, &ticket.EventName)
	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
}

func NewRepositoryTicket(database *sql.DB) RepositoryTicket {
	return &repositoryTicket{
		db:       make(map[int]domain.Ticket),
		database: database,
	}
}
