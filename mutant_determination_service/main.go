package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"meli-exercise/mutant_determination_service/determination"
	"meli-exercise/mutant_determination_service/storage"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mutantStatisticsService string = "http://mutant-statistics-service-cluster-ip-service:8080"
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
				if err == nil {
					fmt.Println("CALLING MUTANT STATISTICS SERVICE")
					body := map[string]bool{"is_mutant": isMutant}
					jsonBody, _ := json.Marshal(body)
					_, callErr := http.Post(
						fmt.Sprintf("%v/stats/", mutantStatisticsService),
						"application/json",
						bytes.NewBuffer(jsonBody),
					)
					fmt.Printf("CALL ACTION ERROR: %v \n", callErr)
				}
			},
		)
	})
	router.Run(":8080")
}
