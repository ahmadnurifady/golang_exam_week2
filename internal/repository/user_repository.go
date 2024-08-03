package repository

import (
	"database/sql"
	"excercise2/internal/domain"

	"github.com/rs/zerolog/log"
)

type RepositoryUser interface {
	CreateUser
	FindAllUser
	FindByIdUser
}

type CreateUser interface {
	Create(request domain.User) (domain.User, error)
}

type FindAllUser interface {
	FindAll() ([]domain.User, error)
}

type FindByIdUser interface {
	FindById(tx *sql.Tx, userId string) (domain.User, error)
}

type repositoryUser struct {
	database *sql.DB
}

// Create implements RepositoryUser.
func (repo *repositoryUser) Create(request domain.User) (domain.User, error) {
	var user domain.User

	err := repo.database.QueryRow("INSERT INTO users (id, name) VALUES ($1, $2)", request.Id, request.Name).Scan(&user.Id, &user.Name)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// FindAll implements RepositoryUser.
func (repo *repositoryUser) FindAll() ([]domain.User, error) {
	var allUsers []domain.User

	rows, err := repo.database.Query("SELECT id, name FROM users")
	if err != nil {
		return []domain.User{}, err
	}

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return []domain.User{}, err
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil

}

// FindById implements RepositoryUser.
func (repo *repositoryUser) FindById(tx *sql.Tx, userId string) (domain.User, error) {

	var user domain.User

	err := tx.QueryRow("SELECT id , name FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name)
	if err != nil {

		log.Info().Any("ERROR at [REPOSITORY] - [USER] - [FindById] - [get data from database]", err).Msg("")

		return domain.User{}, err
	}

	return user, nil

}

func NewRepositoryUser(database *sql.DB) RepositoryUser {
	return &repositoryUser{
		database: database,
	}
}
