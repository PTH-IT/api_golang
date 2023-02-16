package repository

type FirebaseRepository interface {
	Getfirebase() ([]map[string]interface{}, error)
	Putfirebase() error
}
