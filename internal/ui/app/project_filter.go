package app

import (
	"strings"

	"github.com/quix/tforge/internal/project"
)

func (m Model) filteredProjectTargets() []project.Target {
	if strings.TrimSpace(m.projectFilter) == "" {
		return m.projectTargets
	}

	query := strings.ToLower(strings.TrimSpace(m.projectFilter))

	var out []project.Target

	for _, t := range m.projectTargets {
		search := strings.ToLower(
			t.Name + " " +
				string(t.Kind) + " " +
				string(t.Role),
		)

		if strings.Contains(search, query) {
			out = append(out, t)
		}
	}

	return out
}
