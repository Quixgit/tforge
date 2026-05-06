package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/quix/tforge/internal/core/state"
	resourcesmod "github.com/quix/tforge/internal/modules/resources"
)

type Model struct {
	width      int
	height     int
	viewWidth  int
	viewHeight int
	cursor     int

	runtime RuntimeInfo
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewWidth = max(40, msg.Width-4)
		m.viewHeight = max(10, msg.Height-6)

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < 14 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("")
	}

	content := m.renderListView()

	return tea.NewView(
		lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Top,
			content,
			lipgloss.WithWhitespaceChars(" "),
		),
	)
}

func (m Model) renderListView() string {
	var s strings.Builder

	fmt.Fprint(&s, m.renderFilterBox())
	fmt.Fprintln(&s, m.renderResourcesBox())
	fmt.Fprintln(&s, m.renderInfoBar())
	s.WriteString("\n" + m.renderHelpBar() + "\n")

	return s.String()
}

func (m Model) renderFilterBox() string {
	filterContent := "⌕ Press '/' to filter..."
	return borderStyle.Width(m.viewWidth).Render(filterContent) + "\n"
}

func (m Model) renderResourcesBox() string {
	rows := resourcesmod.DemoRows()

	// TODO: replace with real runtime state

	visible := max(1, m.viewHeight-7)

	var b strings.Builder

	for i := 0; i < visible; i++ {
		if i < len(rows) {
			switch rows[i].Kind {
			case resourcesmod.RowModule:
				fmt.Fprintln(&b, m.moduleRow(i, rows[i]))
			case resourcesmod.RowResource:
				fmt.Fprintln(&b, m.resourceRow(i, rows[i]))
			}
		} else {
			fmt.Fprintln(&b)
		}
	}

	return resourceBorderStyle.
		Width(m.viewWidth).
		Render(strings.TrimSuffix(b.String(), "\n"))
}

func (m Model) resourceRow(idx int, row resourcesmod.Row) string {
	r := row.Resource
	if r == nil {
		return ""
	}

	address := r.Address
	if r.Reason != "" {
		address += fmt.Sprintf(" (%s)", r.Reason)
	}

	prefix := treePrefixDefaultStyle.Render(row.TreePrefix)
	if row.TreePrefix != "" {
		prefix += " "
	}

	line := strings.TrimSpace(fmt.Sprintf("%s %s", r.Action.Symbol(), address))

	switch {
	case idx == m.cursor:
		line = cursorStyle.Render(line)
	case r.Selected:
		line = selectedStyle.Render(line)
	}

	line = styleForAction(r.Action).Render(line)

	return prefix + line
}

func (m Model) moduleRow(idx int, row resourcesmod.Row) string {
	symbol := "▾"

	if !row.Expanded {
		symbol = "▸"
	}

	line := fmt.Sprintf("%s %s", symbol, row.Address)

	if idx == m.cursor {
		line = cursorStyle.Render(line)
	} else {
		line = moduleStyle.Render(line)
	}

	return treePrefixCurrentStyle.Render(row.TreePrefix) + line
}

func styleForAction(action state.Action) lipgloss.Style {
	switch action {
	case state.ActionCreate:
		return successStyle
	case state.ActionDelete:
		return errorStyle
	case state.ActionUpdate, state.ActionReplace:
		return warningStyle
	case state.ActionMove, state.ActionImport:
		return lipgloss.NewStyle().Foreground(colorBlue)
	case state.ActionUncertain:
		return dimStyle
	default:
		return lipgloss.NewStyle()
	}
}

func (m Model) renderInfoBar() string {
	engineName := m.runtime.Engine
	if engineName == "" {
		engineName = "mock"
	}

	info := fmt.Sprintf("  Engine: %s", engineName)

	if m.runtime.Dir != "" {
		info += fmt.Sprintf(" | dir: %s", m.runtime.Dir)
	}

	info += " | mock scan complete | 4 selected | 1 warnings"

	return " " +
		successStyle.Render("✓") +
		infoBarStyle.Render(info)
}

func renderKeyHint(key, desc string) string {
	key = "'" + key + "'"

	return helpKeyStyle.Render(key) +
		helpDescStyle.Render(" "+desc)
}

func (m Model) renderHelpBar() string {
	hints := []string{
		renderKeyHint("/", "filter"),
		renderKeyHint("Space", "select"),
		renderKeyHint("Enter", "detail"),
		renderKeyHint("Tab", "action"),
		renderKeyHint("H", "hide unchanged"),
		renderKeyHint("Ctrl+r", "refresh"),
		renderKeyHint("q", "quit"),
	}

	if m.viewWidth >= 90 {
		return " " + strings.Join(hints, "  ")
	}

	mid := (len(hints) + 1) / 2

	return " " +
		strings.Join(hints[:mid], "  ") +
		"\n " +
		strings.Join(hints[mid:], "  ")
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
