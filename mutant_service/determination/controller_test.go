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
	req, _ := sling.New().Post("/mutant").BodyJSON(dna).Request()
	con.Request = req
	IsMutantController(con)
	return w.Result().StatusCode
}

func TestIsMutantController(t *testing.T) {

	type isMutantControllerTest struct {
		dna         Dna
		statusCode  int
		description string
	}

	isMutantControllerTests := []isMutantControllerTest{
		isMutantControllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "ATGC", "ATGC", "TAGC"}},
			statusCode:  http.StatusBadRequest,
			description: "Matrix is 5X6. Must be 6x6",
		},
		isMutantControllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "AZGC", "ATGC", "TAGC", "ATCG"}},
			statusCode:  http.StatusBadRequest,
			description: "Matrix is 6X5. Must be 6x6",
		},
		isMutantControllerTest{
			dna:         Dna{Matrix: []string{"ATGC", "TAGC", "AZGC", "ATGC", "TAGC", "ATCG"}},
			statusCode:  http.StatusBadRequest,
			description: "Invalid char 'Z'",
		},
		isMutantControllerTest{
			dna:         Dna{Matrix: []string{"AAAAGA", "ACGTGC", "AAAAGT", "TGAACG", "CGCCTT", "TCACTG"}},
			statusCode:  http.StatusOK,
			description: "Valid mutant",
		},
		isMutantControllerTest{
			dna:         Dna{Matrix: []string{"ATGCGA", "ACGTGC", "ATGCAG", "TGAACG", "CGCCTT", "TCACTG"}},
			statusCode:  http.StatusForbidden,
			description: "Valid not mutant",
		},
	}

	for _, ct := range isMutantControllerTests {
		result := performRequest(ct.dna)
		if result != ct.statusCode {
			t.Errorf("%v. Status code is %v and should be %v", ct.description, result, ct.statusCode)
		}
	}

}
