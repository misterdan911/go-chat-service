package service

import (
	"context"
	"go-chat-service/db"
	"go-chat-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetRoomList(userId primitive.ObjectID, limit int64, rooms *[]model.Room) error {

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match",
				bson.D{
					{"people", userId},
					{"$or",
						bson.A{
							bson.D{{"lastMessage", bson.D{{"$ne", primitive.Null{}}}}},
							bson.D{{"isGroup", true}},
						},
					},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "users"},
					{"localField", "people"},
					{"foreignField", "_id"},
					{"as", "people"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$people"}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "images"},
					{"localField", "people.picture"},
					{"foreignField", "_id"},
					{"as", "people.picture"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$people.picture"}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "images"},
					{"localField", "picture"},
					{"foreignField", "_id"},
					{"as", "picture"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$picture"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$_id"},
					{"people", bson.D{{"$push", "$people"}}},
					{"title", bson.D{{"$first", "$title"}}},
					{"picture", bson.D{{"$first", "$picture"}}},
					{"isGroup", bson.D{{"$first", "$isGroup"}}},
					{"lastAuthor", bson.D{{"$first", "$lastAuthor"}}},
					{"lastMessage", bson.D{{"$first", "$lastMessage"}}},
					{"lastUpdate", bson.D{{"$first", "$lastUpdate"}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"people.email", 0},
					{"people.password", 0},
					{"people.__v", 0},
					{"people.picture.__v", 0},
					{"picture.__v", 0},
				},
			},
		},
	}

	/*
		matchStage := bson.D{{"$match", bson.D{
			{"people", bson.D{{"$in", bson.A{userId}}}},
		}}}

		lookupStage := bson.D{{"$lookup", bson.D{
			{"from", "users"},
			{"localField", "people"},
			{"foreignField", "_id"},
			{"as", "people"},
		}}}

		//unwindStage := bson.D{{"$unwind", "$people"}}

		lookupStage2 := bson.D{{"$lookup", bson.D{
			{"from", "images"},
			{"localField", "people.picture"},
			{"foreignField", "_id"},
			{"as", "people.picture"},
		}}}

		projectStage := bson.D{{"$project", bson.D{
			{"people.password", 0},
		}}}

	*/

	collection := db.DB.Collection("rooms")
	//cursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{matchStage, lookupStage, lookupStage2, projectStage})
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	if err = cursor.All(context.Background(), rooms); err != nil {
		log.Fatal(err)
	}

	/*
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
	*/

	return nil
}
