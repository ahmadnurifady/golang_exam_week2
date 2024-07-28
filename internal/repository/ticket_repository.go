package repository

import (
	"excercise2/internal/domain"
	"fmt"
)

type RepositoryTicket interface {
	FindAll()
	FindById(ticketId int)
	FindByName(ticketName string)
	UpdateStock(ticketId int) (*domain.Ticket, error)
}

type repositoryTicket struct {
	db map[int]domain.Ticket
}

// UpdateStock implements RepositoryTicket.
func (repo *repositoryTicket) UpdateStock(ticketId int) (*domain.Ticket, error) {
	existingTicket, exists := repo.db[ticketId]
	if !exists {
		return &domain.Ticket{}, fmt.Errorf("ticket dengan id: %d tidak ditemukan", ticketId)
	}

	existingTicket.Stock -= 1

	repo.db[ticketId] = existingTicket
	return &existingTicket, nil

}

// FindByName implements RepositoryTicket.
func (repo *repositoryTicket) FindByName(ticketName string) {
	panic("unimplemented")
}

// FindAll implements RepositoryTicket.
func (repo *repositoryTicket) FindAll() {
	panic("unimplemented")
}

// FindById implements RepositoryTicket.
func (repo *repositoryTicket) FindById(ticketId int) {
	panic("unimplemented")
}

func NewRepositoryTicket() RepositoryTicket {
	return &repositoryTicket{
		db: make(map[int]domain.Ticket),
	}
}
