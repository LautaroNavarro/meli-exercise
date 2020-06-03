package storage

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetConnection(t *testing.T) {

	// newClient func(*options.ClientOptions) (MongoClient, error)
	ctx := context.Background()
	var usedClientOption *options.ClientOptions

	os.Setenv("MONGOUSER", "test_user")
	os.Setenv("MONGOPASSWORD", "test_password")
	os.Setenv("MONGOHOST", "test_host")

	expectedURI := "mongodb://test_user:test_password@test_host/"

	cn := GetConnection(
		ctx,
		func(opt *options.ClientOptions) (MongoClient, error) {
			usedClientOption = opt
			return mongoClientMock{}, nil
		},
	)

	if cn == nil {
		t.Error("Empty connection. Connection must not be nil")
	}

	if usedClientOption.GetURI() != expectedURI {
		t.Errorf("Invalid URI connection. Expected %v , gotten %v", expectedURI, usedClientOption.GetURI())
	}

}
