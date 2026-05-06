package engine

import (
	"os"
	"path/filepath"
)

func Detect(dir string) string {
	if exists(filepath.Join(dir, "terragrunt.hcl")) {
		return "terragrunt"
	}

	if exists(filepath.Join(dir, ".terraform.lock.hcl")) {
		return "terraform"
	}

	if exists(filepath.Join(dir, ".tofu.lock.hcl")) {
		return "tofu"
	}

	return "terraform"
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
