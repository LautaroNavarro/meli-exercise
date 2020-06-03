package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseInterface for Mongo. TODO: When the official mongo library implement interfaces change this
type DatabaseInterface interface {
	Collection(name string) CollectionInterface
}

// CollectionInterface for Mongo. TODO: When the official mongo library implement interfaces change this
type CollectionInterface interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
	Indexes() IndexesInterface
}

// IndexesInterface for Mongo. TODO: When the official mongo library implement interfaces change this
type IndexesInterface interface {
	CreateOne(ctx context.Context, mod mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
}

// ClientInterface for Mongo. TODO: When the official mongo library implement interfaces change this
type ClientInterface interface {
	Database(string) DatabaseInterface
	Connect(ctx context.Context) error
}
