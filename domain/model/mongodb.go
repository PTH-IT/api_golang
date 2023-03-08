package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movies struct {
	Name     string `bson:"name"`
	Title    string `bson:"title"`
	Location string `bson:"ocation"`
}

type GetUser struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserID       string             `bson:"userid"`
	Password     string             `bson:"password"`
	Email        string             `bson:"email"`
	ActiceEmail  string             `bson:"actice"`
	Status       string             `bson:"status"`
	CreatedTime  string             `bson:"createdtime"`
	UpdatedTime  string             `bson:"updatedtime"`
	ConnectionId string             `bson:"connectionid"`
}
type RegisterUserMonggo struct {
	UserID       string `bson:"userid"`
	Password     string `bson:"password"`
	Email        string `bson:"email"`
	ActiceEmail  string `bson:"actice"`
	Status       string `bson:"status"`
	CreatedTime  string `bson:"createdtime"`
	UpdatedTime  string `bson:"updatedtime"`
	ConnectionId string `bson:"connectionid"`
}
type MessageCheckUser struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
