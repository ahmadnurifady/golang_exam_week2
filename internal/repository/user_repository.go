package repository

import (
	"excercise2/internal/domain"
	"fmt"
)

type RepositoryUser interface {
	CreateUser
	FindAllUser
	FindByIdUser
	FindByNameUser
}

type CreateUser interface {
	Create(request *domain.User) (*domain.User, error)
}

type FindAllUser interface {
	FindAll() ([]domain.User, error)
}

type FindByIdUser interface {
	FindById(userId string) (domain.User, error)
}

type FindByNameUser interface {
	FindByName(userName string) (domain.User, error)
}

type repositoryUser struct {
	db map[string]domain.User
}

// Create implements RepositoryUser.
func (repo *repositoryUser) Create(request *domain.User) (*domain.User, error) {
	var user domain.User
	user.Id = request.Id
	user.Name = request.Name

	repo.db[request.Id] = user
	return request, nil
}

// FindAll implements RepositoryUser.
func (repo *repositoryUser) FindAll() ([]domain.User, error) {
	var allUsers []domain.User

	for _, user := range repo.db {
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}

// FindById implements RepositoryUser.
func (repo *repositoryUser) FindById(userId string) (domain.User, error) {

	var findUser domain.User

	if _, exist := repo.db[userId]; !exist {
		return domain.User{}, fmt.Errorf("user dengan id: %s tidak ditemukan", userId)
	}

	for _, user := range repo.db {
		if user.Id == userId {
			findUser = user
		}
	}

	return findUser, nil

}

// FindByName implements RepositoryUser.
func (repo *repositoryUser) FindByName(userName string) (domain.User, error) {
	var findUser domain.User

	for _, user := range repo.db {
		if user.Name == userName {
			findUser = user
		}
	}

	if findUser.Id == "" || findUser.Name == "" {
		return domain.User{}, fmt.Errorf("user dengan name: %s tidak ditemukan", userName)
	}

	return findUser, nil
}

func NewRepositoryUser() RepositoryUser {
	return &repositoryUser{
		db: make(map[string]domain.User),
	}
}
