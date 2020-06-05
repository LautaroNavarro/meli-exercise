package storage

import (
	"fmt"
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"
)

type dialMock struct {
	done map[string]string
}

func (d *dialMock) Close() error {
	return nil
}

func (d *dialMock) Err() error {
	return nil
}

func (d *dialMock) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	stringifiedArgs := fmt.Sprintf("%v", args)
	d.done[commandName] = stringifiedArgs
	return "", nil
}

func (d *dialMock) Send(commandName string, args ...interface{}) error {
	return nil
}

func (d *dialMock) Flush() error {
	return nil
}

func (d *dialMock) Receive() (reply interface{}, err error) {
	return d.done, nil
}

func redisDialMock(network, address string, options ...redis.DialOption) (redis.Conn, error) {
	return &dialMock{done: map[string]string{}}, nil
}

func TestGetRedisConn(t *testing.T) {
	var resultNetwork string
	var resultAddress string
	expectedAuth := "map[AUTH:[test_password]]"
	os.Setenv("REDIS_PASSWORD", "test_password")
	c := GetRedisConn(
		func(network, address string, options ...redis.DialOption) (redis.Conn, error) {
			resultNetwork = network
			resultAddress = address
			return redisDialMock(network, address, options...)
		},
	)
	if resultAddress != "redis-cluster-ip-service:6379" {
		t.Errorf("Address expected to be %v . But was %v", "redis-cluster-ip-service:6379", resultNetwork)
	}
	if resultNetwork != "tcp" {
		t.Errorf("Network expected to be %v . But was %v", "redis-cluster-ip-service:6379", resultNetwork)
	}
	done, _ := c.Receive()

	if fmt.Sprintf("%v", done) != expectedAuth {
		t.Errorf("Auth expected %v . But got %v", expectedAuth, fmt.Sprintf("%v", done))
	}
}
