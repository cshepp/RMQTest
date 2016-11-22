package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

// PublishMessage opens a connection/channel to RabbitMQ, serializes
// the message content, and publishes the message
func PublishMessage(connection Connection, message Message) error {
	// Connect to RMQ
	conn, err := amqp.Dial(connection.GetConnectionString())
	if err != nil {
		return errors.New("Error connecting to RabbitMQ : " + err.Error())
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		return errors.New("Error opening a channel : " + err.Error())
	}
	defer ch.Close()

	content := message.GenerateContent()

	// Send a message
	body, err := json.Marshal(&content)
	if err != nil {
		return errors.New("Error converting message to JSON : " + err.Error())
	}

	err = ch.Publish(
		message.Exchange,
		message.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		return errors.New("Error publishing message : " + err.Error())
	}

	fmt.Println("SENT:")
	fmt.Println(string(body))
	fmt.Println("TO:")
	fmt.Println(connection.Name)

	return nil
}
