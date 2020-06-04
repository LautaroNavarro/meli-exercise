package main

import (
	"meli-exercise/mutant_statistics_service/statistics"

	"meli-exercise/mutant_statistics_service/storage"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func main() {

	router := gin.Default()
	redisConn := storage.GetRedisConn(redis.Dial)
	router.POST("/public/stats/", func(ctx *gin.Context) {
		statistics.GetStatisticsController(ctx, redisConn)
	})

	router.Run(":8080")
}
