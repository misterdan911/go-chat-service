package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-chat-service/service/authservice"
	"os"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}
type ResponseFailed struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func JwtVerifier(c *fiber.Ctx) error {

	var isValid bool = true
	var errorMessage string
	auth := c.Get("Authorization")

	if auth == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get Authorization header",
		})
	}

	// hapus text 'Bearer '
	// extract token-nya aja
	var jwtToken string = strings.Replace(auth, "Bearer ", "", -1)

	// Parse the token
	token, err := jwt.ParseWithClaims(jwtToken, &authservice.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
		return jwtSecret, nil

	})

	// parsing errors result
	if err != nil {
		isValid = false
		errorMessage = "You're Unauthorized due to error parsing the JWT"
	}

	if !token.Valid {
		isValid = false
		errorMessage = "You're Unauthorized due to invalid token"
	}

	if isValid {

		// Extract the claims and the user ID
		if claims, ok := token.Claims.(*authservice.MyCustomClaims); ok {
			// Inject the user ID into the Fiber context
			c.Locals("login_username", claims.Username)
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
		}

		return c.Next()
	} else {
		return c.Status(fiber.StatusForbidden).SendString(errorMessage)
	}

}
