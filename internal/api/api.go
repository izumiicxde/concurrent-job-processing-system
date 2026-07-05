package api

import (
	"concurrent-job-processing-system/internal/api/handlers"
	"concurrent-job-processing-system/internal/api/middlewares"
	"concurrent-job-processing-system/internal/executor"
	"concurrent-job-processing-system/internal/queue"
	"concurrent-job-processing-system/internal/store"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"concurrent-job-processing-system/internal/api/routes"
	"concurrent-job-processing-system/internal/config"
	"concurrent-job-processing-system/internal/logger"
)

type API struct {
	server *http.Server
	logger *logger.Logger
	cfg    *config.Config
}

func New(cfg *config.Config, log *logger.Logger, queue queue.JobQueue, store store.JobStore, executor *executor.Registry) *API {
	api := &API{logger: log, cfg: cfg}

	mux := http.NewServeMux()
	handler := handlers.New(log, queue, store, executor)
	routes.RegisterRoutes(mux, handler)

	api.server = &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middlewares.Logging(log)(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return api
}

func (api *API) Run() {
	go func() {
		api.logger.Info("Server Started", "port", api.cfg.Port)
		err := api.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			api.logger.Error("Server Error: ", "error", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	api.logger.Info("Server shutdown Signal Received")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	signal.Stop(signalChan)
	err := api.server.Shutdown(ctx)
	if err != nil {
		api.logger.Error("Error while shutting down server", "error", err)
	}

	api.logger.Info("Server shutdown gracefully")
	api.logger.Close()
}
