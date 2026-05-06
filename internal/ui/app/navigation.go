package app

func (m *Model) moveCursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
	m.clampCursor()
}

func (m *Model) moveCursorDown() {
	rows := m.visibleRows()
	if m.cursor < len(rows)-1 {
		m.cursor++
	}
	m.clampCursor()
}

func (m *Model) resetCursor() {
	m.cursor = 0
	m.offset = 0
}
