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

func (repo userRepository) GetUser(userId string, password string) (*model.User, error) {

	var user []*model.User
	repo.Begin().Table("user").Where("UserID  = ? and Password = ?", userId, password).Find(&user)
	if len(user) == 0 {
		return nil, nil
	}
	return user[0], nil
}
