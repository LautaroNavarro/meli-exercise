package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/public/hello/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"hello": "world",
		})
	})
	router.Run(":8080")
}
