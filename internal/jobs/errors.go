package jobs

import "errors"

var (
	ErrJobNotFound      = errors.New("the requested Job was not found")
	ErrJobAlreadyExists = errors.New("the job already exists")

	ErrJobQueueClosed   = errors.New("the job queue is closed")
	ErrJobDequeueFailed = errors.New("failed to dequeue job")
)
