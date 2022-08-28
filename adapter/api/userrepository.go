package api

import (
	"github.com/PTH-IT/api_golang/repository"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser() (string, error) {

	return "user", nil
}
