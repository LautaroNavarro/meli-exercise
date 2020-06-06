package getstatistics

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func performRequest(dm redis.Conn) (int, io.ReadCloser) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)
	req, _ := sling.New().Get("/public/stats/").Request()
	con.Request = req
	Controller(con, dm)
	return w.Result().StatusCode, w.Result().Body
}

func TestController(t *testing.T) {

	type controllerTest struct {
		dm          dialMock
		statusCode  int
		body        string
		description string
	}

	controllerTestSuite := []controllerTest{
		controllerTest{
			dm:          dialMock{returnHumanMutant: []byte("2"), returnHuman: []byte("2"), done: []map[string]string{}},
			statusCode:  200,
			body:        `{"count_human_dna":2,"count_mutant_dna":2,"ratio":1}`,
			description: "Two humans, two mutants, ratio 1",
		},
		controllerTest{
			dm:          dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}},
			statusCode:  200,
			body:        `{"count_human_dna":10,"count_mutant_dna":4,"ratio":0.4}`,
			description: "Ten humans, four mutants, ratio 0.4",
		},
	}

	for _, ct := range controllerTestSuite {
		resultStatusCode, resultBody := performRequest(&ct.dm)
		if resultStatusCode != ct.statusCode {
			t.Errorf("%v. Status code is %v and should be %v", ct.description, resultStatusCode, ct.statusCode)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resultBody)
		stringifiedBody := buf.String()
		if stringifiedBody != ct.body {
			t.Errorf("%v. Body is %v and should be %v", ct.description, stringifiedBody, ct.body)
		}
	}

}
