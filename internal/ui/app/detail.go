package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderDetailOverlay(background string) string {
	lines := m.detailLines()

	boxW := min(110, m.width-10)
	boxH := min(32, m.height-6)

	viewportH := max(1, boxH-4)

	if m.detailScroll > max(0, len(lines)-1) {
		m.detailScroll = max(0, len(lines)-1)
	}

	start := m.detailScroll
	end := min(len(lines), start+viewportH)

	visible := lines[start:end]
	for len(visible) < viewportH {
		visible = append(visible, "")
	}

	scrollInfo := fmt.Sprintf(" lines %d-%d/%d ", start+1, end, len(lines))

	content := strings.Join(visible, "\n") +
		"\n" +
		dimStyle.Render(scrollInfo+"↑/↓ scroll | PgUp/PgDn | Home | Esc close")

	box := focusedBorderStyle.
		Width(boxW).
		Height(boxH).
		Render(content)

	w := lipgloss.Width(box)
	h := lipgloss.Height(box)

	x := max(0, (m.width-w)/2)
	y := max(0, (m.height-h)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(box).X(x).Y(y).Z(1)

	return lipgloss.NewCompositor(bg, fg).Render()
}

func (m Model) detailLines() []string {
	row := m.currentRow()

	if row == nil || row.Resource == nil {
		return []string{"No resource selected"}
	}

	r := row.Resource

	lines := []string{
		infoBarStyle.Render("Resource Detail"),
		"",
		fmt.Sprintf("Address: %s", r.Address),
		fmt.Sprintf("Action:  %s", styleForAction(r.Action).Render(string(r.Action))),
		fmt.Sprintf("Symbol:  %s", styleForAction(r.Action).Render(r.Action.Symbol())),
		fmt.Sprintf("Module:  %s", emptyDash(r.Module)),
		fmt.Sprintf("Type:    %s", emptyDash(r.Type)),
		fmt.Sprintf("Name:    %s", emptyDash(r.Name)),
		fmt.Sprintf("Reason:  %s", emptyDash(r.Reason)),
		"",
		infoBarStyle.Render("Changes"),
	}

	diff := renderDiffPreview(r.Before, r.After)
	if diff != "" {
		lines = append(lines, strings.Split(diff, "\n")...)
	}

	lines = append(lines, "", dimStyle.Render("Esc close"))

	return lines
}

func emptyDash(v string) string {
	if v == "" {
		return "-"
	}
	return v
}
