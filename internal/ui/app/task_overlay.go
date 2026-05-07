package app

import (
	"fmt"
	"strings"
)

func (m Model) renderTaskOverlay(background string) string {
	title := "● Running " + m.taskName
	titleStyle := warningStyle

	if m.taskDone {
		title = "✔ Finished " + m.taskName
		titleStyle = successStyle
	}

	boxW := min(120, m.width-10)
	boxH := min(34, m.height-6)

	viewportH := max(1, boxH-8)

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

	progress := progressBar(end, len(logs), 28)
	scrollInfo := fmt.Sprintf(" lines %d-%d/%d ", start+1, end, len(logs))

	footer := "Running... | ↑/↓ scroll | Esc hide"
	if m.taskDone {
		footer = "↑/↓ scroll | PgUp/PgDn | Home/End | Esc close"
	}

	content := titleStyle.Render(title) +
		dimStyle.Render(scrollInfo) +
		"\n" +
		colorBar(progress) +
		"\n\n" +
		strings.Join(visible, "\n") +
		"\n" +
		dimStyle.Render(footer)

	box := focusedBorderStyle.
		Width(boxW).
		Height(boxH).
		Render(content)

	return centeredLayer(background, box, 3, m.width, m.height)
}

func progressBar(current, total, width int) string {
	if total <= 0 {
		total = 1
	}
	if current > total {
		current = total
	}

	filled := width * current / total
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}

	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func colorBar(v string) string {
	return successStyle.Render(v)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
