package worker

import (
	"concurrent-job-processing-system/internal/executor"
	"concurrent-job-processing-system/internal/jobs"
	"concurrent-job-processing-system/internal/queue"
	"concurrent-job-processing-system/internal/store"
	"errors"
	"time"
)

type Worker struct {
	ID       int
	Queue    queue.JobQueue
	Store    store.JobStore
	Executor *executor.Registry
}

func NewWorker(id int, queue queue.JobQueue, store store.JobStore, executor *executor.Registry) *Worker {
	return &Worker{
		ID:       id,
		Queue:    queue,
		Store:    store,
		Executor: executor,
	}
}

func (w *Worker) Start() {
	for {
		job, err := w.Queue.Dequeue()
		if err != nil {
			if errors.Is(err, jobs.ErrJobQueueClosed) {
				return
			}
			continue
		}
		job.Status = jobs.JobStatusRunning
		job.StartedAt = time.Now()
		_ = w.Store.Update(job)

		jobExecutor, err := w.Executor.Get(job.Type)

		if err != nil {
			job.Status = jobs.JobStatusFailed
			job.FinishedAt = time.Now()
			job.Error = err.Error()
			_ = w.Store.Update(job)
			continue
		}

		for {
			if job.Retries <= job.MaxRetries {
				job.Status = jobs.JobStatusRunning

				if err := jobExecutor.Execute(job); err != nil {
					job.Error = err.Error()
					job.Retries++
					job.Status = jobs.JobStatusRetrying
					_ = w.Store.Update(job)
				} else {
					job.Error = ""
					job.Status = jobs.JobStatusCompleted
					job.FinishedAt = time.Now()
					_ = w.Store.Update(job)
					break
				}
			} else {
				job.Status = jobs.JobStatusFailed
				job.FinishedAt = time.Now()
				_ = w.Store.Update(job)
				break
			}
		}
	}
}
