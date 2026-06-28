package main

import (
	"concurrent-job-processing-system/internal/api"
	"concurrent-job-processing-system/internal/config"
	"concurrent-job-processing-system/internal/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg := config.Load()

	log := logger.New(cfg)
	server := api.New(log)
	server.Run()

	if err := log.Close(); err != nil {
		panic(err)
	}
}
