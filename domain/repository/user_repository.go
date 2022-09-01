package repository

type UserRepository interface {
	GetUser() (string, error)
}
