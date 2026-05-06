package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderActionModalOverlay(background string) string {
	var b strings.Builder

	title := fmt.Sprintf("%d resource(s) selected", countSelected(m.selected))
	help := "Enter to choose | Esc to cancel"

	width := max(lipgloss.Width(title), lipgloss.Width(help)) + 8
	centered := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)

	fmt.Fprintln(&b, centered.Render(title))
	fmt.Fprintln(&b, centered.Render(strings.Repeat("─", width-6)))

	for i, action := range actions {
		line := "    " + action
		if i == m.actionCursor {
			line = cursorStyle.Render("  > " + action)
		}
		fmt.Fprintln(&b, line)
	}

	fmt.Fprintln(&b)
	fmt.Fprintln(&b, centered.Render(help))

	modal := focusedBorderStyle.Render(b.String())

	modalW := lipgloss.Width(modal)
	modalH := lipgloss.Height(modal)

	x := max(0, (m.width-modalW)/2)
	y := max(0, (m.height-modalH)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(modal).X(x).Y(y).Z(1)

	return lipgloss.NewCompositor(bg, fg).Render()
}
