package app

import (
	"fmt"
	"strings"

	"github.com/quix/tforge/internal/core/state"
)

type actionStats struct {
	total   int
	create  int
	update  int
	delete  int
	replace int
}

func (m Model) stats() actionStats {
	var s actionStats

	for _, row := range m.visibleRows() {
		if row.Resource == nil {
			continue
		}

		s.total++

		switch row.Resource.Action {
		case state.ActionCreate:
			s.create++
		case state.ActionUpdate:
			s.update++
		case state.ActionDelete:
			s.delete++
		case state.ActionReplace:
			s.replace++
		}
	}

	return s
}

func (m Model) renderHeaderLine() string {
	ws := m.currentWorkspace
	if ws == "" {
		ws = "-"
	}

	s := m.stats()

	left := infoBarStyle.Render(" TFORGE ") +
		dimStyle.Render("• ") +
		successStyle.Render(m.runtime.Engine) +
		dimStyle.Render(" • ws: ") +
		warningStyle.Render(ws)

	right := strings.Join([]string{
		fmt.Sprintf("resources %d", s.total),
		successStyle.Render(fmt.Sprintf("+%d", s.create)),
		warningStyle.Render(fmt.Sprintf("~%d", s.update)),
		errorStyle.Render(fmt.Sprintf("-%d", s.delete)),
		warningStyle.Render(fmt.Sprintf("+/-%d", s.replace)),
		fmt.Sprintf("%d selected", countSelected(m.selected)),
	}, dimStyle.Render("  "))

	line := left + dimStyle.Render("  │  ") + right

	if len(line) > m.viewWidth {
		return left
	}

	return line
}
