package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Room struct {
	ID          primitive.ObjectID   `bson:"_id" json:"_id"`
	People      []primitive.ObjectID `bson:"people" json:"people"`
	Title       string               `bson:"title" json:"title"`
	LastUpdate  time.Time            `bson:"lastUpdate" json:"lastUpdate"`
	LastMessage primitive.ObjectID   `bson:"lastMessage" json:"lastMessage"`
	IsGroup     bool                 `bson:"isGroup" json:"isGroup"`
}

type RoomData struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	People      []UserData         `bson:"people" json:"people"`
	Title       string             `bson:"title" json:"title"`
	LastUpdate  time.Time          `bson:"lastUpdate" json:"lastUpdate"`
	LastMessage Message            `bson:"lastMessage" json:"lastMessage"`
	IsGroup     bool               `bson:"isGroup" json:"isGroup"`
}
