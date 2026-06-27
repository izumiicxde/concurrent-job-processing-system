package api

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
}

const PORT = "8000"

func InitializeAPI() *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("System up and working...."))
		if err != nil {
			log.Fatal(err)
		}
	})

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
