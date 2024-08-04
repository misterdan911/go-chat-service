package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Kayanya gak kepake
type MyCustomClaims struct {
	UserId primitive.ObjectID `bson:"_id" json:"id"`
	jwt.RegisteredClaims
}
