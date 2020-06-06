package updatestatistics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type registration struct {
	IsMutant *bool `json:"is_mutant" binding:"required"`
}

// Controller handles a request and update the redis statistics
func Controller(ctx *gin.Context, redisConn redis.Conn) {

	var re registration
	if err := ctx.ShouldBindJSON(&re); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	hry := humanRegistry{rc: redisConn, ctx: ctx}

	hry.registerHuman(*(re.IsMutant))

}
