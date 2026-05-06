package app

import (
	"strings"
)

func (m Model) renderWorkspaceOverlay(background string) string {
	lines := []string{
		infoBarStyle.Render("Workspaces"),
		"",
	}

	if m.workspaceErr != nil {
		lines = append(lines, errorStyle.Render(m.workspaceErr.Error()))
	}

	if len(m.workspaces) == 0 && m.workspaceErr == nil {
		lines = append(lines, dimStyle.Render("No workspaces found"))
	}

	for i, ws := range m.workspaces {

		label := ws

		if ws == m.currentWorkspace {
			label += "  *"
			label = successStyle.Render(label)
		}

		line := "  " + label

		if i == m.workspaceCursor {
			line = cursorStyle.Render("> " + label)
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Enter switch | j/k move | Ctrl+r refresh | Esc close"))

	box := focusedBorderStyle.
		Width(min(80, m.width-10)).
		Height(min(24, m.height-6)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 3, m.width, m.height)
}
