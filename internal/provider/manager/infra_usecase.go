package manager

import (
	"excercise2/internal/provider/db"
	"excercise2/internal/usecase"
)

type UsecaseManager interface {
	EventUc() usecase.UsecaseEvent
	TransactionUc() usecase.UsecaseTransaction
}

type usecaseManager struct {
	repo     RepoManager
	database db.Connection
}

// EventUc implements UsecaseManager.
func (u *usecaseManager) EventUc() usecase.UsecaseEvent {
	return usecase.NewUsecaseEvent(u.repo.EventRepo(), u.repo.TicketRepo(), u.repo.CompositeRepo(), u.database.GetDb())
}

// TransactionUc implements UsecaseManager.
func (u *usecaseManager) TransactionUc() usecase.UsecaseTransaction {
	return usecase.NewUsecaseTransaction(u.repo.TransactionRepo(), u.repo.UserRepo(), u.repo.EventRepo(), u.repo.TicketRepo(), u.database.GetDb())
}

func NewUcManager(repo RepoManager, database db.Connection) UsecaseManager {
	return &usecaseManager{repo: repo, database: database}
}
