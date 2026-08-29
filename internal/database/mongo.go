package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo establishes a connection to MongoDB and returns the client instance.
// TS Equivalent: export const connectMongo = async (uri: string): Promise<*mongo.Client>
func ConnectMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// 'defer' is a unique Go keyword. It means: "Run this function right before ConnectMongo finishes and returns."
	// We use it here to ensure the memory used by the timeout is cleaned up, no matter what happens below.
	defer cancel()

	// Configure the client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	// TS Equivalent: const client = await MongoClient.connect(uri, options);
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	// mongo.Connect doesn't actually ping the database, it just sets up the config.
	// We must explicitly Ping it to ensure the connection is truly alive.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("MongoDB successfully connected!")

	// Return the pointer to the client so the rest of our app can use it
	return client, nil
}