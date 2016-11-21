package main

import "testing"

func TestGetConnectionString(t *testing.T) {

	connection := Connection{
		Name:     "TestConnection",
		HostName: "amqp.test.com",
		User:     "user",
		Password: "password",
		VHost:    "/Test",
	}

	connString := connection.GetConnectionString()

	if "amqp://user:password@amqp.test.com:5672/%2FTest" != connString {
		t.Error("Expected amqp://user:password@amqp.test.com:5672/%2FTest, got ", connString)
	}
}
