package authservice

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-bootstrap/dto"
	"go-fiber-bootstrap/model"
	"go-fiber-bootstrap/orm"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func AddNewUser(user *model.User) error {

	user.Password, _ = HashPassword(user.Password)
	result := orm.DB.Create(user)

	// Check for errors
	if result.Error != nil {
		return errors.New("Failed creating new user, " + result.Error.Error())
	} else {
		return nil
	}
}

/*
func ValidateSignInnnn() error {
	var isValidUser bool

	// Create a new SignInData struct
	signInData := new(SignInData)

	// Parse the JSON request body into the signInData struct
	if err := c.BodyParser(signInData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse SignIn JSON",
		})
	}

	var user model.User
	err := orm.DB.First(&user, "email = ?", signInData.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		isValidUser = false
	}

	passwordIsSame := passwordIsSame(request.Password, users.Password)

	if !passwordIsSame {
		isValidUser = false
	}

}
*/

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

func GenerateJWT() (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func ValidateSignIn(signedInUser *dto.SignedInUser) (bool, model.User, string) {

	var isValidUser bool
	var user model.User
	var jwtToken string = ""

	err := orm.DB.First(&user, "username = ?", signedInUser.Username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = orm.DB.First(&user, "email = ?", signedInUser.Email).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isValidUser = false
			//validateError := errors.New("Invalid user or email")
			return isValidUser, user, jwtToken
		}
	}

	passwordMatch := CheckPasswordHash(signedInUser.Password, user.Password)

	if !passwordMatch {
		isValidUser = false
		//validateError := errors.New("SignIn failed. Please check your username or email and password")
		return isValidUser, user, jwtToken
	}

	jwtToken, errJwt := GenerateJWT()

	if errJwt != nil {
		log.Fatal("Error generating JWT", err)
	}

	isValidUser = true
	return isValidUser, user, jwtToken
}
