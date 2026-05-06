package app

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderTaskOverlay(background string) string {

	title := "Running " + m.taskName

	content := strings.Join(m.taskLogs, "\n")

	if content == "" {
		content = "Starting..."
	}

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(30, m.height-6)).
		Render(
			infoBarStyle.Render(title) +
				"\n\n" +
				content,
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
