package executor

import (
	"concurrent-job-processing-system/internal/jobs"
	"errors"
	"log"
	"time"
)

var (
	ErrExecutorAlreadyRegistered = errors.New("executor is already registered")
	ErrExecutorNotFound          = errors.New("executor not found in registry")
)

type Executor interface {
	Execute(job *jobs.Job) error
}

type Registry struct {
	executors map[jobs.JobType]Executor
}

func (r *Registry) Register(jobType jobs.JobType, executor Executor) error {
	if _, exists := r.executors[jobType]; exists {
		return ErrExecutorAlreadyRegistered
	}

	r.executors[jobType] = executor
	return nil
}

func (r *Registry) Get(jobType jobs.JobType) (Executor, error) {
	executor, exists := r.executors[jobType]
	if !exists {
		return nil, ErrExecutorNotFound
	}

	return executor, nil
}

func (r *Registry) mustRegister(jobType jobs.JobType, executor Executor) {
	if err := r.Register(jobType, executor); err != nil {
		panic(err)
	}
}

func New() *Registry {
	registry := &Registry{
		executors: make(map[jobs.JobType]Executor),
	}

	registry.mustRegister(jobs.JobTypeSendEmail, &EmailExecutor{})
	registry.mustRegister(jobs.JobTypeSendNotification, &NotificationExecutor{})
	registry.mustRegister(jobs.JobTypeGenerateThumbnail, &ThumbnailExecutor{})
	registry.mustRegister(jobs.JobTypeCompressFiles, &CompressFilesExecutor{})
	registry.mustRegister(jobs.JobTypeExportUserData, &ExportUserDataExecutor{})

	return registry
}

type EmailExecutor struct{}

func (ex *EmailExecutor) Execute(job *jobs.Job) error {
	log.Printf("[EMAIL] Processing job=%s", job.ID)

	time.Sleep(2 * time.Second)

	log.Printf("[EMAIL] Email sent successfully for job=%s", job.ID)

	return nil
}

type NotificationExecutor struct{}

func (ex *NotificationExecutor) Execute(job *jobs.Job) error {
	log.Printf("[NOTIFICATION] Processing job=%s", job.ID)

	time.Sleep(500 * time.Millisecond)

	log.Printf("[NOTIFICATION] Notification sent successfully for job=%s", job.ID)

	return nil
}

type ThumbnailExecutor struct{}

func (ex *ThumbnailExecutor) Execute(job *jobs.Job) error {
	log.Printf("[THUMBNAIL] Processing job=%s", job.ID)

	time.Sleep(3 * time.Second)

	log.Printf("[THUMBNAIL] Thumbnail generated for job=%s", job.ID)

	return nil
}

type CompressFilesExecutor struct{}

func (ex *CompressFilesExecutor) Execute(job *jobs.Job) error {
	log.Printf("[ZIP] Processing job=%s", job.ID)

	time.Sleep(4 * time.Second)

	log.Printf("[ZIP] Archive created successfully for job=%s", job.ID)

	return nil
}

type ExportUserDataExecutor struct{}

func (ex *ExportUserDataExecutor) Execute(job *jobs.Job) error {
	log.Printf("[EXPORT] Processing job=%s", job.ID)

	time.Sleep(2 * time.Second)

	log.Printf("[EXPORT] User data exported successfully for job=%s", job.ID)

	return nil
}
