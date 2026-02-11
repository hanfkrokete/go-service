package job

import "time"

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusDone       Status = "done"
	StatusWaiting    Status = "waiting"
)

type Job struct {
	ID        string
	Type      string
	Payload   map[string]any
	Status    Status
	CreatedAt time.Time
}
