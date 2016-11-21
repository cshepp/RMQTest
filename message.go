package main

import (
	"math/rand"
	"strings"
)

// Message describes a message we can send to the queue
type Message struct {
	Name       string
	Exchange   string
	RoutingKey string
	Properties []MessageProperty
}

// MessageProperty describes a property of a message
type MessageProperty struct {
	Name         string
	DataType     string
	DefaultValue string
}

// MessageContent provides a dynamic structure for json-ifying messages
type MessageContent map[string]interface{}

// GenerateContent returns a string of the json-encoded message with dynamic properties
func (m *Message) GenerateContent() MessageContent {

	data := make(map[string]interface{})

	for _, property := range m.Properties {
		gen := getGenerator(property)
		data[property.Name] = gen()
	}

	return data
}

func getGenerator(prop MessageProperty) func() interface{} {
	if !strings.HasPrefix(prop.DefaultValue, "_GENERATE") {
		return func() interface{} { return prop.DefaultValue }
	}

	switch prop.DefaultValue {
	case "_GENERATE_STRING":
		return generateString
	case "_GENERATE_INT":
		return generateInt
	default:
		return func() interface{} { return "" }
	}
}

func generateString() interface{} {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateInt() interface{} {
	return rand.Int()
}
