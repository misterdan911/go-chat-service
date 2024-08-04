package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
type User struct {
	ID         string   `json:"id" gorm:"primaryKey"`
	Email      string   `json:"email" validate:"required,email,email_unique"`
	FirstName  string   `json:"firstName" validate:"required"`
	Level      string   `json:"level" validate:"required"`
	Password   string   `json:"password,omitempty" validate:"required"`
	Phone      string   `json:"phone,omitempty"`
	LastName   string   `json:"lastName"`
	Username   string   `json:"username" validate:"required,username_unique"`
	Favorites  []string `json:"favorites"`
	Tagline    string   `json:"tagline"`
	LastOnline string   `json:"lastOnline"`
}
*/

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Email      string             `bson:"email" json:"email"`
	FirstName  string             `bson:"firstName" json:"firstName"`
	Level      string             `bson:"level" json:"level"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	Phone      string             `bson:"phone" json:"phone"`
	LastName   string             `bson:"lastName" json:"lastName"`
	Username   string             `bson:"username" json:"username"`
	Picture    primitive.ObjectID `bson:"picture" json:"picture"`
	Favorites  []string           `bson:"favorites" json:"favorites"`
	TagLine    string             `bson:"tagLine" json:"tagLine"`
	LastOnline time.Time          `bson:"lastOnline" json:"lastOnline"`
}

type UserData struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Email      string             `bson:"email" json:"email"`
	FirstName  string             `bson:"firstName" json:"firstName"`
	Level      string             `bson:"level" json:"level"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	Phone      string             `bson:"phone" json:"phone"`
	LastName   string             `bson:"lastName" json:"lastName"`
	Username   string             `bson:"username" json:"username"`
	Picture    Image              `bson:"picture" json:"picture"`
	Favorites  []string           `bson:"favorites" json:"favorites"`
	TagLine    string             `bson:"tagLine" json:"tagLine"`
	LastOnline time.Time          `bson:"lastOnline" json:"lastOnline"`
}
