package controller

import "github.com/gofiber/fiber/v2"

func Info(c *fiber.Ctx) error {

	// Create a JSON response
	response := fiber.Map{
		"version":           "2.9.0",
		"build":             2,
		"nodemailerEnabled": false,
	}

	// Send the JSON response
	return c.JSON(response)
}
