package app

import (
	"fmt"
	"strings"
)

func (m Model) renderModuleDocs() []string {
	lines := []string{
		infoBarStyle.Render("Module Documentation"),
		"",
	}

	if readme := m.readModuleReadme(); readme != "" {
		lines = append(lines, strings.Split(readme, "\n")...)
		return lines
	}

	if m.moduleTarget == nil {
		return append(lines, dimStyle.Render("No active module"))
	}

	lines = append(lines,
		fmt.Sprintf("# %s", m.moduleTarget.Name),
		"",
		dimStyle.Render(m.moduleTarget.Dir),
		"",
	)

	if len(m.parsedModule.Variables) > 0 {
		lines = append(lines,
			infoBarStyle.Render("Inputs"),
			"",
		)

		for _, v := range m.parsedModule.Variables {
			req := "optional"
			if v.Required {
				req = "required"
			}

			lines = append(lines,
				fmt.Sprintf("- %s (%s)", v.Name, req),
			)

			if v.Description != "" {
				lines = append(lines,
					"  "+v.Description,
				)
			}
		}

		lines = append(lines, "")
	}

	if len(m.parsedModule.Outputs) > 0 {
		lines = append(lines,
			infoBarStyle.Render("Outputs"),
			"",
		)

		for _, o := range m.parsedModule.Outputs {
			lines = append(lines,
				"- "+o.Name,
			)
		}
	}

	lines = append(lines,
		"",
		dimStyle.Render("terraform-docs integration planned"),
	)

	return lines
}

func renderScrollable(lines []string, offset int, height int) string {
	if len(lines) <= height {
		return strings.Join(lines, "\n")
	}

	if offset < 0 {
		offset = 0
	}

	if offset > len(lines)-height {
		offset = max(0, len(lines)-height)
	}

	end := min(len(lines), offset+height)

	visible := lines[offset:end]

	indicator := dimStyle.Render(
		fmt.Sprintf("scroll %d-%d/%d", offset+1, end, len(lines)),
	)

	return strings.Join(append(visible, "", indicator), "\n")
}
