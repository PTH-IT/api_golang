package repository

import (
	"github.com/PTH-IT/api_golang/domain/model"
)

type UserRepository interface {
	GetUser() (*model.User, error)
}
