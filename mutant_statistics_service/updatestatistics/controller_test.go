package updatestatistics

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func performRequest(dm redis.Conn, rg registration) int {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	con, _ := gin.CreateTestContext(w)
	req, _ := sling.New().Post("/stats/").BodyJSON(rg).Request()
	con.Request = req
	Controller(con, dm)
	return w.Result().StatusCode
}

func TestGetStatisticsController(t *testing.T) {

	type controllerTest struct {
		dm                     dialMock
		statusCode             int
		expectedRedisExecution string
		isMutant               bool
		description            string
		rg                     registration
	}
	trueValue := true
	falseValue := false
	controllerTestSuite := []controllerTest{
		controllerTest{
			dm:                     dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}},
			statusCode:             200,
			rg:                     registration{IsMutant: &trueValue},
			expectedRedisExecution: `[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]] map[SET:[humans-mutants 5]]]`,
			description:            "If is mutant, add one to humans and one to mutants",
		},
		controllerTest{
			dm:                     dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}},
			statusCode:             200,
			rg:                     registration{IsMutant: &falseValue},
			expectedRedisExecution: `[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]]]`,
			description:            "If is not mutant, add one to humans",
		},
	}

	for _, ct := range controllerTestSuite {
		resultStatusCode := performRequest(&ct.dm, ct.rg)
		if resultStatusCode != ct.statusCode {
			t.Errorf("%v. Status code is %v and should be %v", ct.description, resultStatusCode, ct.statusCode)
		}

		redisExecution, _ := ct.dm.Receive()
		if ct.expectedRedisExecution != fmt.Sprintf("%v", redisExecution) {
			t.Errorf("%v. Body is %v and should be %v", ct.description, redisExecution, ct.expectedRedisExecution)
		}
	}

}
