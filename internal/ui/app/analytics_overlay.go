package app

import (
	"fmt"
	"strings"

	"github.com/quix/tforge/internal/core/state"
)

func (m Model) renderAnalyticsOverlay(background string) string {
	rows := m.visibleRows()

	counts := map[state.Action]int{}
	total := 0

	for _, row := range rows {
		if row.Resource == nil {
			continue
		}

		counts[row.Resource.Action]++
		total++
	}

	lines := []string{
		infoBarStyle.Render("Analytics"),
		"",
		fmt.Sprintf("Engine:    %s", m.runtime.Engine),
		fmt.Sprintf("Workspace: %s", emptyDash(m.currentWorkspace)),
		fmt.Sprintf("Directory:  %s", m.runtime.Dir),
		"",
		fmt.Sprintf("Resources: %d", total),
		fmt.Sprintf("Selected:  %d", countSelected(m.selected)),
		"",
		successStyle.Render(fmt.Sprintf("+ create:  %d", counts[state.ActionCreate])),
		warningStyle.Render(fmt.Sprintf("~ update:  %d", counts[state.ActionUpdate])),
		errorStyle.Render(fmt.Sprintf("- delete:  %d", counts[state.ActionDelete])),
		warningStyle.Render(fmt.Sprintf("+/- replace: %d", counts[state.ActionReplace])),
		dimStyle.Render(fmt.Sprintf("no-op/read: %d", counts[state.ActionNoOp]+counts[state.ActionRead])),
		"",
		dimStyle.Render("Esc close"),
	}

	box := focusedBorderStyle.
		Width(min(90, m.width-10)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 4, m.width, m.height)
}
