package api

import (
	"github.com/PTH-IT/api_golang/domain/model"
	"github.com/PTH-IT/api_golang/domain/repository"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser() (*model.User, error) {

	iptuser := &model.User{
		UserName: "admin",
		Password: "admin",
	}
	return iptuser, nil
}
