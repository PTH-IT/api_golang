package api

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser(userId string, password string) (*model.User, error) {

	iptuser := &model.User{
		UserID:   "admin",
		Password: "admin",
	}
	return iptuser, nil
}
