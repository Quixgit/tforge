package graph

import (
	"strings"
)

func Render(nodes []*Node) string {
	var lines []string

	for _, n := range nodes {
		lines = append(lines, n.Label)

		for i, c := range n.Children {
			prefix := "├─ "

			if i == len(n.Children)-1 {
				prefix = "└─ "
			}

			lines = append(lines,
				prefix+
					statusIcon(c.Status)+
					" "+
					c.Label)
		}

		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}

func statusIcon(s string) string {
	switch s {

	case "running":
		return "●"

	case "complete":
		return "✔"

	case "failed":
		return "✖"

	case "planned":
		return "◌"

	default:
		return "·"
	}
}
