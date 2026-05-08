package project

import (
	"path/filepath"
	"strings"
)

func HasRunnableConfig(dir string) bool {
	if fileExists(filepath.Join(dir, "terragrunt.hcl")) {
		return true
	}

	if hasAny(dir, []string{
		"terraform.tfvars",
		"*.auto.tfvars",
		"backend.tf",
		"provider.tf",
		"providers.tf",
	}) {
		return true
	}

	if fileContains(filepath.Join(dir, "main.tf"), "backend") {
		return true
	}

	if pathLooksRunnable(dir) && hasAny(dir, []string{"*.tf"}) {
		return true
	}

	return false
}

func hasAny(dir string, patterns []string) bool {
	for _, pattern := range patterns {
		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err == nil && len(matches) > 0 {
			return true
		}
	}

	return false
}

func pathLooksRunnable(dir string) bool {
	p := strings.ToLower(filepath.ToSlash(dir))

	hints := []string{
		"/live/",
		"/env/",
		"/envs/",
		"/stacks/",
		"/deployments/",
		"/prod/",
		"/production/",
		"/stage/",
		"/staging/",
		"/dev/",
		"/qa/",
	}

	for _, hint := range hints {
		if strings.Contains(p, hint) {
			return true
		}
	}

	return false
}
