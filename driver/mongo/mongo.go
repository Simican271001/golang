package mongo_driver

import (
	"charum/util"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(databaseName string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Make MongoDB URI
	dbUser := util.GetConfig("DB_USER")
	dbPass := util.GetConfig("DB_PASS")
	dbHost := util.GetConfig("DB_HOST") // Assume DB_HOST does not include "mongodb+srv" protocol and might include a port
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s/?retryWrites=true&w=majority", dbUser, dbPass, dbHost)

	// Set client options
	clientOptions := options.Client().ApplyURI(mongodbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	database := client.Database(databaseName)
	return database
}
