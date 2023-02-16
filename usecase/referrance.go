package usecase

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
)

type Reference interface {
	GetUser(userId string, password string) (*model.User, error)
	AddtUser(userId string, password string) error
	Getmongo() ([]*model.Movies, error)
	Putmongo() error
	Getfirebase() ([]map[string]interface{}, error)
	Putfirebase() error
}

func NewReferrance(
	userRepository repository.UserRepository,
	mongoRepository repository.MonggoRepository,
	firebaseRepository repository.FirebaseRepository,
) Reference {
	return reference{
		userRepository,
		mongoRepository,
		firebaseRepository,
	}
}

type reference struct {
	userRepository     repository.UserRepository
	mongoRepository    repository.MonggoRepository
	firebaseRepository repository.FirebaseRepository
}

func (r reference) GetUser(userId string, password string) (*model.User, error) {

	user, err := r.userRepository.GetUser(userId, password)
	return user, err
}
func (r reference) AddtUser(userId string, password string) error {

	err := r.userRepository.AddUser(userId, password)
	return err
}

func (r reference) Getmongo() ([]*model.Movies, error) {

	result, err := r.mongoRepository.Getmongo()
	return result, err
}
func (r reference) Putmongo() error {

	err := r.mongoRepository.Putmongo()
	return err
}

func (r reference) Getfirebase() ([]map[string]interface{}, error) {

	result, err := r.firebaseRepository.Getfirebase()
	return result, err
}
func (r reference) Putfirebase() error {

	err := r.firebaseRepository.Putfirebase()
	return err
}
