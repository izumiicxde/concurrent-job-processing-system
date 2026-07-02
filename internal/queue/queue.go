package queue

import (
	"concurrent-job-processing-system/internal/jobs"
	"sync"
)

type JobQueue interface {
	Enqueue(job *jobs.Job) error
	Dequeue() (*jobs.Job, error)
	Length() int
	Close()
}

type MemoryQueue struct {
	jobs      chan *jobs.Job
	closeOnce sync.Once
}

func NewMemoryQueue(capacity int) *MemoryQueue {
	return &MemoryQueue{
		jobs: make(chan *jobs.Job, capacity),
	}
}

func (mq *MemoryQueue) Enqueue(job *jobs.Job) error {
	mq.jobs <- job
	return nil
}

func (mq *MemoryQueue) Dequeue() (*jobs.Job, error) {
	var job *jobs.Job
	job, ok := <-mq.jobs
	if !ok {
		return nil, jobs.ErrJobQueueClosed
	}
	return job, nil
}

func (mq *MemoryQueue) Length() int {
	return len(mq.jobs)
}

func (mq *MemoryQueue) Close() {

	mq.closeOnce.Do(func() {
		close(mq.jobs)
	})
}
