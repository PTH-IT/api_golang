package repository

import (
	"PTH-IT/api_golang/domain/model"
)

type MonggoRepository interface {
	GetUser(userId string, password string) (*model.GetUser, error)
	AddUser(userId string, password string) error
	Getmongo() ([]*model.Movies, error)
	AddMovies(name string, title string, location string) error
}
