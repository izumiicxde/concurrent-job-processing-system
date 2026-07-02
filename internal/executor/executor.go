package executor

import (
	"concurrent-job-processing-system/internal/jobs"
	"errors"
)

var (
	ErrExecutorAlreadyRegistered = errors.New("executor is already registered")
	ErrExecutorNotFound          = errors.New("executor not found in registry")
)

type Registry struct {
	executors map[jobs.JobType]Executor
}

func (r *Registry) Register(jobtype jobs.JobType, executor Executor) error {
	_, exists := r.executors[jobtype]
	if exists {
		return ErrExecutorAlreadyRegistered
	}
	r.executors[jobtype] = executor
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

	return registry
}

type Executor interface {
	Execute(job *jobs.Job) error
}

type EmailExecutor struct{}

func (ex *EmailExecutor) Execute(job *jobs.Job) error {
	return nil
}

type NotificationExecutor struct{}

func (ex *NotificationExecutor) Execute(job *jobs.Job) error {
	return nil
}
