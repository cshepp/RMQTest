package main

import "encoding/json"

// Config represents the configuration data required by the program
type Config struct {
	Connections []Connection
	Messages    []Message
}

// ParseConfig deserializes a JSON string into a Config type
func ParseConfig(config string) (Config, error) {

	conf := Config{}

	err := json.Unmarshal([]byte(config), &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
