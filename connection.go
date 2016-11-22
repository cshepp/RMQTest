package main

import (
	"fmt"
	"net/url"
)

// A Connection describes the necessary data for connecting to a RabbitMQ server
type Connection struct {
	Name     string
	HostName string
	User     string
	Password string
	VHost    string
}

// GetConnectionString returns the amqp connection string for a connection
func (c *Connection) GetConnectionString() string {
	encodedVHost := url.QueryEscape(c.VHost)
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.User, c.Password, c.HostName, 5672, encodedVHost)
	return connectionString
}
