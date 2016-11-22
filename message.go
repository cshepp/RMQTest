package main

import (
	"math/rand"
	"strconv"
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
type MessageContent map[string]string

// GenerateContent returns a string of the json-encoded message with dynamic properties
func (m *Message) GenerateContent() MessageContent {

	data := make(MessageContent)

	for _, property := range m.Properties {
		gen := getGenerator(property)
		data[property.Name] = gen()
	}

	return data
}

func getGenerator(prop MessageProperty) func() string {
	if !strings.HasPrefix(prop.DefaultValue, "_GENERATE") {
		return func() string { return prop.DefaultValue }
	}

	switch prop.DefaultValue {
	case "_GENERATE_STRING":
		return generateString
	case "_GENERATE_INT":
		return generateInt
	default:
		return func() string { return "" }
	}
}

// Adapted from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func generateString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateInt() string {
	i := rand.Int()
	return strconv.Itoa(i)
}
