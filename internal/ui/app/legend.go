package app

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderLegend() string {
	full := []string{
		successStyle.Render("+ create"),
		warningStyle.Render("~ update"),
		errorStyle.Render("- delete"),
		warningStyle.Render("+/- replace"),
		dimStyle.Render("no-op"),
	}

	line := " " + strings.Join(full, "  ")
	if lipgloss.Width(line) <= m.viewWidth {
		return line
	}

	short := []string{
		successStyle.Render("+"),
		warningStyle.Render("~"),
		errorStyle.Render("-"),
		warningStyle.Render("+/-"),
		dimStyle.Render("no-op"),
	}

	return " " + strings.Join(short, "  ")
}
