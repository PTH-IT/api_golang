package repository

import (
	"PTH-IT/api_golang/domain/model"
)

type MonggoRepository interface {
	GetUser(userId string, password string) (*model.GetUser, error)
	CheckUserName(userId string, email string) ([]*model.GetUser, error)
	AddUser(userId string, password string, email string) error
	Getmongo() ([]*model.Movies, error)
	AddMovies(name string, title string, location string) error
	GetConnectionID(userId string) (*model.GetUser, error)
	UpdateConnectionID(userId string, connectionid string) error
}
