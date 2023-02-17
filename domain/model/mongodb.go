package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	Name     string `bson:"name"`
	Title    string `bson:"title"`
	Location string `bson:"ocation"`
}

type GetUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      string             `bson:"userid"`
	Password    string             `bson:"password"`
	Status      string             `bson:"status"`
	CreatedTime string             `bson:"createdtime"`
	UpdatedTime string             `bson:"updatedtime"`
}
