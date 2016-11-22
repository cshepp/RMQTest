package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var configfile = flag.String("config", "conf.json", "config file to load")

func main() {

	flag.Parse()

	config := loadConfig(*configfile)
	connection := promptToSelectConnection(config.Connections)
	message := promptToSelectMessage(config.Messages)

	err := PublishMessage(connection, message)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func loadConfig(path string) Config {

	// load contents of config file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading config file")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// parse config file
	config, err := ParseConfig(string(b))
	if err != nil {
		fmt.Println("Error parsing config file")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return config
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
		fmt.Println("Invalid input")
		return promptToSelectConnection(connections)
	}

	if i > len(connections) {
		fmt.Println("Invalid input")
		return promptToSelectConnection(connections)
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
		fmt.Println("Invalid input")
		return promptToSelectMessage(messages)
	}

	if i > len(messages) {
		fmt.Println("Invalid input")
		return promptToSelectMessage(messages)
	}

	return messages[i]
}
