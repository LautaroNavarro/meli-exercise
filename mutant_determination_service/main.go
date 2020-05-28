package main

import (
	"meli-exercise/mutant_determination_service/determination"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/mutant/", determination.IsMutantController)
	router.Run(":8080")
}
