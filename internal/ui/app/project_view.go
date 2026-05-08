package app

import "strings"

func (m Model) renderProjectView() string {
	var s strings.Builder

	title := infoBarStyle.Render(
		"TFORGE • Project Explorer",
	)

	s.WriteString(title)
	s.WriteString("\n\n")

	s.WriteString(
		dimStyle.Render("root: " + m.runtime.Root),
	)

	s.WriteString("\n\n")

	overlay := m.renderProjectOverlay("")

	s.WriteString(overlay)

	return s.String()
}
