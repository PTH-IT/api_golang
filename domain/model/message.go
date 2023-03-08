package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Websocket struct {
	Type     string `json:"type"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Detail   string `json:"detail"`
}
type Message struct {
	Receiver string `bson:"receiver" json:"receiver"`
	Sender   string `bson:"sender" json:"sender"`
	Detail   string `bson:"detail" json:"detail"`
	Time     int64  `bson:"time" json:"time"`
}
type GetMessage struct {
	ID       primitive.ObjectID `bson:"_id"`
	Receiver string             `bson:"receiver" json:"receiver"`
	Sender   string             `bson:"sender" json:"sender"`
	Detail   string             `bson:"detail" json:"detail"`
	Time     int64              `bson:"time" json:"time"`
}
type InputGetMessage struct {
	Receiver string `bson:"receiver" json:"receiver"`
	Sender   string `bson:"sender" json:"sender"`
}
