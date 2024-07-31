package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Kayanya gak kepake
type MyCustomClaims struct {
	UserId primitive.ObjectID `bson:"user_id" json:"user_id"`
	jwt.RegisteredClaims
}
