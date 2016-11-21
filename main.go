package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

func main() {

	config := loadConfig()

	connection := promptToSelectConnection(config.Connections)
	connectionString := connection.GetConnectionString()

	message := promptToSelectMessage(config.Messages)
	messageContent := message.GenerateContent()

	// Connect to RMQ
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Send a message
	body, err := json.Marshal(&messageContent)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	fmt.Println("SENT:")
	fmt.Println(string(body))
	fmt.Println("TO:")
	fmt.Println(connection.Name)
}

func loadConfig() Config {
	// get any args passed in
	args := os.Args[1:]

	// this is the default config file path
	configPath := "conf.json"

	// check to see if a custom config file path was specified
	if len(args) == 1 {
		configPath = args[0]
	}

	// load contents of config file
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	// return parsed config file
	return ParseConfig(string(b))
}

func promptToSelectConnection(connections []Connection) Connection {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number of the connection you'd like to use:")

	// print out a list of available servers
	for i, conn := range connections {
		fmt.Printf("%d: %s \n", i, conn.Name)
	}

	input, _ := reader.ReadString('\n')
	cleanedInputWindows := strings.Replace(input, "\r\n", "", -1)
	cleanedInput := strings.Replace(cleanedInputWindows, "\n", "", -1)

	i, err := strconv.Atoi(cleanedInput)
	if err != nil {
		panic(err)
	}

	if i > len(connections) {
		panic("Invalid number entered")
	}

	return connections[i]
}

func promptToSelectMessage(messages []Message) Message {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number of the message you'd like to send:")

	// print out a list of available messages
	for i, m := range messages {
		fmt.Printf("%d: %s \n", i, m.Name)
	}

	input, _ := reader.ReadString('\n')
	cleanedInputWindows := strings.Replace(input, "\r\n", "", -1)
	cleanedInput := strings.Replace(cleanedInputWindows, "\n", "", -1)

	i, err := strconv.Atoi(cleanedInput)
	if err != nil {
		panic(err)
	}

	if i > len(messages) {
		panic("Invalid number entered")
	}

	return messages[i]
}
