package gormdb

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser() (*model.User, error) {

	var user *model.User
	repo.Begin().Table("user").Find(&user)
	return user, nil
}
