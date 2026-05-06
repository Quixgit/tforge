package runtime

import (
	"context"

	"github.com/quix/tforge/internal/iac/engine"
)

func showJSON(
	ctx context.Context,
	eng engine.Engine,
	dir string,
	planfile string,
) ([]byte, error) {

	base, ok := eng.(interface {
		Binary() string
	})
	if !ok {
		return nil, nil
	}

	runner := getRunner(eng)

	return runner.Output(ctx, spec(
		eng.Name(),
		base.Binary(),
		dir,
		"show-json",
		"show",
		"-json",
		planfile,
	))
}
