package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author  primitive.ObjectID `json:"author" bson:"author"`
	Content string             `json:"content" bson:"content"`
	Room    primitive.ObjectID `json:"room" bson:"room"`
	Date    time.Time          `json:"date" bson:"date"`
}
