package storage

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dnaCollection                      string = "dna"
	mutantDeterminationServiceDatabase string = "mutant_determination_service"
)

// MongoClient is a mongo client interface
type MongoClient interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	Connect(ctx context.Context) error
}

// StoreDna stores the passed dna into MongoDb
func StoreDna(client MongoClient, matrix []string, isMutant bool) (*mongo.InsertOneResult, error) {
	d := Dna{
		IsMutant: isMutant,
		Matrix:   matrix,
	}
	d.setHash()
	return d.Save(context.Background(), client)
}

// Dna is tha mongodb representation for a dna
type Dna struct {
	IsMutant bool     `json:"is_mutant"`
	Hash     string   `json:"hash"`
	Matrix   []string `json:"matrix"`
}

func (d *Dna) setHash() {
	hash := md5.New()
	for _, str := range d.Matrix {
		hash.Write([]byte(str))
	}
	d.Hash = hex.EncodeToString(hash.Sum(nil))
}

// Save the DNA into the mongodb
func (d Dna) Save(ctx context.Context, client MongoClient) (*mongo.InsertOneResult, error) {
	collection := client.Database(mutantDeterminationServiceDatabase).Collection(dnaCollection)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"hash": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	collection.Indexes().CreateOne(ctx, mod)
	result, err := collection.InsertOne(ctx, d)
	if err != nil {
		fmt.Printf("Error inserting, %v", err)
	}
	return result, err
}
