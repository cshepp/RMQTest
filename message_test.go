package main

import (
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

	content := message.GenerateContent()

	if "TEST" != content["PropertyA"] {
		t.Error("Expected default value TEST, got ", content["PropertyA"])
	}

	if reflect.TypeOf(content["PropertyB"]) != reflect.TypeOf("string") {
		t.Error("Expected generated string, got ", content["PropertyB"])
	}

	if reflect.TypeOf(content["PropertyC"]) != reflect.TypeOf(1) {
		t.Error("Expected generated int, got ", content["PropertyC"])
	}

	if content["PropertyD"] != "" {
		t.Error("Expected default empty string, got ", content["PropertyD"])
	}
}
