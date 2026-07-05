package handlers

import (
	"concurrent-job-processing-system/internal/jobs"
	"encoding/json"
	"net/http"
)

type CreateJobResponse struct {
	ID     string         `json:"id"`
	Status jobs.JobStatus `json:"status"`
}

type GetAllJobResponse struct {
	Total int         `json:"total"`
	Jobs  []*jobs.Job `json:"jobs"`
}
type GetJobByIDResponse struct {
	Job *jobs.Job `json:"job"`
}
type DeleteJobByIDResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

func (h *Handler) WriteResponse(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		h.logger.Error("Failed to encode http response", "status_code", statusCode, "error", err)
	}
}

func (h *Handler) WriteError(w http.ResponseWriter, statusCode int, code errorResponseType, message string) {
	errResponse := ErrorResponse{Error: Error{Code: code, Message: message}}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		h.logger.Error("Failed to encode http response", "status_code", statusCode, "error", err)
	}
}
