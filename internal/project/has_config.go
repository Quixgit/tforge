package project

import (
	"os"
	"path/filepath"
)

func HasConfig(dir string) bool {
	if fileExists(filepath.Join(dir, "terragrunt.hcl")) {
		return true
	}

	if fileExists(filepath.Join(dir, ".tofu.lock.hcl")) {
		return true
	}

	matches, err := filepath.Glob(filepath.Join(dir, "*.tf"))
	return err == nil && len(matches) > 0
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
