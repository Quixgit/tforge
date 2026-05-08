package project

import (
	"os"
	"path/filepath"
	"strings"
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
	if hasFiles(dir, []string{
		"terraform.tfvars",
		"*.auto.tfvars",
		"backend.tf",
		"provider.tf",
		"providers.tf",
		"versions.tf",
	}) {
		return true
	}

	if pathLooksLikeStack(dir) {
		return true
	}

	if fileContains(filepath.Join(dir, "main.tf"), "backend") {
		return true
	}

	if fileContains(filepath.Join(dir, "main.tf"), "required_providers") &&
		fileContains(filepath.Join(dir, "main.tf"), "provider") {
		return true
	}

	return false
}

func hasFiles(dir string, patterns []string) bool {
	for _, pattern := range patterns {
		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err == nil && len(matches) > 0 {
			return true
		}
	}

	return false
}

func pathLooksLikeStack(dir string) bool {
	normalized := strings.ToLower(filepath.ToSlash(dir))

	stackHints := []string{
		"/env/",
		"/envs/",
		"/environment/",
		"/environments/",
		"/live/",
		"/stacks/",
		"/deployments/",
		"/prod/",
		"/production/",
		"/stage/",
		"/staging/",
		"/dev/",
		"/qa/",
		"/sandbox/",
	}

	for _, hint := range stackHints {
		if strings.Contains(normalized, hint) {
			return true
		}
	}

	return false
}

func fileContains(path string, needle string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	return strings.Contains(string(data), needle)
}
