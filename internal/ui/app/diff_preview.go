package app

import (
	"fmt"
	"sort"
	"strings"
)

func renderDiffPreview(before, after map[string]any) string {
	if len(before) == 0 && len(after) == 0 {
		return dimStyle.Render("No before/after values available")
	}

	keysMap := map[string]bool{}

	for k := range before {
		keysMap[k] = true
	}
	for k := range after {
		keysMap[k] = true
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	lines := []string{}

	for _, key := range keys {
		bv, bok := before[key]
		av, aok := after[key]

		switch {
		case !bok && aok:
			lines = append(lines, successStyle.Render(fmt.Sprintf("+ %s = %v", key, av)))

		case bok && !aok:
			lines = append(lines, errorStyle.Render(fmt.Sprintf("- %s = %v", key, bv)))

		case fmt.Sprintf("%v", bv) != fmt.Sprintf("%v", av):
			lines = append(lines, warningStyle.Render(fmt.Sprintf("~ %s: %v -> %v", key, bv, av)))

		default:
			lines = append(lines, dimStyle.Render(fmt.Sprintf("  %s = %v", key, av)))
		}

		if len(lines) >= 16 {
			lines = append(lines, dimStyle.Render("... truncated"))
			break
		}
	}

	return strings.Join(lines, "\n")
}
