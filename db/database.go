package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var postCollection *mongo.Collection
var DB *mongo.Database
var ctx context.Context

func ConnectDatabase() {

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	var opts *options.ClientOptions
	useMongodbAtlas := os.Getenv("USE_MONGODB_ATLAS")

	if useMongodbAtlas == "true" {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts = options.Client().ApplyURI("mongodb+srv://misterdan:M1zt3rD4nz212@cluster0.n5fdxh4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	} else {
		opts = options.Client().ApplyURI("mongodb://localhost:27017")
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	/*
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	*/

	// Send a ping to confirm a successful connection
	/*
		if err := client.Database("clover_db").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
			panic(err)
		}
	*/

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("an error ocurred when connect to mongoDB : %v", err)
		panic(err)
	}

	DB = client.Database("clover_db")

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
