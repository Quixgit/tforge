package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m Model) renderHistoryOverlay(background string) string {
	if m.historyDetail {
		return m.renderHistoryDetailOverlay(background)
	}

	lines := []string{
		infoBarStyle.Render("History"),
		"",
	}

	if m.historyErr != nil {
		lines = append(lines, errorStyle.Render(m.historyErr.Error()))
	}

	if len(m.historyEntries) == 0 && m.historyErr == nil {
		lines = append(lines, dimStyle.Render("No history yet"))
	}

	viewportH := min(18, max(1, m.height-14))

	start := 0
	if m.historyCursor >= viewportH {
		start = m.historyCursor - viewportH + 1
	}

	end := min(len(m.historyEntries), start+viewportH)

	for i := start; i < end; i++ {
		e := m.historyEntries[i]

		status := successStyle.Render("✓")
		if !e.Success {
			status = errorStyle.Render("✗")
		}

		line := fmt.Sprintf(
			"%s  %s  %-7s %-9s %s",
			status,
			e.Time.Format("2006-01-02 15:04:05"),
			e.Engine,
			e.Action,
			filepath.Base(e.Dir),
		)

		if i == m.historyCursor {
			line = cursorStyle.Render("> " + line)
		} else {
			line = "  " + line
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("j/k move | Enter logs | Esc close"))

	box := focusedBorderStyle.
		Width(min(120, m.width-10)).
		Height(min(28, m.height-6)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 2, m.width, m.height)
}

func (m Model) renderHistoryDetailOverlay(background string) string {
	if len(m.historyEntries) == 0 || m.historyCursor >= len(m.historyEntries) {
		return m.renderHistoryOverlay(background)
	}

	e := m.historyEntries[m.historyCursor]

	header := fmt.Sprintf(
		"%s | %s | %s | %s",
		e.Time.Format("2006-01-02 15:04:05"),
		e.Engine,
		e.Action,
		e.Dir,
	)

	logs := e.Logs
	viewportH := min(22, max(1, m.height-12))

	if m.historyScroll > max(0, len(logs)-1) {
		m.historyScroll = max(0, len(logs)-1)
	}

	start := m.historyScroll
	end := min(len(logs), start+viewportH)

	visible := logs[start:end]
	for len(visible) < viewportH {
		visible = append(visible, "")
	}

	content := infoBarStyle.Render("History Logs") +
		"\n" +
		dimStyle.Render(header) +
		"\n\n" +
		strings.Join(visible, "\n") +
		"\n" +
		dimStyle.Render(fmt.Sprintf("lines %d-%d/%d | ↑/↓ scroll | Esc back", start+1, end, len(logs)))

	box := focusedBorderStyle.
		Width(min(130, m.width-8)).
		Height(min(32, m.height-6)).
		Render(content)

	return centeredLayer(background, box, 3, m.width, m.height)
}

func centeredLayer(background, box string, z int, width, height int) string {
	w := lipgloss.Width(box)
	h := lipgloss.Height(box)

	x := max(0, (width-w)/2)
	y := max(0, (height-h)/2)

	bg := lipgloss.NewLayer(background)
	fg := lipgloss.NewLayer(box).X(x).Y(y).Z(z)

	return lipgloss.NewCompositor(bg, fg).Render()
}
