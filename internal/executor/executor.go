package executor

import (
	"concurrent-job-processing-system/internal/jobs"
	"concurrent-job-processing-system/internal/logger"
	"errors"
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

func New(logger *logger.Logger) *Registry {
	registry := &Registry{
		executors: make(map[jobs.JobType]Executor),
	}

	registry.mustRegister(
		jobs.JobTypeSendEmail,
		&EmailExecutor{logger: logger},
	)

	registry.mustRegister(
		jobs.JobTypeSendNotification,
		&NotificationExecutor{logger: logger},
	)

	registry.mustRegister(
		jobs.JobTypeGenerateThumbnail,
		&ThumbnailExecutor{logger: logger},
	)

	registry.mustRegister(
		jobs.JobTypeCompressFiles,
		&CompressFilesExecutor{logger: logger},
	)

	registry.mustRegister(
		jobs.JobTypeExportUserData,
		&ExportUserDataExecutor{logger: logger},
	)

	return registry
}

type EmailExecutor struct {
	logger *logger.Logger
}

func (ex *EmailExecutor) Execute(job *jobs.Job) error {
	ex.logger.Info(
		"Executing email job",
		"job_id", job.ID,
		"job_type", job.Type,
	)

	time.Sleep(2 * time.Second)

	ex.logger.Info(
		"Email job completed",
		"job_id", job.ID,
	)

	return nil
}

type NotificationExecutor struct {
	logger *logger.Logger
}

func (ex *NotificationExecutor) Execute(job *jobs.Job) error {
	ex.logger.Info(
		"Executing notification job",
		"job_id", job.ID,
		"job_type", job.Type,
	)

	time.Sleep(500 * time.Millisecond)

	ex.logger.Info(
		"Notification job completed",
		"job_id", job.ID,
	)

	return nil
}

type ThumbnailExecutor struct {
	logger *logger.Logger
}

func (ex *ThumbnailExecutor) Execute(job *jobs.Job) error {
	ex.logger.Info(
		"Executing thumbnail generation job",
		"job_id", job.ID,
		"job_type", job.Type,
	)

	time.Sleep(3 * time.Second)

	ex.logger.Info(
		"Thumbnail generated successfully",
		"job_id", job.ID,
	)

	return nil
}

type CompressFilesExecutor struct {
	logger *logger.Logger
}

func (ex *CompressFilesExecutor) Execute(job *jobs.Job) error {
	ex.logger.Info(
		"Executing file compression job",
		"job_id", job.ID,
		"job_type", job.Type,
	)

	time.Sleep(4 * time.Second)

	ex.logger.Info(
		"Files compressed successfully",
		"job_id", job.ID,
	)

	return nil
}

type ExportUserDataExecutor struct {
	logger *logger.Logger
}

func (ex *ExportUserDataExecutor) Execute(job *jobs.Job) error {
	ex.logger.Info(
		"Executing user data export job",
		"job_id", job.ID,
		"job_type", job.Type,
	)

	time.Sleep(2 * time.Second)

	ex.logger.Info(
		"User data exported successfully",
		"job_id", job.ID,
	)

	return nil
}
