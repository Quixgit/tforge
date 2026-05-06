package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderDetailOverlay(background string) string {
	row := m.currentRow()

	var content string
	if row == nil || row.Resource == nil {
		content = "No resource selected"
	} else {
		r := row.Resource

		lines := []string{
			infoBarStyle.Render("Resource Detail"),
			"",
			fmt.Sprintf("Address: %s", r.Address),
			fmt.Sprintf("Action:  %s", r.Action),
			fmt.Sprintf("Symbol:  %s", r.Action.Symbol()),
			fmt.Sprintf("Module:  %s", emptyDash(r.Module)),
			fmt.Sprintf("Type:    %s", emptyDash(r.Type)),
			fmt.Sprintf("Name:    %s", emptyDash(r.Name)),
			fmt.Sprintf("Reason:  %s", emptyDash(r.Reason)),
			"",
			dimStyle.Render("Next module: before/after diff from terraform show -json"),
			"",
			dimStyle.Render("Esc close"),
		}

		content = strings.Join(lines, "\n")
	}

	box := focusedBorderStyle.
		Width(min(90, m.width-10)).
		Height(min(22, m.height-6)).
		Render(content)

	w := lipgloss.Width(box)
	h := lipgloss.Height(box)

	x := max(0, (m.width-w)/2)
	y := max(0, (m.height-h)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(box).X(x).Y(y).Z(1)

	return lipgloss.NewCompositor(bg, fg).Render()
}

func emptyDash(v string) string {
	if v == "" {
		return "-"
	}
	return v
}
