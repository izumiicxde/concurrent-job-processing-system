package routes

import (
	"concurrent-job-processing-system/internal/api/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handlers *handlers.Handler) {
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/jobs", handlers.GetJobs)
	mux.HandleFunc("GET /job/{id}", handlers.GetJob)
	mux.HandleFunc("DELETE /job/{id}", handlers.DeleteJob)
}
