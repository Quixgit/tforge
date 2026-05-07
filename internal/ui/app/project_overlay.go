package app

import (
	"fmt"
	"strings"
)

func (m Model) renderProjectOverlay(background string) string {
	lines := []string{
		infoBarStyle.Render("Project Targets"),
		dimStyle.Render("root: " + m.runtime.Root),
		"",
	}

	if m.projectErr != nil {
		lines = append(lines, errorStyle.Render(m.projectErr.Error()))
	}

	if len(m.projectTargets) == 0 && m.projectErr == nil {
		lines = append(lines, dimStyle.Render("No Terraform/Terragrunt/OpenTofu targets found"))
	}

	viewportH := min(20, max(1, m.height-14))

	start := 0
	if m.projectCursor >= viewportH {
		start = m.projectCursor - viewportH + 1
	}

	end := min(len(m.projectTargets), start+viewportH)

	for i := start; i < end; i++ {
		t := m.projectTargets[i]

		kind := string(t.Kind)
		line := fmt.Sprintf("%-12s %s", kind, t.Name)

		if i == m.projectCursor {
			line = cursorStyle.Render("> " + line)
		} else {
			line = "  " + line
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Enter switch | j/k move | Esc close"))

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(30, m.height-6)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 6, m.width, m.height)
}
