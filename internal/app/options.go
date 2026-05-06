package app

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	Dir    string
	Engine string

	TerraformBinary  string
	TofuBinary       string
	TerragruntBinary string

	AllowDestroy bool
}

func ParseOptions(args []string) (Options, error) {
	fs := flag.NewFlagSet("tforge", flag.ContinueOnError)

	var opts Options

	fs.StringVar(&opts.Dir, "dir", ".", "working directory with IaC files")
	fs.StringVar(&opts.Engine, "engine", "auto", "engine: auto, terraform, tofu, terragrunt")

	fs.StringVar(&opts.TerraformBinary, "terraform-binary", "terraform", "terraform binary")
	fs.StringVar(&opts.TofuBinary, "tofu-binary", "tofu", "tofu binary")
	fs.StringVar(&opts.TerragruntBinary, "terragrunt-binary", "terragrunt", "terragrunt binary")
	fs.BoolVar(&opts.AllowDestroy, "allow-destroy", false, "allow destroy actions")

	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}

	dir, err := filepath.Abs(opts.Dir)
	if err != nil {
		return Options{}, err
	}

	info, err := os.Stat(dir)
	if err != nil {
		return Options{}, err
	}

	if !info.IsDir() {
		return Options{}, fmt.Errorf("dir is not a directory: %s", dir)
	}

	opts.Dir = dir

	switch opts.Engine {
	case "auto", "terraform", "tofu", "terragrunt":
	default:
		return Options{}, fmt.Errorf("unsupported engine: %s", opts.Engine)
	}

	return opts, nil
}
