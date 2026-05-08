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

	title := fmt.Sprintf(
		"MODULE INSPECTOR • %s",
		m.moduleTarget.Name,
	)

	s.WriteString(infoBarStyle.Render(title))
	s.WriteString("\n\n")

	s.WriteString(
		dimStyle.Render("path: " + m.moduleTarget.Dir),
	)

	s.WriteString("\n\n")

	content := strings.Join([]string{
		"",
		"  Variables     coming next",
		"  Outputs       coming next",
		"  Resources     coming next",
		"  Providers     coming next",
		"  Security      coming next",
		"  Docs          coming next",
		"",
		"  This target was classified as reusable module.",
		"  Plan/apply disabled intentionally.",
		"",
		"  Esc back",
		"",
	}, "\n")

	s.WriteString(
		focusedBorderStyle.Width(90).Render(content),
	)

	return s.String()
}

func isModuleTarget(t project.Target) bool {
	return t.Role == project.RoleModule
}
