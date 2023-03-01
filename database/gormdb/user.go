package gormdb

import (
	"time"

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
	DB.Table("user").Where("user_id  = ? and password = ?", userId, password).Find(&user)
	if len(user) == 0 {
		return nil, nil
	}
	return user[0], nil
}
func (repo userRepository) AddUser(userId string, password string) error {
	var err error
	user := &model.User{
		UserID:      userId,
		Password:    password,
		Status:      "0",
		CreatedTime: time.Now().UTC().Local().Format(time.RFC3339Nano),
		UpdatedTime: time.Now().UTC().Local().Format(time.RFC3339Nano),
	}
	err = DB.Table("user").Create(user).Error

	if err != nil {
		return err
	}
	return nil
}
