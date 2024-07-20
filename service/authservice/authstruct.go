package authservice

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Username string `bson:"username" json:"username"`
	jwt.RegisteredClaims
}
