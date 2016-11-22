package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGenerateContent(t *testing.T) {

	properties := []MessageProperty{
		MessageProperty{
			Name:         "PropertyA",
			DataType:     "string",
			DefaultValue: "TEST",
		},
		MessageProperty{
			Name:         "PropertyB",
			DataType:     "string",
			DefaultValue: "_GENERATE_STRING",
		},
		MessageProperty{
			Name:         "PropertyC",
			DataType:     "int",
			DefaultValue: "_GENERATE_INT",
		},
		MessageProperty{
			Name:         "PropertyD",
			DataType:     "int",
			DefaultValue: "_GENERATE_YO_MAMMA",
		},
	}

	message := Message{
		Name:       "test",
		Exchange:   "test",
		RoutingKey: "test",
		Properties: properties,
	}

	j, err := message.GenerateContent()

	if err != nil {
		t.Error(err)
	}

	var content = make(map[string]interface{})
	if err := json.Unmarshal(j, &content); err != nil {
		t.Error(err)
	}

	if "TEST" != content["PropertyA"] {
		t.Error("Expected default value TEST, got ", content["PropertyA"])
	}

	if reflect.TypeOf(content["PropertyB"]) != reflect.TypeOf("string") {
		t.Error("Expected generated string, got ", reflect.TypeOf(content["PropertyB"]))
	}

	// json.Unmarshal converts ints to float64s, but
	// we really don't care, as long as it's numeric
	if reflect.TypeOf(content["PropertyC"]) != reflect.TypeOf(1.0) {
		t.Error("Expected generated int, got ", reflect.TypeOf(content["PropertyC"]))
	}

	if content["PropertyD"] != "" {
		t.Error("Expected default empty string, got ", content["PropertyD"])
	}
}
