package app

import "strings"

func projectGroup(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) == 0 || parts[0] == "" {
		return "Other"
	}

	return parts[0]
}

func projectShortName(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) <= 1 {
		return name
	}

	return strings.Join(parts[1:], "/")
}
