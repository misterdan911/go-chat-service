package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go-chat-service/controller"
	"go-chat-service/internal/hub"
	"go-chat-service/middleware"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	apiBaseRoute := app.Group("/api")

	authRoute := apiBaseRoute.Group("/auth")
	authRoute.Post("/signup", controller.SignUp)
	authRoute.Post("/login", controller.SignIn)

	userRoute := apiBaseRoute.Group("/user", middleware.JwtVerifier)
	userRoute.Get("/", controller.GetAllUser)

	// Initialize a new hub
	h := hub.NewHub()

	// Run the hub in a separate goroutine
	go h.Run()

	app.Use("/ws", middleware.WsAllowUpgrade)
	app.Use("ws/chat", websocket.New(controller.HandleChat(h)))

}
