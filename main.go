package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "go-chat-service/docs"
	"log"
	"os"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

// @securitydefinitions.oauth2.application	OAuth2Application
// @tokenUrl								https://example.com/oauth/token
// @scope.write							Grants write access
// @scope.admin							Grants read and write access to administrative information
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//db.ConnectDatabase()

	//mysocketio.Clients = nil

	app := fiber.New()
	//routes.Setup(app)

	var port string
	port = os.Getenv("PORT")

	if port == "" {
		port = os.Getenv("DEFAULT_PORT")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
