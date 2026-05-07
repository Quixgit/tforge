package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderConfirmOverlay(background string) string {
	selected := m.selectedAddresses()

	title := warningStyle.Render(fmt.Sprintf("Confirm %s for %d resource(s)?", m.confirmAction, len(selected)))

	lines := []string{title, ""}

	if m.confirmAction == "destroy" {
		lines = append(lines, errorStyle.Render("Destroy is disabled by default safety policy."))
		lines = append(lines, dimStyle.Render("Later we will add --allow-destroy or config-based override."))
		lines = append(lines, "")
		if m.confirmAction == "apply" {
			lines = append(lines, successStyle.Render("Apply will use the last saved plan from cache."))
			lines = append(lines, dimStyle.Render("Refresh with Ctrl+r before applying if configuration changed."))
			lines = append(lines, "")
		}
	}

	maxShow := min(10, len(selected))
	for i := 0; i < maxShow; i++ {
		lines = append(lines, "  "+selected[i])
	}

	if len(selected) > maxShow {
		lines = append(lines, dimStyle.Render(fmt.Sprintf("  ... and %d more", len(selected)-maxShow)))
	}

	lines = append(lines, "")

	cancel := confirmButton("Cancel", m.confirmCursor == 0, false)
	confirm := confirmButton("Confirm", m.confirmCursor == 1, m.confirmAction == "destroy")

	lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, cancel, "  ", confirm))
	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("←/→ switch | Enter select | Esc cancel"))

	box := focusedBorderStyle.
		Width(min(90, m.width-10)).
		Render(strings.Join(lines, "\n"))

	w := lipgloss.Width(box)
	h := lipgloss.Height(box)

	x := max(0, (m.width-w)/2)
	y := max(0, (m.height-h)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(box).X(x).Y(y).Z(2)

	return lipgloss.NewCompositor(bg, fg).Render()
}
