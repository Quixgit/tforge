package execution

import "strings"

func IsStalePlanError(line string) bool {
	line = strings.ToLower(line)

	return strings.Contains(line, "saved plan is stale") ||
		strings.Contains(line, "plan file can no longer be applied")
}
