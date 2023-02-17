package monggodb

import (
	"PTH-IT/api_golang/config"
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
	"context"
	"encoding/json"
	"time"

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

func (r MongoDriverRepository) GetUser(userId string, password string) (*model.GetUser, error) {

	var listUser []*model.GetUser

	client, err := Connect()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("users")
	var result []bson.M

	filter := bson.D{{"userid", userId}, {"password", password}}
	s, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	s.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonToByte, &listUser)
	if err != nil {
		return nil, err
	}
	if len(listUser) == 0 {
		return nil, err
	}

	return listUser[0], nil
}

func (r MongoDriverRepository) AddUser(userId string, password string) error {

	client, err := Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("users")
	user := &model.User{
		UserID:      userId,
		Password:    password,
		Status:      "0",
		CreatedTime: time.Now().UTC().Local().Format(time.RFC3339Nano),
		UpdatedTime: time.Now().UTC().Local().Format(time.RFC3339Nano),
	}
	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil

}
