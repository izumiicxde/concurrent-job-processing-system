package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port     string `env:"PORT"`
	LogPath  string `env:"LOG_PATH"`
	LogLevel string `env:"LOG_LEVEL"`
}

func Load() *Config {
	var config Config
	err := env.Parse(config)
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	return &config
}
