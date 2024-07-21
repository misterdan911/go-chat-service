package service

import (
	"context"
	"go-chat-service/db"
	"go-chat-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetRoomList(userId primitive.ObjectID, limit int64, rooms *[]model.Room) error {
	collection := db.DB.Collection("rooms")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Define the filter
	filter := bson.M{
		"people": bson.M{"$in": bson.A{userId}},
		"$or": bson.A{
			bson.M{"lastMessage": bson.M{"$ne": nil}},
			bson.M{"isGroup": true},
		},
	}

	// Define the options
	opts := options.Find()
	opts.SetSort(bson.M{"lastUpdate": -1})
	opts.SetLimit(limit)
	opts.SetProjection(bson.M{
		"people.email":    0,
		"people.password": 0,
		"people.friends":  0,
		"people.__v":      0,
	})

	// Find the rooms
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	// Prepare the results
	if err = cursor.All(ctx, rooms); err != nil {
		return err
	}

	return nil
}
