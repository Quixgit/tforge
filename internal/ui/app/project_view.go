package app

import "strings"

func (m Model) renderProjectView() string {
	var s strings.Builder

	s.WriteString(infoBarStyle.Render("TFORGE • Project Explorer"))
	s.WriteString("\n\n")
	s.WriteString(dimStyle.Render("root: " + m.runtime.Root))
	s.WriteString("\n\n")
	s.WriteString(m.renderProjectOverlay(""))

	return s.String()
}
