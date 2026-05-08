package app

import (
	"strings"
)

func (m Model) renderErrorOverlay() string {
	if m.err == nil {
		return ""
	}

	lines := []string{
		"Scan failed",
		"",
		m.err.Error(),
		"",
	}

	if m.needsInit {
		lines = append(lines,
			infoBarStyle.Render("This target requires terraform init."),
			"",
			"I run terraform init",
			"Ctrl+r retry scan",
			"",
		)
	}

	lines = append(lines,
		"q quit",
	)

	return focusedBorderStyle.
		Width(min(140, m.width-10)).
		Render(strings.Join(lines, "\n"))
}
