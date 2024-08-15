package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-chat-service/db"
	"go-chat-service/internal/mysocketio"
	"go-chat-service/routes"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	db.ConnectDatabase()

	mysocketio.Clients = nil

	app := fiber.New()
	routes.Setup(app)

	var port string
	port = os.Getenv("PORT")

	if port == "" {
		port = os.Getenv("DEFAULT_PORT")
	}

	log.Fatal(app.Listen(":" + port))
}
