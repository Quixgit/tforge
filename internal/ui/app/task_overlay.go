package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderTaskOverlay(background string) string {
	title := "Running " + m.taskName
	if m.taskDone {
		title = "Finished " + m.taskName
	}

	boxW := min(120, m.width-10)
	boxH := min(34, m.height-6)

	viewportH := max(1, boxH-7)

	logs := m.taskLogs
	if len(logs) == 0 {
		logs = []string{"Starting..."}
	}

	start := m.taskScroll - viewportH + 1
	if start < 0 {
		start = 0
	}
	if start > len(logs) {
		start = len(logs)
	}

	end := min(len(logs), start+viewportH)

	visible := logs[start:end]

	for len(visible) < viewportH {
		visible = append(visible, "")
	}

	scrollInfo := fmt.Sprintf(" lines %d-%d/%d ", start+1, end, len(logs))

	footer := "↑/↓ scroll | PgUp/PgDn | Home/End | Esc close"
	if !m.taskDone {
		footer = "Running... | ↑/↓ scroll | Esc hide"
	}

	content := infoBarStyle.Render(title) +
		dimStyle.Render(scrollInfo) +
		"\n\n" +
		strings.Join(visible, "\n") +
		"\n" +
		dimStyle.Render(footer)

	box := focusedBorderStyle.
		Width(boxW).
		Height(boxH).
		Render(content)

	w := lipgloss.Width(box)
	h := lipgloss.Height(box)

	x := max(0, (m.width-w)/2)
	y := max(0, (m.height-h)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(box).X(x).Y(y).Z(1)

	return lipgloss.NewCompositor(bg, fg).Render()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
