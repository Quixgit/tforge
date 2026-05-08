package project

import (
	"os"
	"path/filepath"
)

func Classify(dir string, kind Kind) Role {
	if kind == KindHelm {
		return RoleHelm
	}

	if kind == KindTerragrunt {
		return RoleStack
	}

	if hasRunnableHints(dir) {
		return RoleStack
	}

	return RoleModule
}

func hasRunnableHints(dir string) bool {
	patterns := []string{
		"terraform.tfvars",
		"*.auto.tfvars",
		"backend.tf",
		"provider.tf",
		"providers.tf",
	}

	for _, pattern := range patterns {
		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err == nil && len(matches) > 0 {
			return true
		}
	}

	if containsBackendBlock(filepath.Join(dir, "main.tf")) {
		return true
	}

	return false
}

func containsBackendBlock(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	return contains(string(data), "backend")
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && filepath.Base(s) != "" && stringContains(s, sub)
}

func stringContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
