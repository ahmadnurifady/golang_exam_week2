package repository

import (
	"excercise2/internal/domain"
)

type RepositoryTransaction interface {
	CreateTransaction
	FindAllTransaction
}

type CreateTransaction interface {
	Create(request *domain.Transaction) (*domain.Transaction, error)
}

type FindAllTransaction interface {
	FindAll() ([]domain.Transaction, error)
}

type repositoryTransaction struct {
	db map[string]domain.Transaction
}

// Create implements RepositoryTransaction.
func (repo *repositoryTransaction) Create(request *domain.Transaction) (*domain.Transaction, error) {

	repo.db[request.Id] = *request
	return request, nil
}

// FindAll implements RepositoryTransaction.
func (repo *repositoryTransaction) FindAll() ([]domain.Transaction, error) {
	var allTransaction []domain.Transaction

	for _, tx := range repo.db {
		allTransaction = append(allTransaction, tx)
	}

	return allTransaction, nil
}

func NewRepositoryTransaction() RepositoryTransaction {
	return &repositoryTransaction{
		db: make(map[string]domain.Transaction),
	}
}
