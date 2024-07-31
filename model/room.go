package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Room struct {
	ID          primitive.ObjectID   `bson:"_id" json:"_id"`
	People      []primitive.ObjectID `bson:"people" json:"people"`
	LastUpdate  time.Time            `bson:"lastUpdate" json:"lastUpdate"`
	LastMessage primitive.ObjectID   `bson:"lastMessage" json:"lastMessage"`
	IsGroup     bool                 `bson:"isGroup" json:"isGroup"`
}

type RoomData struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	People      []User             `bson:"people" json:"people"`
	LastUpdate  time.Time          `bson:"lastUpdate" json:"lastUpdate"`
	LastMessage primitive.ObjectID `bson:"lastMessage" json:"lastMessage"`
	IsGroup     bool               `bson:"isGroup" json:"isGroup"`
}
