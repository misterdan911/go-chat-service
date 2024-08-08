package main

import (
	"github.com/gofiber/fiber/v2"
	_ "go-chat-service/docs"
	"log"
	"os"
)

func main() {

	/*
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		db.ConnectDatabase()

		mysocketio.Clients = nil
	*/

	app := fiber.New()
	//routes.Setup(app)

	var port string
	port = os.Getenv("PORT")

	/*
		if port == "" {
			port = os.Getenv("DEFAULT_PORT")
		}
	*/

	log.Fatal(app.Listen(":" + port))
}
