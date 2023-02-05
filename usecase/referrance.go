package usecase

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
)

type Reference interface {
	GetUser(userId string, password string) (*model.User, error)
}

func NewReferrance(
	userRepository repository.UserRepository,
) Reference {
	return reference{
		userRepository,
	}
}

type reference struct {
	userRepository repository.UserRepository
}

func (r reference) GetUser(userId string, password string) (*model.User, error) {

	user, err := r.userRepository.GetUser(userId, password)
	return user, err
}
