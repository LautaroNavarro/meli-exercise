package main

import (
	"meli-exercise/mutant_statistics_service/getstatistics"
	"meli-exercise/mutant_statistics_service/storage"
	"meli-exercise/mutant_statistics_service/updatestatistics"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func main() {

	router := gin.Default()
	redisConn := storage.GetRedisConn(redis.Dial)
	router.GET("/public/stats/", func(ctx *gin.Context) {
		getstatistics.Controller(ctx, redisConn)
	})
	router.POST("/stats/", func(ctx *gin.Context) {
		updatestatistics.Controller(ctx, redisConn)
	})
	router.Run(":8080")
}
