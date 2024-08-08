package controller

import "github.com/gofiber/fiber/v2"

func GetAllUser(c *fiber.Ctx) error {
	return c.SendString("Helooooooooooooooooooooo")
}
