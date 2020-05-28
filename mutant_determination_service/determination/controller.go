package determination

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	mutant     string = "MUTANT"
	notMutant  string = "NOT_MUTANT"
	invalidDna string = "INVALID_DNA"
)

// IsMutantController handle a request with a DNA inside the body, and inject into the context status ok if the
// provided dna is mutant, else inject forbidden
func IsMutantController(ctx *gin.Context) {
	var dna Dna
	if err := ctx.ShouldBindJSON(&dna); err != nil {
		fmt.Println("INVALID REQUEST")
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	isMutant, err := IsMutant(dna.Matrix)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error":         invalidDna,
			"error_message": err.Error(),
		})
		return
	}

	if isMutant {
		ctx.JSON(http.StatusOK, map[string]string{"result": mutant})
	} else {
		ctx.JSON(http.StatusForbidden, map[string]string{"result": notMutant})
	}
}
