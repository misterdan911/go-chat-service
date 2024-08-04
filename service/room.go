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

func GetRoomList(userId primitive.ObjectID, limit int64, rooms *[]model.RoomData) error {

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
					{"from", "messages"},
					{"localField", "lastMessage"},
					{"foreignField", "_id"},
					{"as", "lastMessage"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$lastMessage"}}}},
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
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$picture"},
					{"preserveNullAndEmptyArrays", true},
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
		bson.D{{"$sort", bson.D{{"lastUpdate", -1}}}},
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
		bson.D{{"$limit", limit}},
	}

	collection := db.DB.Collection("rooms")
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

	return nil
}
