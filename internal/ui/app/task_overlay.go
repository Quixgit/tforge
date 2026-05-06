package app

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderTaskOverlay(background string) string {
	title := "Running " + m.taskName
	if m.taskDone {
		title = "Finished " + m.taskName
	}

	logs := m.taskLogs
	maxLines := max(1, min(24, m.height-12))
	if len(logs) > maxLines {
		logs = logs[len(logs)-maxLines:]
	}

	content := strings.Join(logs, "\n")
	if content == "" {
		content = "Starting..."
	}

	footer := "\n\nEsc close"
	if !m.taskDone {
		footer = "\n\nRunning... | Esc hide"
	}

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(32, m.height-6)).
		Render(
			infoBarStyle.Render(title) +
				"\n\n" +
				content +
				dimStyle.Render(footer),
		)

	boxW := lipgloss.Width(box)
	boxH := lipgloss.Height(box)

	x := max(0, (m.width-boxW)/2)
	y := max(0, (m.height-boxH)/2)

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
