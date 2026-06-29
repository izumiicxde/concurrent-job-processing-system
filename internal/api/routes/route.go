package routes

import (
	"concurrent-job-processing-system/internal/api/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handlers *handlers.Handler) {
	mux.HandleFunc("/health", handlers.Health)
}
