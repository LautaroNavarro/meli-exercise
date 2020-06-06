package updatestatistics

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

type humanRegistry struct {
	rc  redis.Conn
	ctx context.Context
}

func (hr *humanRegistry) registerHuman(isMutant bool) error {
	var humans float64
	var mutants float64
	var err error
	humans, err = redis.Float64(hr.rc.Do("GET", "humans"))
	mutants, err = redis.Float64(hr.rc.Do("GET", "humans-mutants"))
	_, err = hr.rc.Do("SET", "humans", humans+1)
	if isMutant {
		_, err = hr.rc.Do("SET", "humans-mutants", mutants+1)
	}
	return err
}
