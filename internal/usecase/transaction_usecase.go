package usecase

import (
	"context"
	"database/sql"
	"excercise2/internal/domain"
	"excercise2/internal/domain/dto"
	"excercise2/internal/repository"
	"fmt"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type UsecaseTransaction interface {
	UsecaseCreateTransaction
	UsecaseFindAllTransaction
}

type UsecaseCreateTransaction interface {
	CreateUsecase(ctx context.Context, request dto.CreateTxDto) (domain.Transaction, error)
}

type UsecaseFindAllTransaction interface {
	FindAllUsecase(ctx context.Context) ([]domain.Transaction, error)
}

type usecaseTransaction struct {
	repo       repository.RepositoryTransaction
	repoUser   repository.RepositoryUser
	repoEvent  repository.RepositoryEvent
	repoTicket repository.RepositoryTicket
	sync.WaitGroup
	database *sql.DB
}

// CreateUsecase implements UsecaseTransaction.
func (uc *usecaseTransaction) CreateUsecase(ctx context.Context, request dto.CreateTxDto) (domain.Transaction, error) {

	tx, err := uc.database.Begin()
	if err != nil {
		return domain.Transaction{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var createTx domain.Transaction
	uuidGenerate, _ := uuid.NewV4()

	user, err := uc.repoUser.FindById(tx, request.UserId)
	if err != nil {
		log.Info().Any("ERROR at [USECASE] - [TRANSACTION] - [CreateUsecase] - [tx begin]", err).Msg("")
		return domain.Transaction{}, err
	}

	event, err := uc.repoEvent.FindById(tx, ctx, request.EventId)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("event dengan id: %s tidak ditemukan", request.EventId)
	}

	allTicket, err := uc.repoTicket.FindByEventName(tx, event.EventName)
	if err != nil {
		return domain.Transaction{}, err
	}

	createTx.Id = uuidGenerate.String()
	createTx.User = user
	createTx.Event = event
	for _, ticket := range allTicket {
		if ticket.Type == request.TicketType {
			createTx.Ticket = ticket
			_, err := uc.repoTicket.UpdateStock(tx, ticket.Id)
			if err != nil {
				log.Info().Any("ERROR at [USECASE] - [TRANSACTION] - [CreateUsecase] - [update stock ticket]", err).Msg("")
				return domain.Transaction{}, err
			}
		}
	}
	fmt.Println("ticket user == ", createTx.Ticket)

	createTx.Create_at = time.Now()
	createTx.Update_at = time.Now()

	_, err = uc.repo.Create(ctx, createTx)
	if err != nil {
		return domain.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.Transaction{}, err
	}

	return createTx, nil
}

// FindAllUsecase implements UsecaseTransaction.
func (uc *usecaseTransaction) FindAllUsecase(ctx context.Context) ([]domain.Transaction, error) {
	var allTransaction []domain.Transaction

	tx, err := uc.database.Begin()
	if err != nil {
		return []domain.Transaction{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	findAll, err := uc.repo.FindAll(tx)
	if err != nil {
		return []domain.Transaction{}, err
	}

	for _, valTx := range findAll {
		findUser, err := uc.repoUser.FindById(tx, valTx.User.Id)
		if err != nil {
			log.Info().Any("ERROR at [USECASE] - [TRANSACTION] - [FindAllUsecase] - [get data from repository user]", err).Msg("")
			return []domain.Transaction{}, err
		}

		findEvent, err := uc.repoEvent.FindById(tx, ctx, valTx.Event.Id)
		if err != nil {
			log.Info().Any("ERROR at [USECASE] - [TRANSACTION] - [FindAllUsecase] - [get data from repository event]", err).Msg("")
			return []domain.Transaction{}, err
		}

		findTicket, err := uc.repoTicket.FindById(tx, valTx.Ticket.Id)
		if err != nil {
			log.Info().Any("ERROR at [USECASE] - [TRANSACTION] - [FindAllUsecase] - [get data from repository ticket]", err).Msg("")
			return []domain.Transaction{}, err
		}

		valTx.User = findUser
		valTx.Event = findEvent
		valTx.Ticket = findTicket

		allTransaction = append(allTransaction, valTx)
	}

	fmt.Println("all transaction in usecase == ", allTransaction)

	return allTransaction, err
}

func NewUsecaseTransaction(repo repository.RepositoryTransaction, repoUser repository.RepositoryUser, repoEvent repository.RepositoryEvent, repoTicket repository.RepositoryTicket, database *sql.DB) UsecaseTransaction {
	return &usecaseTransaction{
		repo:       repo,
		repoUser:   repoUser,
		repoEvent:  repoEvent,
		database:   database,
		repoTicket: repoTicket,
	}
}
