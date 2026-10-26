package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func NewConfig() Config {
	port := 8080
	if value := os.Getenv("PORT"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 && parsed <= 65535 {
			port = parsed
		}
	}

	return Config{Port: port}
}
