package app

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderActionModal() string {

	items := make([]string, 0, len(actions))

	for i, action := range actions {

		line := "  " + action

		if i == m.actionCursor {
			line = cursorStyle.Render("> " + action)
		}

		items = append(items, line)
	}

	content := strings.Join(items, "\n")

	title := infoBarStyle.Render("Actions")

	box := focusedBorderStyle.
		Padding(1, 3).
		Render(title + "\n\n" + content)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		box,
	)
}
