package app

import (
	"fmt"
	"strings"

	"github.com/quix/tforge/internal/project"
)

func (m Model) renderModuleInspector() string {
	if m.moduleTarget == nil {
		return ""
	}

	var s strings.Builder

	title := fmt.Sprintf("MODULE INSPECTOR • %s", m.moduleTarget.Name)

	s.WriteString(infoBarStyle.Render(title))
	s.WriteString("\n\n")
	s.WriteString(dimStyle.Render("path: " + m.moduleTarget.Dir))
	s.WriteString("\n\n")

	tabs := []string{
		tabLabel("1", "Variables", m.moduleTab == 0),
		tabLabel("2", "Outputs", m.moduleTab == 1),
		tabLabel("3", "Resources", m.moduleTab == 2),
		tabLabel("4", "Providers", m.moduleTab == 3),
	}

	lines := []string{
		" " + strings.Join(tabs, "  "),
		"",
		fmt.Sprintf(" Variables: %d   Outputs: %d   Resources: %d   Providers: %d",
			len(m.parsedModule.Variables),
			len(m.parsedModule.Outputs),
			len(m.parsedModule.Resources),
			len(m.parsedModule.Providers),
		),
		"",
	}

	switch m.moduleTab {
	case 0:
		lines = append(lines, renderStringList("Variables", m.parsedModule.Variables)...)
	case 1:
		lines = append(lines, renderStringList("Outputs", m.parsedModule.Outputs)...)
	case 2:
		lines = append(lines, renderStringList("Resources", m.parsedModule.Resources)...)
	case 3:
		lines = append(lines, renderStringList("Providers", m.parsedModule.Providers)...)
	}

	lines = append(lines, "", dimStyle.Render("1/2/3/4 switch tabs | Esc back"))

	s.WriteString(
		focusedBorderStyle.Width(100).Render(strings.Join(lines, "\n")),
	)

	return s.String()
}

func tabLabel(key string, label string, active bool) string {
	text := key + " " + label

	if active {
		return cursorStyle.Render(text)
	}

	return dimStyle.Render(text)
}

func renderStringList(title string, items []string) []string {
	lines := []string{infoBarStyle.Render(title), ""}

	if len(items) == 0 {
		return append(lines, dimStyle.Render("No items found"))
	}

	for i, item := range items {
		if i >= 18 {
			lines = append(lines, dimStyle.Render(fmt.Sprintf("...and %d more", len(items)-i)))
			break
		}

		lines = append(lines, " • "+item)
	}

	return lines
}

func isModuleTarget(t project.Target) bool {
	return t.Role == project.RoleModule
}
