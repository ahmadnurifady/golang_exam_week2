package usecase

import (
	"excercise2/internal/domain"
	"excercise2/internal/domain/dto"
	"excercise2/internal/repository"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type UsecaseTransaction interface {
	UsecaseCreateTransaction
	UsecaseFindAllTransaction
}

type UsecaseCreateTransaction interface {
	CreateUsecase(request dto.CreateTxDto) (*domain.Transaction, error)
}

type UsecaseFindAllTransaction interface {
	FindAllUsecase() ([]domain.Transaction, error)
}

type usecaseTransaction struct {
	repo      repository.RepositoryTransaction
	repoUser  repository.RepositoryUser
	repoEvent repository.RepositoryEvent
}

// CreateUsecase implements UsecaseTransaction.
func (uc *usecaseTransaction) CreateUsecase(request dto.CreateTxDto) (*domain.Transaction, error) {

	var createTx domain.Transaction
	uuidGenerate, _ := uuid.NewV4()

	user, err := uc.repoUser.FindById(request.UserId)
	if err != nil {
		return &domain.Transaction{}, fmt.Errorf("user dengan id: %s tidak ditemukan", request.UserId)
	}

	event, err := uc.repoEvent.FindById(request.EventId)
	if err != nil {
		return &domain.Transaction{}, fmt.Errorf("event dengan id: %s tidak ditemukan", request.EventId)
	}

	createTx.Id = uuidGenerate.String()
	createTx.User = user
	createTx.EventName = event.EventName
	for _, ticket := range event.Ticket {
		if ticket.Type == request.TicketType {
			createTx.Ticket.Type = ticket.Type
			createTx.Ticket.Price = ticket.Price
			uc.repoEvent.UpdateEventTicketStock(request.EventId, request.TicketType)
		}
	}
	createTx.Create_at = time.Now()
	createTx.Update_at = time.Now()

	result, err := uc.repo.Create(&createTx)
	if err != nil {
		return &domain.Transaction{}, err
	}

	return result, nil
}

// FindAllUsecase implements UsecaseTransaction.
func (uc *usecaseTransaction) FindAllUsecase() ([]domain.Transaction, error) {
	findAll, err := uc.repo.FindAll()
	if err != nil {
		return []domain.Transaction{}, err
	}
	return findAll, err
}

func NewUsecaseTransaction(repo repository.RepositoryTransaction, repoUser repository.RepositoryUser, repoEvent repository.RepositoryEvent) UsecaseTransaction {
	return &usecaseTransaction{
		repo:      repo,
		repoUser:  repoUser,
		repoEvent: repoEvent,
	}
}
