package storage

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoClientMock ...
type mongoClientMock struct {
}

// mongoDatabaseMock ...
type mongoDatabaseMock struct {
}

// mongoCollectionMock ...
type mongoCollectionMock struct {
}

// mongoSingleResultMock ...
type mongoSingleResultMock struct {
}

// mongoIndexesMock ...
type mongoIndexesMock struct {
}

// Database ...
func (mc *mongoClientMock) Database(dbName string) DatabaseInterface {
	return &mongoDatabaseMock{}
}

// Connect ...
func (mc *mongoClientMock) Connect(ctx context.Context) error {
	return nil
}

// Collection ...
func (md *mongoDatabaseMock) Collection(colName string) CollectionInterface {
	return &mongoCollectionMock{}
}

// InsertOne ...
func (mc *mongoCollectionMock) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return document, nil
}

// Indexes ...
func (mc *mongoCollectionMock) Indexes() IndexesInterface {
	return &mongoIndexesMock{}
}

// CreateOne ...
func (in *mongoIndexesMock) CreateOne(ctx context.Context, mod mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return "", nil
}

// Decode ...
func (sr *mongoSingleResultMock) Decode(v interface{}) error {
	return nil
}

func TestGetConnection(t *testing.T) {

	ctx := context.Background()
	var usedClientOption *options.ClientOptions

	os.Setenv("MONGOUSER", "test_user")
	os.Setenv("MONGOPASSWORD", "test_password")
	os.Setenv("MONGOHOST", "test_host")

	expectedURI := "mongodb://test_user:test_password@test_host/"

	cn := GetConnection(
		ctx,
		func(opt *options.ClientOptions) (ClientInterface, error) {
			usedClientOption = opt
			return &mongoClientMock{}, nil
		},
	)

	if cn == nil {
		t.Error("Empty connection. Connection must not be nil")
	}

	if usedClientOption.GetURI() != expectedURI {
		t.Errorf("Invalid URI connection. Expected %v , gotten %v", expectedURI, usedClientOption.GetURI())
	}

}
