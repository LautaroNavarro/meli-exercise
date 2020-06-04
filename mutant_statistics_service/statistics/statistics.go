package statistics

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

type statistics struct {
	rc                 redis.Conn
	ctx                context.Context
	humansMutantsCache bool
	humansMutants      float64
	humansCache        bool
	humans             float64
	ratio              float64
}

func (st *statistics) getHumans() float64 {
	val, err := redis.Float64(st.rc.Do("GET", "humans"))
	if err != nil {
		val = 0
	}
	st.humans = val
	st.humansCache = true
	return st.humans
}

func (st *statistics) getHumansMutants() float64 {
	val, err := redis.Float64(st.rc.Do("GET", "humans-mutants"))
	if err != nil {
		val = 0
	}
	st.humansMutants = val
	st.humansMutantsCache = true
	return st.humansMutants
}

func (st *statistics) getRatio() float64 {
	if st.humansCache == false {
		st.getHumans()
	}
	if st.humansMutantsCache == false {
		st.getHumansMutants()
	}
	if st.humans == 0 {
		return 0
	}
	return float64(st.humansMutants / st.humans)
}
