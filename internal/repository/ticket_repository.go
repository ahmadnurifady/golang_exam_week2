package repository

import (
	"database/sql"
	"excercise2/internal/domain"
	"sync"

	"github.com/rs/zerolog/log"
)

type RepositoryTicket interface {
	CreateTicket
	CreateTicketWithTx
	FindById
	FindByEventName
	UpdateStock
	UpdateStockNotOne
}

type UpdateStockNotOne interface {
	UpdateStockNotOne(tx *sql.Tx, ticketId string, howManyUpdateStock int) (domain.Ticket, error)
}

type CreateTicket interface {
	Create(request domain.Ticket) (domain.Ticket, error)
}

type CreateTicketWithTx interface {
	CreateTicketWithTx(tx *sql.Tx, request domain.Ticket) (domain.Ticket, error)
}

type FindById interface {
	FindById(tx *sql.Tx, ticketId string) (domain.Ticket, error)
}

type FindByEventName interface {
	FindByEventName(tx *sql.Tx, ticketName string) ([]domain.Ticket, error)
}

type UpdateStock interface {
	UpdateStock(tx *sql.Tx, ticketId string) (domain.Ticket, error)
}

type repositoryTicket struct {
	database *sql.DB
	sync.Mutex
}

// UpdateStockNotOne implements RepositoryTicket.
func (repo *repositoryTicket) UpdateStockNotOne(tx *sql.Tx, ticketId string, howManyUpdateStock int) (domain.Ticket, error) {
	var wg sync.WaitGroup
	var ticket domain.Ticket
	var err error

	mtx := &repo.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mtx.Lock()
		defer mtx.Unlock()

		_, err = tx.Exec("UPDATE ticket SET stock = stock - $2 WHERE id = $1", ticketId, howManyUpdateStock)
		if err != nil {
			log.Info().Any("ERROR at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]: %v", err)
			return
		}
		log.Info().Any("SUCCESS at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]", "").Msg("")
	}()

	wg.Wait()

	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
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

	var wg sync.WaitGroup
	var ticket domain.Ticket
	var err error

	mtx := &repo.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mtx.Lock()
		defer mtx.Unlock()

		_, err = tx.Exec("UPDATE ticket SET stock = stock - 1 WHERE id = $1", ticketId)
		if err != nil {
			log.Info().Any("ERROR at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]: %v", err)
			return
		}
		log.Info().Any("SUCCESS at [REPOSITORY] - [TICKET] - [UpdateStock] - [set update data query]", "").Msg("")
	}()

	wg.Wait()

	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil

}

// FindByName implements RepositoryTicket.
func (repo *repositoryTicket) FindByEventName(tx *sql.Tx, ticketName string) ([]domain.Ticket, error) {
	var allTicket []domain.Ticket

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

	return allTicket, nil

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
		database: database,
	}
}
