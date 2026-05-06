package app

import (
	"fmt"

	"github.com/quix/tforge/internal/iac/engine"
	"github.com/quix/tforge/internal/iac/terraform"
	"github.com/quix/tforge/internal/iac/terragrunt"
	"github.com/quix/tforge/internal/iac/tofu"
)

type Runtime struct {
	Options  Options
	Registry engine.Registry
	Engine   engine.Engine
}

func NewRuntime(opts Options) (Runtime, error) {
	registry := engine.NewRegistry()

	registry.Register(terraform.New(opts.TerraformBinary))
	registry.Register(tofu.New(opts.TofuBinary))
	registry.Register(terragrunt.New(opts.TerragruntBinary))

	engineName := opts.Engine
	if engineName == "auto" {
		engineName = engine.Detect(opts.Dir)
	}

	selected, err := registry.Get(engineName)
	if err != nil {
		return Runtime{}, fmt.Errorf("select engine: %w", err)
	}

	return Runtime{
		Options:  opts,
		Registry: registry,
		Engine:   selected,
	}, nil
}
