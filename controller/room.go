package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-chat-service/model"
	"go-chat-service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type RoomListDto struct {
	Limit int64 `json:"limit"`
}

func RoomList(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(primitive.ObjectID)
	fmt.Println("UserId: " + userId.Hex())

	// convert string to primitive.ObjectID
	// primitive.ObjectIDFromHex(hexString)

	roomListDto := new(RoomListDto)

	if err := c.BodyParser(roomListDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse RoomList JSON",
		})
	}

	// Extract limit from query parameters
	limit := int64(30)
	if roomListDto.Limit != 0 {
		limit = roomListDto.Limit
	}

	rooms := new([]model.RoomData)

	err := service.GetRoomList(userId, roomListDto.Limit, rooms)
	if err != nil {
		return err
	}

	// Send the response
	response := map[string]interface{}{
		"limit": limit,
		"rooms": rooms,
	}

	// supaya field2 response json nya sesuai urutan kita
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// supaya response headernya 'application/json'
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)

}
