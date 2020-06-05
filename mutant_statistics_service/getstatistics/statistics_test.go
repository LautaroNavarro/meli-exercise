package getstatistics

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

func TestStatistics_GetHumans(t *testing.T) {
	d := dialMock{returnHuman: []byte("12"), done: []map[string]string{}}
	s := statistics{ctx: context.Background(), rc: &d}
	humans := s.getHumans()
	if humans != float64(12) {
		t.Errorf("Humans expected to be 12 but was %v", humans)
	}
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]]]", fmt.Sprintf("%v", done))
	}
}
func TestStatistics_GetHumansMutants(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("2"), done: []map[string]string{}}
	s := statistics{ctx: context.Background(), rc: &d}
	humantMutant := s.getHumansMutants()
	if humantMutant != float64(2) {
		t.Errorf("humantMutant expected to be 2 but was %v", humantMutant)
	}
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans-mutants]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans-mutants]]]", fmt.Sprintf("%v", done))
	}
}

func TestStatistics_GetRatio(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("2"), returnHuman: []byte("2"), done: []map[string]string{}}
	s := statistics{ctx: context.Background(), rc: &d}
	ratio := s.getRatio()
	if ratio != float64(1) {
		t.Errorf("Ratio expected to be 1 but was %v", ratio)
	}
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]] map[GET:[humans-mutants]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]] map[GET:[humans-mutants]]]", fmt.Sprintf("%v", done))
	}
}

func TestStatistics_GetRatio_CaseTwo(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("10"), done: []map[string]string{}}
	s := statistics{ctx: context.Background(), rc: &d}
	ratio := s.getRatio()
	if ratio != float64(0.4) {
		t.Errorf("Ratio expected to be 0.4 but was %v", ratio)
	}
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]] map[GET:[humans-mutants]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]] map[GET:[humans-mutants]]]", fmt.Sprintf("%v", done))
	}
}

func TestStatistics_GetRatio_ZeroDivision(t *testing.T) {
	d := dialMock{returnHumanMutant: []byte("4"), returnHuman: []byte("0"), done: []map[string]string{}}
	s := statistics{ctx: context.Background(), rc: &d}
	ratio := s.getRatio()
	if ratio != float64(0) {
		t.Errorf("Ratio expected to be 0 but was %v", ratio)
	}
	done, _ := d.Receive()
	if fmt.Sprintf("%v", done) != "[map[GET:[humans]] map[GET:[humans-mutants]]]" {
		t.Errorf("Expected to be executed %v but executed %v", "[map[GET:[humans]] map[GET:[humans-mutants]]]", fmt.Sprintf("%v", done))
	}
}
