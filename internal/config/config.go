package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultServerPort = "8080"
)

type Config struct {
	ServerPort int
}

func LoadFromEnv() (*Config, error) {
	conf := &Config{}
	var err error
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = defaultServerPort
	}

	conf.ServerPort, err = strconv.Atoi(serverPort)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s as int: %w", os.Getenv("SERVER_PORT"), err)
	}

	return conf, nil
}
