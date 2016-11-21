package main

import "encoding/json"

type Config struct {
	Connections []Connection
	Messages    []Message
}

func ParseConfig(config string) Config {

	conf := Config{}

	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
