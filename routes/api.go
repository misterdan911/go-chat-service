package routes

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-chat-service/controller"
	"go-chat-service/internal/mysocketio"
	"go-chat-service/middleware"
)

type Iseng struct {
	User string `bson:"user" json:"user"`
}

func Setup(app *fiber.App) {

	app.Use(cors.New())

	// Setup the middleware to retrieve the data sent in first GET request
	/*
		app.Use(func(c *fiber.Ctx) error {
			// IsWebSocketUpgrade returns true if the client
			// requested upgrade to the WebSocket protocol.
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			return fiber.ErrUpgradeRequired
		})
	*/

	// The key for the map is message.to
	mysocketio.Clients = make(map[string]string)
	mysocketio.DefineSocketAction()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket) {

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// socketio to manage Emit/EmitTo/Broadcast
		//clients[userId] = kws.UUID
		mysocketio.Clients[userId] = kws.UUID

		for _, socketClient := range mysocketio.Clients {
			fmt.Println("socketClient: " + socketClient)
		}

		// Every websocket connection has an optional session key => value storage
		kws.SetAttribute("user_id", userId)

		//Broadcast to all the connected users the newcomerout, err := json.Marshal(message2)
		//		if err != nil {
		//			panic(err)
		//		}
		//kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, socketio.TextMessage)
		//Write welcome message
		//kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)), socketio.TextMessage)

		/*
			var message2 Iseng
			message2 = Iseng{
				User: "budi",
			}

			kws.Emit(out, socketio.TextMessage)
		*/

		mysocketio.Kws = kws
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	apiBaseRoute := app.Group("/api")

	apiBaseRoute.Post("/register", controller.SignUp)
	apiBaseRoute.Post("/login", controller.SignIn)
	apiBaseRoute.Get("/info", controller.Info)
	apiBaseRoute.Post("/message", controller.SaveMessage)

	//apiBaseRoute.Post("/rooms/list", controller.RoomList)

	// Group routes that require JWT verification
	protectedRoutes := apiBaseRoute.Group("/", middleware.JwtVerifier)
	protectedRoutes.Post("/rooms/list", controller.RoomList)

	//apiBaseRoute.Post("/meeting/list", controller.RoomList)
	//apiBaseRoute.Post("/favorites/list", controller.RoomList)
	//apiBaseRoute.Post("/search", controller.RoomList)

	userRoute := apiBaseRoute.Group("/user", middleware.JwtVerifier)
	userRoute.Get("/", controller.GetAllUser)

	/*
		// Initialize a new hub
		h := hub.NewHub()

		// Run the hub in a separate goroutine
		go h.Run()

		app.Use("/ws", middleware.WsAllowUpgrade)
		app.Use("ws/chat", websocket.New(controller.HandleChat(h)))
	*/

}
