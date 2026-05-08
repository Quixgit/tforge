package app

import (
	"fmt"
	"strings"
)

func (m Model) renderBatchOverlay(background string) string {
	lines := []string{
		infoBarStyle.Render("Batch " + m.batchAction),
		"",
	}

	if len(m.batchItems) == 0 {
		lines = append(lines, dimStyle.Render("No selected projects"))
	}

	done := 0
	for _, item := range m.batchItems {
		icon := "◌"
		style := dimStyle

		switch item.Status {
		case batchRunning:
			icon = "●"
			style = warningStyle
		case batchSuccess:
			icon = "✔"
			style = successStyle
			done++
		case batchFailed:
			icon = "✖"
			style = errorStyle
			done++
		}

		line := fmt.Sprintf("%s %-10s %s", icon, item.Status, item.Target.Name)
		lines = append(lines, style.Render(line))

		if item.Error != "" {
			lines = append(lines, errorStyle.Render("  "+item.Error))
		}
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render(fmt.Sprintf("%d/%d completed | Esc close", done, len(m.batchItems))))

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(32, m.height-6)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 7, m.width, m.height)
}
