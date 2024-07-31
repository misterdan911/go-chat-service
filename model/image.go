package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Shield     string             `bson:"shield" json:"shield"`
	Name       string             `bson:"name" json:"name"`
	Author     primitive.ObjectID `bson:"author" json:"author"`
	Size       int64              `bson:"size" json:"size"`
	Location   string             `bson:"location" json:"location"`
	ShieldedID string             `bson:"shieldedID" json:"shieldedID"`
}
