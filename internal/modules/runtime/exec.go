package runtime

import (
	"context"

	"github.com/quix/tforge/internal/iac/engine"
)

func planTerraform(
	ctx context.Context,
	eng engine.Engine,
	dir string,
	outfile string,
) error {

	base, ok := eng.(interface {
		Binary() string
	})
	if !ok {
		return nil
	}

	runner := getRunner(eng)

	_, err := runner.Output(ctx, spec(
		eng.Name(),
		base.Binary(),
		dir,
		"plan",
		"plan",
		"-input=false",
		"-no-color",
		"-out="+outfile,
	))

	return err
}

func planTerragrunt(
	ctx context.Context,
	eng engine.Engine,
	dir string,
	outfile string,
) error {

	return planTerraform(ctx, eng, dir, outfile)
}
