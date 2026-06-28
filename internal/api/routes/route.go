package routes

import (
	"concurrent-job-processing-system/internal/api/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
}
