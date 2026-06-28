package api

import (
	"concurrent-job-processing-system/internal/api/routes"
	"concurrent-job-processing-system/internal/logger"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const PORT = "8000"

type API struct {
	server *http.Server
	logger *logger.Logger
}

func New(log *logger.Logger) *API {

	api := &API{logger: log}

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	api.server = &http.Server{
		Addr:         ":" + PORT,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return api
}

func (api *API) Run() {

	go func() {
		fmt.Println("Server running on port: ", PORT)
		err := api.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server Error: ", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	fmt.Println("Shutdown Signal Received")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	signal.Stop(signalChan)
	err := api.server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Shutdown error", err)
	}

	fmt.Println("Server shutdown gracefully..")
}
