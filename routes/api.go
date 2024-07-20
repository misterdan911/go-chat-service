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

	apiBaseRoute.Post("/register", controller.SignUp)
	apiBaseRoute.Post("/login", controller.SignIn)

	//apiBaseRoute.Post("/rooms/list", controller.RoomList)

	// Group routes that require JWT verification
	protectedRoutes := apiBaseRoute.Group("/", middleware.JwtVerifier)
	protectedRoutes.Post("/rooms/list", controller.RoomList)

	//apiBaseRoute.Post("/meeting/list", controller.RoomList)
	//apiBaseRoute.Post("/favorites/list", controller.RoomList)
	//apiBaseRoute.Post("/search", controller.RoomList)

	userRoute := apiBaseRoute.Group("/user", middleware.JwtVerifier)
	userRoute.Get("/", controller.GetAllUser)

	// Initialize a new hub
	h := hub.NewHub()

	// Run the hub in a separate goroutine
	go h.Run()

	app.Use("/ws", middleware.WsAllowUpgrade)
	app.Use("ws/chat", websocket.New(controller.HandleChat(h)))

}
