package controller

import "github.com/gofiber/fiber/v2"

// GetAllUser godoc
//
//	@Summary		Get All User
//	@Description	Get All User
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		model.User
//	@Router			/api/v1/user [get]
func GetAllUser(c *fiber.Ctx) error {
	return c.SendString("Helooooooooooooooooooooo")
}
