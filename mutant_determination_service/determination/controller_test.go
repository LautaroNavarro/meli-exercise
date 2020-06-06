package determination

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

func performRequest(dna Dna) int {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)
	req, _ := sling.New().Post("/public/mutant").BodyJSON(dna).Request()
	con.Request = req
	IsMutantController(con, func(matrix []string, isMutant bool) {})
	return w.Result().StatusCode
}

func TestIsMutantController(t *testing.T) {

	type controllerTest struct {
		dna         Dna
		statusCode  int
		description string
	}

	controllerTestSuite := []controllerTest{
		controllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "ATGC", "ATGC", "TAGC"}},
			statusCode:  http.StatusBadRequest,
			description: "Matrix is 5X6. Must be 6x6",
		},
		controllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "AZGC", "ATGC", "TAGC", "ATCG"}},
			statusCode:  http.StatusBadRequest,
			description: "Matrix is 6X5. Must be 6x6",
		},
		controllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "AZGC", "ATGC", "TAGC", "ATCG"}},
			statusCode:  http.StatusBadRequest,
			description: "Invalid char 'Z'",
		},
		controllerTest{
			dna:         Dna{Matrix: []string{"AAAAGA", "ACGTGC", "AAAAGT", "TGAACG", "CGCCTT", "TCACTG"}},
			statusCode:  http.StatusOK,
			description: "Valid mutant",
		},
		controllerTest{
			dna:         Dna{Matrix: []string{"ATGCGA", "ACGTGC", "ATGCAG", "TGAACG", "CGCCTT", "TCACTG"}},
			statusCode:  http.StatusForbidden,
			description: "Valid not mutant",
		},
	}

	for _, ct := range controllerTestSuite {
		result := performRequest(ct.dna)
		if result != ct.statusCode {
			t.Errorf("%v. Status code is %v and should be %v", ct.description, result, ct.statusCode)
		}
	}

}
