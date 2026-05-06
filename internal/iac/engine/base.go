package engine

import (
	"context"
	"strings"
	"time"

	"github.com/quix/tforge/internal/core/events"
	"github.com/quix/tforge/internal/iac/runtime"
)

type Base struct {
	EngineName string
	Bin        string
	Runner     runtime.Runner
}

func NewBase(name, bin string) Base {
	return Base{
		EngineName: name,
		Bin:        bin,
		Runner:     runtime.NewRunner(30 * time.Minute),
	}
}

func (b Base) Name() string {
	return b.EngineName
}

func (b Base) Binary() string {
	return b.Bin
}

func (b Base) Init(ctx context.Context, dir string) error {
	_, err := b.Runner.Output(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "init",
		Args:    []string{"init", "-input=false"},
	})
	return err
}

func (b Base) Plan(ctx context.Context, dir string) (<-chan events.Event, error) {
	return b.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "plan",
		Args:    []string{"plan", "-input=false", "-no-color"},
	})
}

func (b Base) Apply(ctx context.Context, dir string) (<-chan events.Event, error) {
	return b.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "apply",
		Args:    []string{"apply", "-input=false", "-auto-approve", "-no-color"},
	})
}

func (b Base) Destroy(ctx context.Context, dir string) (<-chan events.Event, error) {
	return b.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "destroy",
		Args:    []string{"destroy", "-input=false", "-auto-approve", "-no-color"},
	})
}

func (b Base) StatePull(ctx context.Context, dir string) ([]byte, error) {
	return b.Runner.Output(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "state pull",
		Args:    []string{"state", "pull"},
	})
}

func (b Base) Workspaces(ctx context.Context, dir string) ([]string, error) {
	out, err := b.Runner.Output(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "workspace list",
		Args:    []string{"workspace", "list"},
	})
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	workspaces := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(strings.TrimPrefix(line, "*"))
		if line != "" {
			workspaces = append(workspaces, line)
		}
	}

	return workspaces, nil
}

func (b Base) SelectWorkspace(ctx context.Context, dir string, name string) error {
	_, err := b.Runner.Output(ctx, runtime.CommandSpec{
		Engine:  b.EngineName,
		Binary:  b.Bin,
		Dir:     dir,
		Command: "workspace select",
		Args:    []string{"workspace", "select", name},
	})
	return err
}
