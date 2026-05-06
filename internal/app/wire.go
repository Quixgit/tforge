package app

import (
	"fmt"

	"github.com/quix/tforge/internal/config"
	"github.com/quix/tforge/internal/iac/engine"
	"github.com/quix/tforge/internal/iac/terraform"
	"github.com/quix/tforge/internal/iac/terragrunt"
	"github.com/quix/tforge/internal/iac/tofu"
)

type Runtime struct {
	Options  Options
	Config   config.Config
	Registry engine.Registry
	Engine   engine.Engine
}

func NewRuntime(opts Options) (Runtime, error) {
	cfg, err := config.Load(opts.Dir)
	if err != nil {
		return Runtime{}, fmt.Errorf("load config: %w", err)
	}

	if opts.AllowDestroy {
		cfg.Security.AllowDestroy = true
	}

	if opts.Engine != "auto" {
		cfg.Engine = opts.Engine
	}

	if opts.TerraformBinary != "" {
		cfg.Terraform.Binary = opts.TerraformBinary
	}
	if opts.TofuBinary != "" {
		cfg.Tofu.Binary = opts.TofuBinary
	}
	if opts.TerragruntBinary != "" {
		cfg.Terragrunt.Binary = opts.TerragruntBinary
	}

	registry := engine.NewRegistry()

	registry.Register(terraform.New(cfg.Terraform.Binary))
	registry.Register(tofu.New(cfg.Tofu.Binary))
	registry.Register(terragrunt.New(cfg.Terragrunt.Binary))

	engineName := cfg.Engine
	if engineName == "auto" {
		engineName = engine.Detect(opts.Dir)
	}

	selected, err := registry.Get(engineName)
	if err != nil {
		return Runtime{}, fmt.Errorf("select engine: %w", err)
	}

	return Runtime{
		Options:  opts,
		Config:   cfg,
		Registry: registry,
		Engine:   selected,
	}, nil
}
