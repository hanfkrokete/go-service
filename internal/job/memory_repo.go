package job

import (
	"context"
	"sync"
)

type MemoryRepo struct {
	data map[string]Job
	mu   sync.RWMutex
}

func (mr *MemoryRepo) Create(ctx context.Context, job Job) (Job, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	if _, ok := mr.data[job.ID]; ok {
		return Job{}, ErrAlreadyExists
	}
	mr.data[job.ID] = job
	return job, nil

}

func (mr *MemoryRepo) Get(ctx context.Context, id string) (Job, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	job, ok := mr.data[id]
	if !ok {
		return Job{}, ErrNotFound
	}

	return job, nil
}

func (mr *MemoryRepo) List(ctx context.Context) ([]Job, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	jobs := make([]Job, 0, len(mr.data))
	for _, job := range mr.data {
		jobs = append(jobs, job)
	}
	return jobs, nil
}
