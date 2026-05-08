package app

import "github.com/quix/tforge/internal/project"

type batchItemStatus string

const (
	batchPending batchItemStatus = "pending"
	batchRunning batchItemStatus = "running"
	batchSuccess batchItemStatus = "success"
	batchFailed  batchItemStatus = "failed"
)

type batchItem struct {
	Target project.Target
	Status batchItemStatus
	Error  string
}

type batchStartedMsg struct{}
type batchNextMsg struct{}
type batchItemFinishedMsg struct {
	index int
	err   error
}
