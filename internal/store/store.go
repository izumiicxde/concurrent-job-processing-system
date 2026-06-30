package store

import "concurrent-job-processing-system/internal/jobs"

type JobStore interface {
	Create(job *jobs.Job) error
	Get(id string) (*jobs.Job, error)
	List() ([]*jobs.Job, error)
	Update(job *jobs.Job) error
	Delete(id string) error
}
