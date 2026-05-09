package app

import "github.com/quix/tforge/internal/project"

func (m Model) availableActions() []string {
	if m.activeTarget != nil && m.activeTarget.Role == project.RoleModule {
		return []string{
			"validate",
			"security",
			"graph",
			"docs",
			"projects",
		}
	}

	return []string{
		"plan",
		"apply",
		"destroy",
		"refresh",
	}
}
