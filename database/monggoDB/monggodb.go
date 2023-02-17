package monggodb

import (
	"PTH-IT/api_golang/config"
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func NewMongoDriver() repository.MonggoRepository {
	return MongoDriverRepository{}
}

type MongoDriverRepository struct {
}

func (r MongoDriverRepository) Getmongo() ([]*model.Movies, error) {
	var movies []*model.Movies
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("movies")
	var result []bson.M

	s, err := coll.Find(context.TODO(), bson.D{})
	s.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonToByte, &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil

}
func (r MongoDriverRepository) Putmongo() error {

	client, err := Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("movies")
	movies := &model.Movies{
		Name:     "tesst",
		Title:    "tesst",
		Location: "tesst",
	}

	_, err = coll.InsertOne(context.TODO(), movies)
	if err != nil {
		return err
	}
	return nil

}
