package usecase

import (
	"fmt"

	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
	InforLog "PTH-IT/api_golang/log/infor"
)

type Reference interface {
	GetUserGormdb(userId string, password string) (*model.User, error)
	AddtUserGormdb(userId string, password string) error

	Getfirebase() ([]map[string]interface{}, error)
	Putfirebase() error

	GetMovies() ([]*model.Movies, error)
	AddMovies(name string, title string, location string) error
	GetUser(userId string, password string) (*model.GetUser, error)
	CheckUserName(userId string, email string) ([]*model.GetUser, error)
	AddUser(userId string, password string, email string) error
	GetConnectionID(userId string) (*model.GetUser, error)
	UpdateConnectionID(userId string, connectionid string) error
	SaveMessage(message *model.Message) error
	Getmessage(message *model.InputGetMessage) ([]*model.GetMessage, error)
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
func (r reference) CheckUserName(userId string, email string) ([]*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	user, err := r.mongoRepository.CheckUserName(userId, email)
	return user, err
}
func (r reference) GetUser(userId string, password string) (*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	user, err := r.mongoRepository.GetUser(userId, password)
	return user, err
}

func (r reference) GetConnectionID(userId string) (*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	user, err := r.mongoRepository.GetConnectionID(userId)
	return user, err
}
func (r reference) UpdateConnectionID(userId string, connectionid string) error {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	err := r.mongoRepository.UpdateConnectionID(userId, connectionid)
	return err
}

func (r reference) AddUser(userId string, password string, email string) error {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.AddUser call"))
	err := r.mongoRepository.AddUser(userId, password, email)
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
func (r reference) SaveMessage(message *model.Message) error {

	err := r.mongoRepository.SaveMessage(message)
	return err
}
func (r reference) Getmessage(message *model.InputGetMessage) ([]*model.GetMessage, error) {
	result, err := r.mongoRepository.Getmessage(message)
	if err != nil {
		return nil, err
	}
	return result, nil
}
