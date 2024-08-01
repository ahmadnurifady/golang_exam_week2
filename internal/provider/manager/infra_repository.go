package manager

import (
	"excercise2/internal/provider/db"
	"excercise2/internal/repository"
)

type RepoManager interface {
	UserRepo() repository.RepositoryUser
	TicketRepo() repository.RepositoryTicket
	EventRepo() repository.RepositoryEvent
	TransactionRepo() repository.RepositoryTransaction
	CompositeRepo() repository.CompositeEventTicket
}

type repoManager struct {
	database db.Connection
}

// CompositeRepo implements RepoManager.
func (r *repoManager) CompositeRepo() repository.CompositeEventTicket {
	return repository.NewRepositoryCompositeEventTicket(r.database.GetDb())
}

// EventRepo implements RepoManager.
func (r *repoManager) EventRepo() repository.RepositoryEvent {
	return repository.NewRepositoryEvent(r.database.GetDb())
}

// TicketRepo implements RepoManager.
func (r *repoManager) TicketRepo() repository.RepositoryTicket {
	return repository.NewRepositoryTicket(r.database.GetDb())
}

// TransactionRepo implements RepoManager.
func (r *repoManager) TransactionRepo() repository.RepositoryTransaction {
	return repository.NewRepositoryTransaction(r.database.GetDb())
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.RepositoryUser {
	return repository.NewRepositoryUser(r.database.GetDb())
}

func NewRepoManager(database db.Connection) RepoManager {
	return &repoManager{database: database}
}
