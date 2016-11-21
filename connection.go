package main

import (
	"fmt"
	"net/url"
)

type Connection struct {
	Name     string
	HostName string
	User     string
	Password string
	VHost    string
}

func (c *Connection) GetConnectionString() string {
	encodedVHost := url.QueryEscape(c.VHost)
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", c.User, c.Password, c.HostName, 5672, encodedVHost)
	return connectionString
}
