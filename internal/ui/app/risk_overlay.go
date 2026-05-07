package app

import (
	"fmt"
	"strings"

	"github.com/quix/tforge/internal/risk"
)

func (m Model) renderRiskOverlay(background string) string {
	findings := risk.AnalyzeRows(m.rows)

	lines := []string{
		infoBarStyle.Render("Risk Report"),
		"",
	}

	if len(findings) == 0 {
		lines = append(lines, successStyle.Render("No risky changes detected"))
	}

	for _, f := range findings {
		style := dimStyle

		switch f.Level {
		case risk.LevelHigh:
			style = errorStyle
		case risk.LevelMedium:
			style = warningStyle
		case risk.LevelLow:
			style = successStyle
		}

		lines = append(lines, style.Render(fmt.Sprintf("[%s] %s", strings.ToUpper(string(f.Level)), f.Address)))
		lines = append(lines, dimStyle.Render("  "+f.Message))
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Esc close"))

	box := focusedBorderStyle.
		Width(min(110, m.width-10)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 5, m.width, m.height)
}
