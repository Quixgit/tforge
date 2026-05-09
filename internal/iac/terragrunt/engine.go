package terragrunt

import (
	"context"

	"github.com/quix/tforge/internal/core/events"
	"github.com/quix/tforge/internal/iac/engine"
	"github.com/quix/tforge/internal/iac/runtime"
)

type Engine struct {
	engine.Base
	RunAll bool
}

func New(binary string) engine.Engine {
	if binary == "" {
		binary = "terragrunt"
	}

	return Engine{
		Base:   engine.NewBase("terragrunt", binary),
		RunAll: true,
	}
}

func (e Engine) Validate(ctx context.Context, dir string) (<-chan events.Event, error) {
	args := []string{"validate", "-no-color"}

	if e.RunAll {
		args = []string{"run-all", "validate", "-no-color", "--terragrunt-non-interactive"}
	}

	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "validate",
		Args:    args,
	})
}

func (e Engine) Plan(ctx context.Context, dir string) (<-chan events.Event, error) {
	args := []string{"plan", "-input=false", "-no-color"}

	if e.RunAll {
		args = []string{"run-all", "plan", "-input=false", "-no-color", "--terragrunt-non-interactive"}
	}

	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "plan",
		Args:    args,
	})
}

func (e Engine) Apply(ctx context.Context, dir string) (<-chan events.Event, error) {
	args := []string{"apply", "-input=false", "-auto-approve", "-no-color"}

	if e.RunAll {
		args = []string{"run-all", "apply", "-input=false", "-auto-approve", "-no-color", "--terragrunt-non-interactive"}
	}

	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "apply",
		Args:    args,
	})
}

func (e Engine) Destroy(ctx context.Context, dir string) (<-chan events.Event, error) {
	args := []string{"destroy", "-input=false", "-auto-approve", "-no-color"}

	if e.RunAll {
		args = []string{"run-all", "destroy", "-input=false", "-auto-approve", "-no-color", "--terragrunt-non-interactive"}
	}

	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "destroy",
		Args:    args,
	})
}
