package main

import (
	"concurrent-job-processing-system/internal/api"
	"concurrent-job-processing-system/internal/config"
	"concurrent-job-processing-system/internal/executor"
	"concurrent-job-processing-system/internal/logger"
	"concurrent-job-processing-system/internal/queue"
	"concurrent-job-processing-system/internal/store"
	"concurrent-job-processing-system/internal/worker"
	"context"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg := config.Load()

	log := logger.New(cfg)
	memoryStore := store.New()
	memoryQueue := queue.NewMemoryQueue(200)
	jobExecutor := executor.New(log)

	workersPool := worker.NewWorkerPool(20, memoryQueue, memoryStore, jobExecutor, log)
	server := api.New(cfg, log, memoryQueue, memoryStore, jobExecutor)

	workersPool.Start()
	server.Run()
	if err := workersPool.Shutdown(context.Background()); err != nil {
		os.Exit(1)
	}
}
