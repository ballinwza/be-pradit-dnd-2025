package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)


func connectMongo(uri string) (*mongo.Client) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("MongoDB connect error:", err)
	}
	
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		log.Fatalf("Error disconnection: %v", err)
	// 	}
	// 	fmt.Println("Disconnected from MongoDB")
	// }()
	
	return client

	
}
