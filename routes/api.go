package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go-chat-service/controller"
	"go-chat-service/middleware"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	apiBaseRoute := app.Group("/api/v1")

	authRoute := apiBaseRoute.Group("/auth")
	authRoute.Post("/signup", controller.SignUp)
	authRoute.Post("/signin", controller.SignIn)

	userRoute := apiBaseRoute.Group("/user", middleware.JwtVerifier)
	userRoute.Get("/", controller.GetAllUser)

}
