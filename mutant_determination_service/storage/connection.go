package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient wrapper for mongo. TODO: When the official mongo library implement interfaces change this
type MongoClient struct {
	Cl *mongo.Client
}

// MongoDatabase wrapper for mongo. TODO: When the official mongo library implement interfaces change this
type MongoDatabase struct {
	Db *mongo.Database
}

// MongoCollection wrapper for mongo. TODO: When the official mongo library implement interfaces change this
type MongoCollection struct {
	Coll *mongo.Collection
}

// MongoSingleResult wrapper for mongo. TODO: When the official mongo library implement interfaces change this
type MongoSingleResult struct {
	sr *mongo.SingleResult
}

// MongoIndexes wrapper for mongo. TODO: When the official mongo library implement interfaces change this
type MongoIndexes struct {
	in mongo.IndexView
}

// Database wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (mc *MongoClient) Database(dbName string) DatabaseInterface {
	db := mc.Cl.Database(dbName)
	return &MongoDatabase{Db: db}
}

// Connect wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (mc *MongoClient) Connect(ctx context.Context) error {
	return mc.Cl.Connect(ctx)
}

// Collection wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (md *MongoDatabase) Collection(colName string) CollectionInterface {
	collection := md.Db.Collection("dna")
	return &MongoCollection{Coll: collection}
}

// InsertOne wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (mc *MongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return mc.Coll.InsertOne(ctx, document)
}

// Indexes wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (mc *MongoCollection) Indexes() IndexesInterface {
	return &MongoIndexes{in: mc.Coll.Indexes()}
}

// CreateOne wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (in *MongoIndexes) CreateOne(ctx context.Context, mod mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return in.in.CreateOne(ctx, mod)
}

// Decode wrapper for mongo. TODO: When the official mongo library implement interfaces change this
func (sr *MongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

// GetConnection returns a mongo.Client connection
func GetConnection(ctx context.Context, newClient func(*options.ClientOptions) (ClientInterface, error)) ClientInterface {
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
