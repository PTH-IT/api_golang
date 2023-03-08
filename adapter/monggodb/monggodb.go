package monggodb

import (
	"PTH-IT/api_golang/config"
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/domain/repository"
	errormessage "PTH-IT/api_golang/log/error"
	"PTH-IT/api_golang/utils"
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
func (r MongoDriverRepository) AddMovies(name string, title string, location string) error {

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
		Name:     name,
		Title:    title,
		Location: location,
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

	filter := bson.D{{Key: "userid", Value: userId}, {Key: "password", Value: password}}
	s, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = s.All(context.TODO(), &result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = json.Unmarshal(jsonToByte, &listUser)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	if len(listUser) == 0 {
		return nil, nil
	}

	return listUser[0], nil
}

func (r MongoDriverRepository) GetConnectionID(userId string) (*model.GetUser, error) {

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

	filter := bson.D{{Key: "userid", Value: userId}}
	s, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = s.All(context.TODO(), &result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = json.Unmarshal(jsonToByte, &listUser)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	if len(listUser) == 0 {
		return nil, nil
	}

	return listUser[0], nil
}

func (r MongoDriverRepository) UpdateConnectionID(userId string, connectionid string) error {

	var listUser []*model.GetUser

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
	var result []bson.M

	filter := bson.D{{Key: "userid", Value: userId}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "connectionid", Value: connectionid}}}}
	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errormessage.PrintError("1", err)
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return errormessage.PrintError("1", err)
	}
	err = json.Unmarshal(jsonToByte, &listUser)
	if err != nil {
		return errormessage.PrintError("1", err)
	}
	if len(listUser) == 0 {
		return nil
	}

	return nil
}
func (r MongoDriverRepository) CheckUserName(userId string, email string) ([]*model.GetUser, error) {
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

	filter := bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "userid", Value: userId}}, bson.D{{Key: "email", Value: email}}}}}
	s, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = s.All(context.TODO(), &result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = json.Unmarshal(jsonToByte, &listUser)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	if len(listUser) == 0 {
		return nil, nil
	}

	return listUser, nil
}
func (r MongoDriverRepository) AddUser(userId string, password string, email string) error {

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
	user := &model.RegisterUserMonggo{
		UserID:      userId,
		Password:    password,
		Email:       email,
		Status:      "0",
		ActiceEmail: "0",
		CreatedTime: utils.GettimeNow(),
		UpdatedTime: utils.GettimeNow(),
	}
	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil

}
func (r MongoDriverRepository) SaveMessage(message *model.Message) error {
	client, err := Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("message")
	_, err = coll.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil

}
func (r MongoDriverRepository) Getmessage(message *model.InputGetMessage) ([]*model.GetMessage, error) {

	var listmessage []*model.GetMessage

	client, err := Connect()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("message")
	var result []bson.M
	filter := bson.D{{Key: "$or",
		Value: bson.A{
			bson.D{{Key: "sender", Value: message.Sender}, {Key: "receiver", Value: message.Receiver}},
			bson.D{{Key: "sender", Value: message.Receiver}, {Key: "receiver", Value: message.Sender}}}}}
	s, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = s.All(context.TODO(), &result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	jsonToByte, err := json.Marshal(result)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}
	err = json.Unmarshal(jsonToByte, &listmessage)
	if err != nil {
		return nil, errormessage.PrintError("1", err)
	}

	return listmessage, nil
}
