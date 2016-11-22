package main

import (
	"io/ioutil"
	"testing"
)

func TestParseConfig(t *testing.T) {

	b, err := ioutil.ReadFile("sample.conf.json")
	if err != nil {
		panic(err)
	}

	configString := string(b)
	result, err := ParseConfig(configString)

	if len(result.Connections) != 1 {
		t.Error("Expected 1, got", len(result.Connections))
	}

	if len(result.Messages) != 1 {
		t.Error("Expected 1, got", len(result.Messages))
	}
}
