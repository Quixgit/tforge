package app

import (
	"fmt"
	"strings"
)

func (m Model) renderProjectOverlay(background string) string {
	selected := len(m.selectedProjects)

	lines := []string{
		infoBarStyle.Render(fmt.Sprintf("Project Targets • %d selected", selected)),
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
		marker := "[ ]"

		if m.selectedProjects[t.Dir] {
			marker = "[✓]"
		}

		role := string(t.Role)
		line := fmt.Sprintf("%s %-9s %-12s %s", marker, role, kind, t.Name)

		if i == m.projectCursor {
			line = cursorStyle.Render("> " + line)
		} else {
			line = "  " + line
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Space toggle | Enter open stack | P plan selected stacks | A apply selected stacks | Esc close"))

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(30, m.height-6)).
		Render(strings.Join(lines, "\n"))

	if background == "" {
		return box
	}

	return centeredLayer(background, box, 6, m.width, m.height)
}
