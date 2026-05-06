package terragrunt

import (
	"context"

	"github.com/quix/tforge/internal/core/events"
	"github.com/quix/tforge/internal/iac/engine"
	"github.com/quix/tforge/internal/iac/runtime"
)

type Engine struct {
	engine.Base
}

func New(binary string) engine.Engine {
	if binary == "" {
		binary = "terragrunt"
	}

	return Engine{
		Base: engine.NewBase("terragrunt", binary),
	}
}

func (e Engine) Plan(ctx context.Context, dir string) (<-chan events.Event, error) {
	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "plan",
		Args:    []string{"plan", "-input=false", "-no-color"},
	})
}

func (e Engine) Apply(ctx context.Context, dir string) (<-chan events.Event, error) {
	return e.Runner.Stream(ctx, runtime.CommandSpec{
		Engine:  e.Name(),
		Binary:  e.Binary(),
		Dir:     dir,
		Command: "apply",
		Args:    []string{"apply", "-input=false", "-auto-approve", "-no-color"},
	})
}
