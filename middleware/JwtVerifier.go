package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"

	"fmt"
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

	/*
		// Get all headers
		headers := c.GetReqHeaders()

		// Iterate over headers and print them
		for key, value := range headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	*/

	//auth := c.Request.Header["Authorization"]
	auth := c.Get("Authorization")

	if auth == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get Authorization header",
		})
	}

	// kalau tidak membawa jwt token
	//if len(auth) <= 0 {
	//	isValid = false
	//}

	// hapus text 'Bearer '
	// extract token-nya aja
	var jwtToken string = strings.Replace(auth, "Bearer ", "", -1)

	// verivy jwt token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
		return jwtSecret, nil
	})

	// parsing errors result
	if err != nil {
		/*
			writer.WriteHeader(http.StatusUnauthorized)
			_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
			if err2 != nil {
				return
			}
		*/
		isValid = false
		errorMessage = "You're Unauthorized due to error parsing the JWT"
	}

	if !token.Valid {
		isValid = false
		errorMessage = "You're Unauthorized due to invalid token"
	}

	if isValid {
		return c.Next()
	} else {
		/*
			var response ResponseFailed
			response.Code = 401
			response.Message = "failed"
			// response.Message = "Unauthorized"
			response.Message = errorMessage
			c.JSON(200, response)
			c.Abort()
			return
		*/

		//abortRequest(c, errorMessage)
		return c.Status(fiber.StatusForbidden).SendString(errorMessage)
	}

}

func abortRequest(ctx *fiber.Ctx, errorMessage string) {
	//var response ResponseFailed
	//response.Code = 401
	//response.Message = errorMessage
	//ctx.JSON(200, response)
	//ctx.Abort()
	//return

	//return ctx.JSON(fiber.Map{
	//	"code":    401,
	//	"message": errorMessage,
	//})

	err := ctx.SendStatus(fiber.StatusBadRequest)
	if err != nil {
		log.Fatal(errorMessage)
	}
}
