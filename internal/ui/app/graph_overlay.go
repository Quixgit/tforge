package app

import (
	"fmt"
	"sort"
	"strings"
)

func (m Model) renderGraphOverlay() string {
	lines := []string{}

	focus := m.currentProjectTarget()

	if focus != nil {
		lines = append(lines,
			"Dependency Graph • Focus",
			"",
			focus.Name,
			"",
		)

		if node, ok := m.dependencyGraph.Nodes[focus.Name]; ok {
			if len(node.Deps) == 0 {
				lines = append(lines, "  └── no dependencies")
			}

			for _, dep := range node.Deps {
				lines = append(lines, fmt.Sprintf("  └── %s", dep))
			}
		} else {
			lines = append(lines, "  └── no graph data")
		}

		lines = append(lines, "", "G close graph")
		return focusedBorderStyle.
			Width(min(140, m.width-10)).
			Render(strings.Join(lines, "\n"))
	}

	lines = append(lines,
		"Dependency Graph",
		"",
	)

	names := make([]string, 0, len(m.dependencyGraph.Nodes))
	for name := range m.dependencyGraph.Nodes {
		names = append(names, name)
	}
	sort.Strings(names)

	for i, name := range names {
		if i >= 20 {
			lines = append(lines, dimStyle.Render(fmt.Sprintf("...and %d more", len(names)-i)))
			break
		}

		node := m.dependencyGraph.Nodes[name]
		lines = append(lines, node.Name)

		if len(node.Deps) == 0 {
			lines = append(lines, "  └── no dependencies")
		}

		for _, dep := range node.Deps {
			lines = append(lines, fmt.Sprintf("  └── %s", dep))
		}

		lines = append(lines, "")
	}

	lines = append(lines, "G close graph")

	return focusedBorderStyle.
		Width(min(140, m.width-10)).
		Render(strings.Join(lines, "\n"))
}
