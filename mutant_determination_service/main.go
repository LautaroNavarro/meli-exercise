package main

import (
	"context"
	"fmt"
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
			func(matrix []string, isMutant bool) {
				_, err := storage.StoreDna(client, matrix, isMutant)
				if err != nil {
					fmt.Println("Not MSS service")
				} else {
					fmt.Println("Calling MSS service")
				}
			},
		)
	})
	router.Run(":8080")
}
