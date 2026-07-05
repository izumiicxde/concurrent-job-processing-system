package handlers

import (
	"concurrent-job-processing-system/internal/executor"
	"concurrent-job-processing-system/internal/helpers"
	"concurrent-job-processing-system/internal/jobs"
	"concurrent-job-processing-system/internal/logger"
	"concurrent-job-processing-system/internal/queue"
	"concurrent-job-processing-system/internal/store"
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	logger   *logger.Logger
	queue    queue.JobQueue
	store    store.JobStore
	executor *executor.Registry
}

func New(logger *logger.Logger, queue queue.JobQueue, store store.JobStore, executor *executor.Registry) *Handler {
	return &Handler{logger: logger, queue: queue, store: store, executor: executor}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, err := w.Write([]byte("System Up and Running on port 8000"))
		if err != nil {
			h.logger.HTTPError(r.Method, "/", http.StatusInternalServerError, r.RemoteAddr, err)
		}
	}
}

//  JOB HANDLERS

// GetJobs : Get all jobs.
func (h *Handler) GetJobs(w http.ResponseWriter, _ *http.Request) {
	var allJobs []*jobs.Job
	allJobs, err := h.store.List()
	if err != nil {
		h.WriteError(w, http.StatusInternalServerError, JOB_FETCH_FAILED, err.Error())
		return
	}

	h.WriteResponse(w, http.StatusOK, &GetAllJobResponse{Total: len(allJobs), Jobs: allJobs})
}

// GetJob : Get a single job by ID
func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request) {
	var job *jobs.Job
	id := r.PathValue("id")
	if id == "" {
		h.WriteError(w, http.StatusBadRequest, INVALID_JOB_ID, "Provide a valid job ID")
		return
	}

	job, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, jobs.ErrJobNotFound) {
			h.WriteError(w, http.StatusBadRequest, JOB_NOT_FOUND, err.Error())
			return
		}
		h.WriteError(w, http.StatusInternalServerError, JOB_FETCH_FAILED, err.Error())
		return
	}
	h.WriteResponse(w, http.StatusOK, &GetJobByIDResponse{Job: job})
}

// CreateJob : Create a new Job
func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req jobs.CreateJobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.WriteError(w, http.StatusBadRequest, INVALID_JOB_PAYLOAD, "Invalid payload")
		return
	}

	if _, err := h.executor.Get(req.Type); err != nil {
		h.WriteError(w, http.StatusBadRequest, INVALID_JOB_TYPE, "The job type given is invalid")
		return
	}
	job := &jobs.Job{
		ID:         helpers.CreateJobID(),
		Type:       req.Type,
		Payload:    req.Payload,
		Status:     jobs.JobStatusPending,
		Priority:   jobs.JobPriorityNormal,
		MaxRetries: req.MaxRetries,
	}
	if err := h.store.Create(job); err != nil {
		h.WriteError(w, http.StatusInternalServerError, JOB_CREATION_FAILED, err.Error())
		return
	}

	if err := h.queue.Enqueue(job); err != nil {
		job.Status = jobs.JobStatusFailed
		_ = h.store.Delete(job.ID)

		h.WriteError(w, http.StatusInternalServerError, JOB_CREATION_FAILED, err.Error())
		return
	}
	job.Status = jobs.JobStatusQueued
	_ = h.store.Update(job)

	h.WriteResponse(w, http.StatusAccepted, &CreateJobResponse{ID: job.ID, Status: job.Status})
}

// DeleteJob : Delete a job by ID
func (h *Handler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var job *jobs.Job
	id := r.PathValue("id")
	if id == "" {
		h.WriteError(w, http.StatusBadRequest, INVALID_JOB_ID, "Provide a valid job ID")
		return
	}
	job, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, jobs.ErrJobNotFound) {
			h.WriteError(w, http.StatusBadRequest, JOB_NOT_FOUND, err.Error())
			return
		}
		h.WriteError(w, http.StatusInternalServerError, JOB_FETCH_FAILED, err.Error())
		return
	}

	if job.Status != jobs.JobStatusPending && job.Status != jobs.JobStatusFailed && job.Status != jobs.JobStatusCancelled && job.Status != jobs.JobStatusCompleted {
		h.WriteError(w, http.StatusBadRequest, JOB_DELETION_FAILED, "Job cannot be deleted while queued, running, or retrying")
		return
	}
	if err := h.store.Delete(id); err != nil {
		if errors.Is(err, jobs.ErrJobNotFound) {
			h.WriteError(w, http.StatusNotFound, JOB_NOT_FOUND, err.Error())
			return
		}
		h.WriteError(w, http.StatusInternalServerError, JOB_DELETION_FAILED, "Failed to delete the Job")
		return
	}

	h.WriteResponse(w, http.StatusOK, &DeleteJobByIDResponse{ID: id, Deleted: true})
}
