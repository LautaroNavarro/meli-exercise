package storage

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

const redisHost string = "redis-cluster-ip-service"
const redisPort string = "6379"

// GetRedisConn returns a redis Conn
func GetRedisConn(redisDial func(network, address string, options ...redis.DialOption) (redis.Conn, error)) redis.Conn {

	password := os.Getenv("REDIS_PASSWORD")

	host := fmt.Sprintf("%v:%v", redisHost, redisPort)

	conn, err := redisDial("tcp", host)

	_, err = conn.Do("AUTH", password)

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	return conn

}
