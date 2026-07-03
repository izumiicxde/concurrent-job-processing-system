package worker

import (
	"concurrent-job-processing-system/internal/executor"
	"concurrent-job-processing-system/internal/logger"
	"concurrent-job-processing-system/internal/queue"
	"concurrent-job-processing-system/internal/store"
	"sync"
)

type WorkerPool struct {
	workers []*Worker
	wg      sync.WaitGroup
	logger  *logger.Logger
	queue   queue.JobQueue
}

func NewWorkerPool(workerCount int, queue queue.JobQueue, store store.JobStore, executor *executor.Registry, logger *logger.Logger) *WorkerPool {
	var workers = make([]*Worker, 0, workerCount)
	for i := range workerCount {
		workers = append(workers, NewWorker(i+1, queue, store, executor))
	}

	return &WorkerPool{
		workers: workers,
		queue:   queue,
		logger:  logger,
	}
}

func (wp *WorkerPool) Start() {
	wp.logger.Info("Starting worker pool", "worker_count", len(wp.workers))
	for _, worker := range wp.workers {
		wp.wg.Add(1)
		go func(worker *Worker) {
			defer wp.wg.Done()
			wp.logger.Info("Starting worker", "worker_id", worker.ID)
			worker.Start()
			wp.logger.Info("Worker stopped", "worker_id", worker.ID)
		}(worker)
	}
}
func (wp *WorkerPool) Shutdown() {
	wp.logger.Info("Worker pool shutdown initiated", "worker_count", len(wp.workers))
	wp.logger.Info("Closing job queue")
	wp.queue.Close()

	wp.logger.Info("Waiting for all workers to finish and exit")
	wp.wg.Wait()
	wp.logger.Info("Worker pool shutdown complete")
}
