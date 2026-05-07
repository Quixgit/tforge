package execution

import "time"

type Status string

const (
	StatusQueued   Status = "queued"
	StatusRunning  Status = "running"
	StatusComplete Status = "complete"
	StatusFailed   Status = "failed"
)

type ResourceState struct {
	Address   string
	Action    string
	Status    Status
	StartedAt time.Time
	EndedAt   time.Time
	Error     string
}

func (r ResourceState) Duration() time.Duration {
	if r.StartedAt.IsZero() {
		return 0
	}

	if r.EndedAt.IsZero() {
		return time.Since(r.StartedAt)
	}

	return r.EndedAt.Sub(r.StartedAt)
}
