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
	GetStatisticsController(con, dm)
	return w.Result().StatusCode, w.Result().Body
}

func TestGetStatisticsController(t *testing.T) {

	type getStatisticsControllerTestSuit struct {
		dm          dialMock
		statusCode  int
		body        string
		description string
	}

	getStatistics := []getStatisticsControllerTestSuit{
		getStatisticsControllerTestSuit{
			dm:          dialMock{returnHumanMutant: []byte("2"), returnHuman: []byte("2"), done: []map[string]string{}},
			statusCode:  200,
			body:        `{"count_human_dna":2,"count_mutant_dna":2,"ratio":1}`,
			description: "Two humans, two mutants, ratio 1",
		},
		getStatisticsControllerTestSuit{
			dm:          dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}},
			statusCode:  200,
			body:        `{"count_human_dna":10,"count_mutant_dna":4,"ratio":0.4}`,
			description: "Ten humans, four mutants, ratio 0.4",
		},
	}

	for _, mt := range getStatistics {
		resultStatusCode, resultBody := performRequest(&mt.dm)
		if resultStatusCode != mt.statusCode {
			t.Errorf("%v. Status code is %v and should be %v", mt.description, resultStatusCode, mt.statusCode)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resultBody)
		stringifiedBody := buf.String()
		if stringifiedBody != mt.body {
			t.Errorf("%v. Body is %v and should be %v", mt.description, stringifiedBody, mt.body)
		}
	}

}
