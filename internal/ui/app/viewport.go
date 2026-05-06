package app

func (m *Model) clampCursor() {
	rows := m.visibleRows()
	if len(rows) == 0 {
		m.cursor = 0
		m.offset = 0
		return
	}

	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor > len(rows)-1 {
		m.cursor = len(rows) - 1
	}

	visible := m.visibleResourceRows()
	if visible <= 0 {
		visible = 1
	}

	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+visible {
		m.offset = m.cursor - visible + 1
	}

	if m.offset < 0 {
		m.offset = 0
	}
}

func (m Model) visibleResourceRows() int {
	return max(1, m.viewHeight-7)
}
