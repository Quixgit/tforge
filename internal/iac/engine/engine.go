package engine

import (
	"context"

	"github.com/quix/tforge/internal/core/events"
)

type Engine interface {
	Name() string
	Binary() string

	Init(ctx context.Context, dir string) error
	Plan(ctx context.Context, dir string) (<-chan events.Event, error)
	Apply(ctx context.Context, dir string) (<-chan events.Event, error)
	Destroy(ctx context.Context, dir string) (<-chan events.Event, error)

	StatePull(ctx context.Context, dir string) ([]byte, error)
	Workspaces(ctx context.Context, dir string) ([]string, error)
	SelectWorkspace(ctx context.Context, dir string, name string) error
}
