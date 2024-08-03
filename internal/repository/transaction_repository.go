package repository

import (
	"context"
	"database/sql"
	"excercise2/internal/domain"

	"github.com/rs/zerolog/log"
)

type RepositoryTransaction interface {
	CreateTransaction
	CreateTransactionWithTx
	FindAllTransaction
}

type CreateTransactionWithTx interface {
	CreateTransactionWithTx(ctx context.Context, tx *sql.Tx, request domain.Transaction) (domain.Transaction, error)
}

type CreateTransaction interface {
	Create(ctx context.Context, request domain.Transaction) (domain.Transaction, error)
}

type FindAllTransaction interface {
	FindAll(tx *sql.Tx) ([]domain.Transaction, error)
}

type repositoryTransaction struct {
	db *sql.DB
}

// CreateTransactionWithTx implements RepositoryTransaction.
func (repo *repositoryTransaction) CreateTransactionWithTx(ctx context.Context, tx *sql.Tx, request domain.Transaction) (domain.Transaction, error) {
	var transaction domain.Transaction

	err := repo.db.QueryRowContext(ctx, "INSERT INTO transactions(id, user_id, event_id, ticket_id, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, event_id, ticket_id, create_at, update_at", request.Id, request.User.Id, request.Event.Id, request.Ticket.Id, request.Create_at, request.Update_at).Scan(&transaction.Id, &transaction.User.Id, &transaction.Event.Id, &transaction.Ticket.Id, &transaction.Create_at, &transaction.Update_at)
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [TRANSACTION] - [CreateTransactionWithTx] - [create data to database]", err).Msg("")
		return domain.Transaction{}, err
	}

	return transaction, nil
}

// Create implements RepositoryTransaction.
func (repo *repositoryTransaction) Create(ctx context.Context, request domain.Transaction) (domain.Transaction, error) {

	var transaction domain.Transaction

	err := repo.db.QueryRowContext(ctx, "INSERT INTO transactions(id, user_id, event_id, ticket_id, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, event_id, ticket_id, create_at, update_at", request.Id, request.User.Id, request.Event.Id, request.Ticket.Id, request.Create_at, request.Update_at).Scan(&transaction.Id, &transaction.User.Id, &transaction.Event.Id, &transaction.Ticket.Id, &transaction.Create_at, &transaction.Update_at)
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [TRANSACTION] - [Create] - [create data to database]", err).Msg("")
		return domain.Transaction{}, err
	}

	return transaction, nil
}

// FindAll implements RepositoryTransaction.
func (repo *repositoryTransaction) FindAll(tx *sql.Tx) ([]domain.Transaction, error) {
	var allTransaction []domain.Transaction

	rows, err := tx.Query("SELECT id, user_id, event_id, ticket_id, create_at, update_at FROM transactions")
	if err != nil {
		log.Info().Any("ERROR at [REPOSITORY] - [TRANSACTION] - [FindAll] - [get data from database query]", err).Msg("")
		return []domain.Transaction{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(&transaction.Id, &transaction.User.Id, &transaction.Event.Id, &transaction.Ticket.Id, &transaction.Create_at, &transaction.Update_at)
		if err != nil {
			log.Info().Any("ERROR at [REPOSITORY] - [TRANSACTION] - [FindAll] - [rows scan data]", err).Msg("")
			return []domain.Transaction{}, err
		}
		allTransaction = append(allTransaction, transaction)

	}

	return allTransaction, nil

}

func NewRepositoryTransaction(db *sql.DB) RepositoryTransaction {
	return &repositoryTransaction{
		db: db,
	}
}
