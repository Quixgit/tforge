package app

import "strings"

func requiresTerraformInit(msg string) bool {
	msg = strings.ToLower(msg)

	hints := []string{
		"required plugins are not installed",
		"terraform init",
		"provider plugins",
		"no package for",
		"inconsistent dependency lock file",
	}

	for _, h := range hints {
		if strings.Contains(msg, h) {
			return true
		}
	}

	return false
}
