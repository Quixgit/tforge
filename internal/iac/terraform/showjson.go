package terraform

import (
	"context"

	"github.com/quix/tforge/internal/iac/runtime"
)

func ShowJSON(ctx context.Context, binary, dir, planfile string) ([]byte, error) {
	runner := runtime.NewRunner(10)

	return runner.Output(ctx, runtime.CommandSpec{
		Engine:  "terraform",
		Binary:  binary,
		Dir:     dir,
		Command: "show-json",
		Args: []string{
			"show",
			"-json",
			planfile,
		},
	})
}
