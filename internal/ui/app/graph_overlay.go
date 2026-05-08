package app

import (
	"fmt"
	"strings"
)

func (m Model) renderGraphOverlay() string {
	lines := []string{
		"Dependency Graph",
		"",
	}

	count := 0

	for _, node := range m.dependencyGraph.Nodes {
		if count >= 20 {
			break
		}

		lines = append(lines,
			node.Name,
		)

		if len(node.Deps) == 0 {
			lines = append(lines,
				"  └── no dependencies",
			)
		}

		for _, dep := range node.Deps {
			lines = append(lines,
				fmt.Sprintf("  └── %s", dep),
			)
		}

		lines = append(lines, "")
		count++
	}

	lines = append(lines,
		"G close graph",
	)

	return focusedBorderStyle.
		Width(min(140, m.width-10)).
		Render(strings.Join(lines, "\n"))
}
