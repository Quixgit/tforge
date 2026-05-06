package app

import resources "github.com/quix/tforge/internal/modules/resources"

func (m Model) currentRow() *resources.Row {
	rows := m.visibleRows()

	if m.cursor < 0 || m.cursor >= len(rows) {
		return nil
	}

	return &rows[m.cursor]
}
