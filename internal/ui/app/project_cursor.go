package app

func (m *Model) clampProjectCursor() {
	targets := m.filteredProjectTargets()

	if len(targets) == 0 {
		m.projectCursor = 0
		return
	}

	if m.projectCursor >= len(targets) {
		m.projectCursor = len(targets) - 1
	}

	if m.projectCursor < 0 {
		m.projectCursor = 0
	}
}
