package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	rows := []string{
		m.resourceLine(0, "", "+/-", "aws_db_instance.main (cannot_update)", warningStyle),
		m.resourceLine(1, "", "~", "aws_s3_bucket.assets", warningStyle),
		m.resourceLine(2, "", "~", "aws_s3_bucket.logs", warningStyle),
		m.resourceLine(3, "", "", "aws_s3_bucket.uploads", lipgloss.NewStyle()),
		m.resourceLine(4, "", "", "data.aws_region.current", lipgloss.NewStyle()),

		m.moduleLine(5, "", "module.api", true),

		m.resourceLine(6, "├──", "", "module.api.aws_cloudwatch_log_group.api", lipgloss.NewStyle()),
		m.resourceLine(7, "├──", "+", "module.api.aws_cloudwatch_log_group.api_v2", successStyle),
		m.resourceLine(8, "├──", "", "module.api.aws_iam_policy.lambda_exec", lipgloss.NewStyle()),
		m.resourceLine(9, "├──", "-", "module.api.aws_iam_role.api_lambda (delete_because_no_resource_config)", errorStyle),
		m.resourceLine(10, "├──", "~", "module.api.aws_lambda_function.api", warningStyle),
		m.resourceLine(11, "└──", "", "module.api.aws_route53_record.api", lipgloss.NewStyle()),

		m.moduleLine(12, "", "module.networking", true),

		m.resourceLine(13, "├──", "~", "module.networking.aws_security_group.web", warningStyle),
		m.resourceLine(14, "└──", "", "module.networking.aws_vpc.main", lipgloss.NewStyle()),
	}

	visible := max(1, m.viewHeight-7)

	var b strings.Builder

	for i := 0; i < visible; i++ {
		if i < len(rows) {
			fmt.Fprintln(&b, rows[i])
		} else {
			fmt.Fprintln(&b)
		}
	}

	return resourceBorderStyle.
		Width(m.viewWidth).
		Render(strings.TrimSuffix(b.String(), "\n"))
}

func (m Model) resourceLine(
	idx int,
	prefix string,
	action string,
	address string,
	actionStyle lipgloss.Style,
) string {

	prefixRendered := treePrefixDefaultStyle.Render(prefix)

	if prefix != "" {
		prefixRendered += " "
	}

	line := strings.TrimSpace(fmt.Sprintf("%s %s", action, address))

	switch {

	case idx == m.cursor:
		line = cursorStyle.Render(line)

	case idx == 1 || idx == 2 || idx == 10 || idx == 13:
		line = selectedStyle.Render(line)
	}

	if actionStyle.GetForeground() != nil {
		line = actionStyle.Render(line)
	}

	return prefixRendered + line
}

func (m Model) moduleLine(
	idx int,
	prefix string,
	name string,
	expanded bool,
) string {

	symbol := "▾"

	if !expanded {
		symbol = "▸"
	}

	line := fmt.Sprintf("%s %s", symbol, name)

	if idx == m.cursor {
		line = cursorStyle.Render(line)
	} else {
		line = moduleStyle.Render(line)
	}

	return treePrefixCurrentStyle.Render(prefix) + line
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
