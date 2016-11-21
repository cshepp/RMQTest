package main

import "encoding/json"

type Config struct {
	Connections []Connection
	Messages    []Message
}

func ParseConfig(config string) (Config, error) {

	conf := Config{}

	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
