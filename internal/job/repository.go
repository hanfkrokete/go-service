package job

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("job not found")
	ErrAlreadyExists = errors.New("job already exists")
)

type Repository interface {
	Create(ctx context.Context, job Job) (Job, error)
	Get(ctx context.Context, id string) (Job, error)
	List(ctx context.Context) ([]Job, error)
}
