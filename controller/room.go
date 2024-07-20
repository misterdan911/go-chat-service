package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type RoomListDto struct {
	Limit int `json:"limit"`
}

func RoomList(c *fiber.Ctx) error {

	username := c.Locals("login_username").(string)
	fmt.Println("Username: " + username)

	/*
		roomListDto := new(RoomListDto)

		if err := c.BodyParser(roomListDto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse RoomList JSON",
			})
		}

		roomList := service.GetRoomList(roomListDto.Limit)
	*/

	return nil
}
