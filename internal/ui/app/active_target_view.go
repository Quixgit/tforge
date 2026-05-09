package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/quix/tforge/internal/project"
)

func (m Model) renderActiveTargetView() string {
	if m.activeTarget == nil {
		return m.renderListView()
	}

	if m.activeTarget.Role == project.RoleModule {
		return m.renderActiveModuleView()
	}

	return m.renderListView()
}

func (m Model) renderActiveModuleView() string {
	header := infoBarStyle.Render(
		fmt.Sprintf("TFORGE • MODULE • %s", m.activeTarget.Name),
	)

	meta := dimStyle.Render("path: " + m.activeTarget.Dir)

	leftW := 24
	rightW := max(60, m.width-leftW-8)
	paneH := max(18, m.height-10)

	left := focusedBorderStyle.
		Width(leftW).
		Height(paneH).
		Render(m.renderModuleNav())

	right := focusedBorderStyle.
		Width(rightW).
		Height(paneH).
		Render(m.renderModuleMainContent())

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)

	footer := dimStyle.Render(
		"1-5 tabs | ↑/↓ scroll | G graph | Tab actions | O projects | q quit",
	)

	content := strings.Join([]string{
		header,
		meta,
		"",
		row,
		"",
		footer,
	}, "\n")

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		content,
	)
}

func (m Model) renderModuleNav() string {
	items := []string{
		"Variables",
		"Outputs",
		"Resources",
		"Providers",
		"Graph",
		"Security",
		"Docs",
	}

	var lines []string
	lines = append(lines, infoBarStyle.Render("Module"), "")

	for i, item := range items {
		active := false

		if i < 4 && m.moduleTab == i && !m.graphMode {
			active = true
		}

		if item == "Graph" && m.graphMode {
			active = true
		}

		line := fmt.Sprintf(" %d %s", i+1, item)
		if i >= 4 {
			line = "   " + item
		}

		if active {
			lines = append(lines, cursorStyle.Render(line))
		} else {
			lines = append(lines, dimStyle.Render(line))
		}
	}

	lines = append(lines, "", dimStyle.Render("Actions"), "")
	lines = append(lines,
		"G graph",
		"O projects",
		"q quit",
	)

	return strings.Join(lines, "\n")
}

func (m Model) renderModuleMainContent() string {
	if m.graphMode {
		return m.renderGraphOverlay()
	}

	return m.renderModuleInspector()
}
