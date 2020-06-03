package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetConnection returns a mongo.Client connection
func GetConnection(ctx context.Context, newClient func(*options.ClientOptions) (MongoClient, error)) MongoClient {
	conString := fmt.Sprintf(
		"mongodb://%v:%v@%v/",
		os.Getenv("MONGOUSER"),
		os.Getenv("MONGOPASSWORD"),
		os.Getenv("MONGOHOST"),
	)

	client, err := newClient(options.Client().ApplyURI(conString))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	return client
}
