package updatestatistics

import (
	"context"
	"fmt"
	"testing"
)

type dialMock struct {
	done              []map[string]string
	returnHuman       []byte
	returnHumanMutant []byte
}

func (d *dialMock) Close() error {
	return nil
}

func (d *dialMock) Err() error {
	return nil
}

func (d *dialMock) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	stringifiedArgs := fmt.Sprintf("%v", args)
	d.done = append(d.done, map[string]string{commandName: stringifiedArgs})
	if len(args) != 0 && args[0] == "humans" {
		return d.returnHuman, nil
	}
	return d.returnHumanMutant, nil
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

func TestHumanRegistry_RegisterHuman(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}}
	s := humanRegistry{ctx: context.Background(), rc: &d}
	s.registerHuman(true)
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]] map[SET:[humans-mutants 5]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]] map[SET:[humans-mutants 5]]]", fmt.Sprintf("%v", done))
	}
}

func TestHumanRegistry_RegisterHumanNotMutant(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}}
	s := humanRegistry{ctx: context.Background(), rc: &d}
	s.registerHuman(false)
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]] map[GET:[humans-mutants]] map[SET:[humans 11]]]", fmt.Sprintf("%v", done))
	}
}
