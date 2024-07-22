package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Shield     string             `bson:"shield"`
	Name       string             `bson:"name"`
	Author     primitive.ObjectID `bson:"author"`
	Size       int64              `bson:"size"`
	Location   string             `bson:"location"`
	ShieldedID string             `bson:"shieldedID"`
}
