package main

import (
	"context"
	"meli-exercise/mutant_determination_service/determination"
	"meli-exercise/mutant_determination_service/storage"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	router := gin.Default()
	ctx := context.Background()

	client := storage.GetConnection(ctx, func(opt *options.ClientOptions) (storage.ClientInterface, error) {
		client, _ := mongo.NewClient(opt)
		return &storage.MongoClient{Cl: client}, nil
	})

	router.POST("/public/mutant/", func(ctx *gin.Context) {
		determination.IsMutantController(
			ctx,
			func(matrix []string, isMutant bool) { storage.StoreDna(client, matrix, isMutant) },
		)
	})
	router.Run(":8080")
}
