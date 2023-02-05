package repository

import (
	"PTH-IT/api_golang/domain/model"
)

type UserRepository interface {
	GetUser() (*model.User, error)
}
