package app

import "strings"

func renderLegend() string {
	items := []string{
		successStyle.Render("+ create"),
		warningStyle.Render("~ update"),
		errorStyle.Render("- delete"),
		warningStyle.Render("+/- replace"),
		dimStyle.Render("no-op"),
	}

	return " " + strings.Join(items, "  ")
}
