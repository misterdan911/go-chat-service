package authservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-chat-service/db"
	"go-chat-service/dto"
	"go-chat-service/model"
	"go-chat-service/orm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func AddNewUser(user *model.User) error {

	user.Password, _ = HashPassword(user.Password)

	user.ID = primitive.NewObjectID()
	insertOneResult, err := db.DB.Collection("users").InsertOne(context.Background(), user)

	fmt.Println(insertOneResult)

	// Check for errors
	if err != nil {
		return errors.New("Failed creating new user, " + err.Error())
	} else {
		return nil
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Custom validation function
func EmailUnique(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var user model.User

	result := orm.DB.Find(&user, "email = ?", email)

	if result.RowsAffected == 0 {
		return true
	} else {
		return false
	}
}

func UsernameUnique(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	var user model.User

	result := orm.DB.Find(&user, "username = ?", username)

	if result.RowsAffected == 0 {
		return true
	} else {
		return false
	}
}

func GenerateJWT(user *model.User) (string, error) {

	// get Picture data of the user
	var image model.Image
	err := db.DB.Collection("images").FindOne(context.Background(), bson.M{"_id": user.Picture}).Decode(&image)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"id":        user.ID,
		"email":     user.Email,
		"level":     user.Level,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"username":  user.Username,
		"picture":   image,
	})

	// Sign and get the complete encoded token as a string using the secret
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func ValidateSignIn(signedInUser *dto.SignedInUser) (bool, model.User, string, string) {

	var isValidUser bool
	var user model.User
	var jwtToken string = ""

	err := db.DB.Collection("users").FindOne(context.Background(), bson.M{"username": signedInUser.Email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = db.DB.Collection("users").FindOne(context.Background(), bson.M{"email": signedInUser.Email}).Decode(&user)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					isValidUser = false
					return isValidUser, user, jwtToken, "email"
				}
			}
		}
	}

	/*
		passwordMatch := CheckPasswordHash(signedInUser.Password, user.Password)

		if !passwordMatch {
			isValidUser = false
			return isValidUser, user, jwtToken, "password"
		}
	*/

	jwtToken, errJwt := GenerateJWT(&user)

	if errJwt != nil {
		log.Fatal("Error generating JWT", err)
	}

	isValidUser = true
	return isValidUser, user, jwtToken, ""
}
