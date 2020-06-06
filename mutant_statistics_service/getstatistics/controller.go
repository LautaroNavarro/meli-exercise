package getstatistics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// Controller handles a request and inyect the statistics for human DNAs
func Controller(ctx *gin.Context, redisConn redis.Conn) {

	st := statistics{rc: redisConn, ctx: ctx}

	ctx.JSON(http.StatusOK, map[string]float64{
		"count_mutant_dna": st.getHumansMutants(),
		"count_human_dna":  st.getHumans(),
		"ratio":            st.getRatio(),
	})
}
