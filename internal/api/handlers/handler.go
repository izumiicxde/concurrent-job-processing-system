package handlers

import (
	"concurrent-job-processing-system/internal/logger"
	"log"
	"net/http"
)

type Handler struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, err := w.Write([]byte("System Up and Running on port 8000"))
		if err != nil {
			log.Println("ERROR: GET /health route failed. ", err)
		}
	}
}

//  JOB HANDLERS

// GetJobs : Get all jobs.
func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {

}

// GetJob : Get a single job by ID
func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request) {

}

// CreateJob : Create a new Job
func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {

}

// DeleteJob : Delete a job by ID
func (h *Handler) DeleteJob(w http.ResponseWriter, r *http.Request) {

}
