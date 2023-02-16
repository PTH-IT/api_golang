package firebasedb

import (
	"PTH-IT/api_golang/domain/repository"
	"context"
	"fmt"
)

func NewFirebaseRepository() repository.FirebaseRepository {
	return firebaseRepository{}
}

type firebaseRepository struct {
}

func (r firebaseRepository) Getfirebase() ([]map[string]interface{}, error) {
	ctx := context.Background()

	client, err := Connect(ctx)
	if err != nil {
		return nil, err
	}
	// The app only has access as defined in the Security Rules
	ref := client.NewRef("user_scores")
	var get []map[string]interface{}
	if err := ref.Get(context.TODO(), &get); err != nil {
		return get, err
	}

	return get, nil
}
func (r firebaseRepository) Putfirebase() error {
	ctx := context.Background()

	client, err := Connect(ctx)
	if err != nil {
		return err
	}

	// The app only has access as defined in the Security Rules
	ref := client.NewRef("user_scores/" + fmt.Sprint(1))

	if err := ref.Set(context.TODO(), map[string]interface{}{"score": 40}); err != nil {
		return err
	}

	return nil
}
