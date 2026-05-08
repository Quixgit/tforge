package app

import "github.com/quix/tforge/internal/project"

func (m Model) currentProjectTarget() *project.Target {
	targets := m.filteredProjectTargets()

	if len(targets) == 0 || m.projectCursor >= len(targets) {
		return nil
	}

	return &targets[m.projectCursor]
}
