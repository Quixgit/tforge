package engine

import (
	"os"
	"path/filepath"
)

func Detect(dir string) string {
	if exists(filepath.Join(dir, "terragrunt.hcl")) {
		return "terragrunt"
	}

	if exists(filepath.Join(dir, ".tofu.lock.hcl")) {
		return "tofu"
	}

	if exists(filepath.Join(dir, ".terraform.lock.hcl")) {
		return "terraform"
	}

	if hasFile(dir, "*.tf") {
		return "terraform"
	}

	return "terraform"
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func hasFile(dir, pattern string) bool {
	matches, err := filepath.Glob(filepath.Join(dir, pattern))
	return err == nil && len(matches) > 0
}
