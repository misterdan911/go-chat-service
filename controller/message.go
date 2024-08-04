package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
	"go-chat-service/db"
	"go-chat-service/dto"
	"go-chat-service/internal/mysocketio"
	"go-chat-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type MessageResponse struct {
	Message MessageDetail `json:"message" bson:"message"`
	Room    model.Room    `json:"room" bson:"room"`
}

type MessageDetail struct {
	ID      string    `json:"_id" bson:"_id"`
	Author  Author    `json:"author" bson:"author"`
	Content string    `json:"content" bson:"content"`
	Room    string    `json:"room" bson:"room"`
	Date    time.Time `json:"date" bson:"date"`
	Version int       `json:"__v" bson:"__v"`
}

type Author struct {
	ID         string        `json:"_id" bson:"_id"`
	FirstName  string        `json:"firstName" bson:"firstName"`
	Level      string        `json:"level" bson:"level"`
	Phone      string        `json:"phone" bson:"phone"`
	LastName   string        `json:"lastName" bson:"lastName"`
	Username   string        `json:"username" bson:"username"`
	Favorites  []interface{} `json:"favorites" bson:"favorites"`
	TagLine    string        `json:"tagLine" bson:"tagLine"`
	LastOnline time.Time     `json:"lastOnline" bson:"lastOnline"`
	Picture    Picture       `json:"picture" bson:"picture"`
}

type Picture struct {
	ID         string `json:"_id" bson:"_id"`
	Shield     string `json:"shield" bson:"shield"`
	Name       string `json:"name" bson:"name"`
	Author     string `json:"author" bson:"author"`
	Size       int    `json:"size" bson:"size"`
	Version    int    `json:"__v" bson:"__v"`
	Location   string `json:"location" bson:"location"`
	ShieldedID string `json:"shieldedID" bson:"shieldedID"`
}

type RoomDetail struct {
	ID          string    `json:"_id" bson:"_id"`
	People      []string  `json:"people" bson:"people"`
	IsGroup     bool      `json:"isGroup" bson:"isGroup"`
	Version     int       `json:"__v" bson:"__v"`
	LastAuthor  string    `json:"lastAuthor" bson:"lastAuthor"`
	LastMessage string    `json:"lastMessage" bson:"lastMessage"`
	LastUpdate  time.Time `json:"lastUpdate" bson:"lastUpdate"`
}

func Message(c *fiber.Ctx) error {

	// Create a new User struct
	socketMessage := new(dto.SocketMessage)

	// Parse the JSON request body into the user struct
	if err := c.BodyParser(socketMessage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	var message model.Message
	message.ID = primitive.NewObjectID()
	message.Date = time.Now()
	message.Author = socketMessage.AuthorID
	message.Room = socketMessage.RoomID
	message.Content = socketMessage.Content

	// save message to DB
	_, err := db.DB.Collection("messages").InsertOne(context.Background(), message)

	var updatedDocument model.Room
	filter := bson.D{{"_id", message.Room}}
	update := bson.D{{"$set", bson.D{{"lastMessage", message.ID}, {"lastUpdate", message.Date}}}}
	errUpdate := db.DB.Collection("rooms").FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedDocument)

	if errUpdate != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		//the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Fatal("filter did not match any documents in rooms")
		}
		log.Fatal("Error updating lastMessage")
	}

	//fmt.Println("message.ID: " + message.ID.Hex())

	// Check for errors
	if err != nil {
		return errors.New("Failed Sending message, " + err.Error())
	}

	messagePipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"_id", message.ID}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "users"},
					{"localField", "author"},
					{"foreignField", "_id"},
					{"as", "author"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$author"}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "images"},
					{"localField", "author.picture"},
					{"foreignField", "_id"},
					{"as", "author.picture"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$author.picture"}}}},
	}

	messageCollection := db.DB.Collection("messages")
	cursor, err := messageCollection.Aggregate(context.Background(), messagePipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	// Create a slice of MessageDetail
	var messageDetails []MessageDetail

	if err = cursor.All(context.Background(), &messageDetails); err != nil {
		log.Fatal(err)
	}

	var room model.Room
	err = db.DB.Collection("rooms").FindOne(context.Background(), bson.M{"_id": message.Room}).Decode(&room)

	messageResponse := MessageResponse{
		Message: messageDetails[0],
		Room:    room,
	}

	out, err := json.Marshal(messageResponse)
	if err != nil {
		panic(err)
	}

	// Tes Emit
	//mysocketio.Kws.Emit([]byte(messageResponse.Message.Content), socketio.TextMessage)
	//mysocketio.Kws.Emit(out, socketio.TextMessage)

	var counter int = 0

	for _, userId := range updatedDocument.People {

		fmt.Println(counter)
		fmt.Println(userId.Hex())
		if message.Author != userId {
			err = mysocketio.Kws.EmitTo(mysocketio.Clients[userId.Hex()], out, socketio.TextMessage)
		}
		if err != nil {
			fmt.Println(err)
		}

		counter++
	}

	// supaya field2 response json nya sesuai urutan kita
	jsonResponse, err := json.Marshal(messageResponse)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// supaya response headernya 'application/json'
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonResponse)
}
