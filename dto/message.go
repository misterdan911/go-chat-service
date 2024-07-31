package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SocketMessage struct {
	RoomID      primitive.ObjectID `json:"roomID" bson:"roomID"`
	AuthorID    primitive.ObjectID `json:"authorID" bson:"authorID"`
	Content     string             `json:"content" bson:"content"`
	ContentType string             `json:"contentType" bson:"contentType"`
}
