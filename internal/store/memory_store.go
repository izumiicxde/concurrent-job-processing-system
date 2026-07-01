package store

import (
	"concurrent-job-processing-system/internal/jobs"
	"sync"
)

type MemoryStore struct {
	jobs map[string]*jobs.Job
	mu   sync.RWMutex
}

func New() JobStore {
	return &MemoryStore{
		jobs: make(map[string]*jobs.Job, 0),
	}
}

func (ms *MemoryStore) Get(id string) (*jobs.Job, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	job, exists := ms.jobs[id]

	if !exists {
		return nil, jobs.ErrJobNotFound
	}
	return job, nil
}

func (ms *MemoryStore) Create(job *jobs.Job) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	_, exists := ms.jobs[job.ID]
	if exists {
		return jobs.ErrJobAlreadyExists
	}
	ms.jobs[job.ID] = job
	return nil
}

func (ms *MemoryStore) List() ([]*jobs.Job, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	var allJobs = make([]*jobs.Job, 0, len(ms.jobs))
	if len(ms.jobs) <= 0 {
		return allJobs, nil
	}

	for _, value := range ms.jobs {
		allJobs = append(allJobs, value)
	}
	return allJobs, nil
}

func (ms *MemoryStore) Update(job *jobs.Job) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	_, exists := ms.jobs[job.ID]
	if !exists {
		return jobs.ErrJobNotFound
	}
	ms.jobs[job.ID] = job // May be redundant for shared pointers, but supports updates from new job instances.
	return nil
}

func (ms *MemoryStore) Delete(ID string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	_, exists := ms.jobs[ID]
	if !exists {
		return jobs.ErrJobNotFound
	}
	delete(ms.jobs, ID)
	return nil
}
