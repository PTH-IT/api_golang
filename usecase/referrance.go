package usecase

import (
	"fmt"

	"github.com/PTH-IT/api_golang/domain/model"
	"github.com/PTH-IT/api_golang/domain/repository"
	InforLog "github.com/PTH-IT/api_golang/log/infor"
)

type Reference interface {
	GetUserGormdb(userId string, password string) (*model.User, error)
	AddtUserGormdb(userId string, password string) error

	Getfirebase() ([]map[string]interface{}, error)
	Putfirebase() error

	GetMovies() ([]*model.Movies, error)
	AddMovies(name string, title string, location string) error
	GetUser(userId string, password string) (*model.GetUser, error)
	AddUser(userId string, password string) error
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

func (r reference) GetUserGormdb(userId string, password string) (*model.User, error) {

	user, err := r.userRepository.GetUser(userId, password)
	return user, err
}
func (r reference) AddtUserGormdb(userId string, password string) error {

	err := r.userRepository.AddUser(userId, password)
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

func (r reference) GetUser(userId string, password string) (*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	user, err := r.mongoRepository.GetUser(userId, password)
	return user, err
}
func (r reference) AddUser(userId string, password string) error {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.AddUser call"))
	err := r.mongoRepository.AddUser(userId, password)
	return err
}

func (r reference) GetMovies() ([]*model.Movies, error) {

	result, err := r.mongoRepository.Getmongo()
	return result, err
}
func (r reference) AddMovies(name string, title string, location string) error {

	err := r.mongoRepository.AddMovies(name, title, location)
	return err
}
