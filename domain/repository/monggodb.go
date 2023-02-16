package repository

import (
	"PTH-IT/api_golang/domain/model"
)

type MonggoRepository interface {
	Getmongo() ([]*model.Movies, error)
	Putmongo() error
}
