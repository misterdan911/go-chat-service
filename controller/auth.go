package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-chat-service/dto"
	"go-chat-service/model"
	"go-chat-service/service/authservice"
	"log"
	"time"
)

type SignInSuccessResponse struct {
	Token string `json:"token"`
}

type SignInFailMessage struct {
	SignIn string `json:"signin"`
}

type UserNotFound struct {
	Email string `json:"email"`
}

type WrongPassword struct {
	Password string `json:"password"`
}

var validate = validator.New()

func init() {
	err := validate.RegisterValidation("email_unique", authservice.EmailUnique)
	if err != nil {
		log.Fatal("Failed to register custom validation 'email_unique'")
	}

	err2 := validate.RegisterValidation("username_unique", authservice.UsernameUnique)
	if err2 != nil {
		log.Fatal("Failed to register custom validation 'email_unique'")
	}
}

// SignUp godoc
//
//	@Summary		Signing up new user
//	@Description	Signing up new user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	body		ExampleSignedUpUser	true	"Add user"
//	@Success		200		{object}	AppResponse
//	@Failure		400		{object}	ExampleSignedUpUser
//	@Router			/api/register [post]
func SignUp(c *fiber.Ctx) error {

	// Create a new User struct
	userSignUp := new(dto.SignedUpUser)

	// Parse the JSON request body into the user struct
	if err := c.BodyParser(userSignUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Validate the user struct
	/*
		if err := validate.Struct(user); err != nil {
			// Format validation errors
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				var message string

				switch err.Tag() {
				case "email_unique":
					message = "Email has already been registered"
				case "username_unique":
					message = "Username has already been registered"
				default:
					//message = fmt.Sprintf("Field '%s' is invalid", err.Field())
					message = fmt.Sprintf("Field '%s' %s", err.Field(), err.Tag())
				}

				errors[err.Field()] = message
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "fail",
				"data":   errors,
			})
		}
	*/

	// Create
	user := new(model.User)
	user.Email = userSignUp.Email
	user.FirstName = userSignUp.FirstName
	user.Level = "standard"
	user.Password = userSignUp.Password
	user.Phone = ""
	user.LastName = userSignUp.LastName
	user.Username = userSignUp.Username
	user.Favorites = make([]string, 0)
	user.TagLine = "New Clover User"
	user.LastOnline = time.Now()

	err := authservice.AddNewUser(user)
	user.Password = ""
	var response interface{}

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		response = ResponseError{
			Status:  "error",
			Message: err.Error(),
		}
	} else {
		response = user
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)
}

// SignIn godoc
//
//	@Summary		Signing in user
//	@Description	Signing in user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	body		ExampleSignedInUser	true	"SignIn User"
//	@Success		200		{object}	AppResponse
//	@Router			/api/login [post]
func SignIn(c *fiber.Ctx) error {

	signedInUser := new(dto.SignedInUser)

	if err := c.BodyParser(signedInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse SignIn JSON",
		})
	}

	isValid, userData, jwt, invalidType := authservice.ValidateSignIn(signedInUser)
	userData.Password = ""

	var response interface{}

	if isValid {
		response = SignInSuccessResponse{Token: jwt}
	} else {
		if invalidType == "email" {
			response = UserNotFound{Email: "User not Found"}
		} else {
			response = WrongPassword{Password: "Wrong Password"}
		}
	}

	// supaya field2 response json nya sesuai urutan kita
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// supaya response headernya 'application/json'
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)
}
