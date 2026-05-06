package tasks

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

type Task struct {
	Action  string
	Status  Status
	Logs    []string
	Running bool
}
