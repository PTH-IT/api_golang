package monggodb

import (
	"PTH-IT/api_golang/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client ,error) {
	uri := fmt.Sprintf(config.Getconfig().Monggo.Host, config.Getconfig().Monggo.User, config.Getconfig().Monggo.Pass)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	
	return client, err 

}
