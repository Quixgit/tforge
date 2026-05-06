package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/quix/tforge/internal/core/events"
	"github.com/quix/tforge/internal/core/state"
	"github.com/quix/tforge/internal/history"
	resources "github.com/quix/tforge/internal/modules/resources"
)

type Model struct {
	width      int
	height     int
	viewWidth  int
	viewHeight int
	cursor     int
	offset     int

	filtering bool
	filter    string
	hideNoop  bool

	actionMode   bool
	actionCursor int

	detailMode   bool
	detailScroll int

	confirmMode   bool
	confirmAction string
	confirmCursor int

	historyMode    bool
	historyDetail  bool
	historyCursor  int
	historyScroll  int
	historyErr     error
	historyEntries []history.Entry

	workspaceMode    bool
	analyticsMode    bool
	providersMode    bool
	workspaceCursor  int
	workspaceErr     error
	workspaces       []string
	currentWorkspace string

	selected map[string]bool

	taskMode   bool
	taskLogs   []string
	taskName   string
	taskDone   bool
	taskScroll int
	taskEvents <-chan events.Event

	loading bool
	err     error
	rows    []resources.Row

	runtime RuntimeInfo
}

func New() Model {
	return Model{
		selected: map[string]bool{},
	}
}

func (m Model) Init() tea.Cmd {
	return scanCmd(m.runtime)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case workspacesLoadedMsg:
		m.workspaceErr = msg.err
		m.workspaces = msg.workspaces
		m.workspaceCursor = 0

		if len(msg.workspaces) > 0 && m.currentWorkspace == "" {
			m.currentWorkspace = msg.workspaces[0]
		}

		return m, nil

	case workspaceSwitchedMsg:

		if msg.err != nil {
			m.workspaceErr = msg.err
			return m, nil
		}

		m.currentWorkspace = msg.workspace
		m.workspaceMode = false
		m.loading = true

		return m, scanCmd(m.runtime)

	case historyLoadedMsg:
		m.historyErr = msg.err
		m.historyEntries = msg.entries
		m.historyCursor = 0
		m.historyScroll = 0
		return m, nil

	case scanFinishedMsg:
		m.loading = false
		m.err = msg.err
		m.rows = msg.rows
		return m, nil

	case historySavedMsg:
		if msg.err != nil {
			m.taskLogs = append(m.taskLogs, "history save failed: "+msg.err.Error())
		}
		return m, nil

	case taskFinishedMsg:

		if msg.err != nil {
			m.taskLogs = append(m.taskLogs, "")
			m.taskLogs = append(m.taskLogs, "ERROR:")
			m.taskLogs = append(m.taskLogs, msg.err.Error())
			m.taskLogs = append(m.taskLogs, "")
			m.taskLogs = append(m.taskLogs, "Safety policy blocked this action.")
		} else {
			m.taskLogs = append(m.taskLogs, "")
			m.taskLogs = append(m.taskLogs, "Task completed successfully")
		}

		m.taskDone = true
		return m, saveHistoryCmd(m.runtime, m.taskName, m.taskLogs, true)

	case taskStartedMsg:

		if msg.err != nil {
			m.taskLogs = append(m.taskLogs, "")
			m.taskLogs = append(m.taskLogs, "ERROR:")
			m.taskLogs = append(m.taskLogs, msg.err.Error())
			m.taskDone = true
			return m, saveHistoryCmd(m.runtime, m.taskName, m.taskLogs, false)
		}

		m.taskEvents = msg.events
		m.taskLogs = append(m.taskLogs, "Process started")
		return m, waitTaskEventCmd(msg.events)

	case taskEventMsg:

		if !msg.ok {
			m.taskLogs = append(m.taskLogs, "Stream closed")
			m.taskDone = true
			return m, nil
		}

		switch msg.event.Type {
		case events.TypeStarted:
			m.taskLogs = append(m.taskLogs, "Started: "+msg.event.Command)
		case events.TypeStdout:
			if msg.event.Line != "" {
				m.taskLogs = append(m.taskLogs, msg.event.Line)
				m.taskScroll = max(0, len(m.taskLogs)-1)
			}
		case events.TypeStderr:
			if msg.event.Line != "" {
				m.taskLogs = append(m.taskLogs, "stderr: "+msg.event.Line)
				m.taskScroll = max(0, len(m.taskLogs)-1)
			}
		case events.TypeFinished:
			m.taskLogs = append(m.taskLogs, fmt.Sprintf("Finished with exit code %d", msg.event.ExitCode))
			m.taskScroll = max(0, len(m.taskLogs)-1)
			m.taskDone = true
			return m, saveHistoryCmd(m.runtime, m.taskName, m.taskLogs, msg.event.ExitCode == 0)
		case events.TypeError:
			m.taskLogs = append(m.taskLogs, "ERROR: "+msg.event.Error)
			m.taskDone = true
			return m, saveHistoryCmd(m.runtime, m.taskName, m.taskLogs, false)
		}

		if len(m.taskLogs) > 200 {
			m.taskLogs = m.taskLogs[len(m.taskLogs)-200:]
		}

		return m, waitTaskEventCmd(m.taskEvents)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewWidth = max(40, msg.Width-4)
		m.viewHeight = max(10, msg.Height-6)
		return m, nil

	case tea.KeyMsg:
		key := msg.String()

		if m.taskMode {

			switch key {

			case "esc", "q":
				m.taskMode = false

			case "up", "k":
				if m.taskScroll > 0 {
					m.taskScroll--
				}

			case "down", "j":
				if m.taskScroll < max(0, len(m.taskLogs)-1) {
					m.taskScroll++
				}

			case "pgup":
				m.taskScroll = max(0, m.taskScroll-10)

			case "pgdown":
				m.taskScroll = min(max(0, len(m.taskLogs)-1), m.taskScroll+10)

			case "home":
				m.taskScroll = 0

			case "end":
				m.taskScroll = max(0, len(m.taskLogs)-1)
			}

			return m, nil
		}

		if m.providersMode {
			switch key {
			case "esc", "p", "P", "q":
				m.providersMode = false
			}
			return m, nil
		}

		if m.analyticsMode {
			switch key {
			case "esc", "a", "A", "q":
				m.analyticsMode = false
			}
			return m, nil
		}

		if m.workspaceMode {
			switch key {
			case "esc", "w", "W":
				m.workspaceMode = false
			case "up", "k":
				if m.workspaceCursor > 0 {
					m.workspaceCursor--
				}
			case "down", "j":
				if m.workspaceCursor < len(m.workspaces)-1 {
					m.workspaceCursor++
				}
			case "enter":
				if len(m.workspaces) > 0 &&
					m.workspaceCursor < len(m.workspaces) {

					return m, switchWorkspaceCmd(
						m.runtime,
						m.workspaces[m.workspaceCursor],
					)
				}

			case "ctrl+r":
				return m, loadWorkspacesCmd(m.runtime)
			}
			return m, nil
		}

		if m.workspaceMode {
			switch key {
			case "esc", "w", "W":
				m.workspaceMode = false
			case "up", "k":
				if m.workspaceCursor > 0 {
					m.workspaceCursor--
				}
			case "down", "j":
				if m.workspaceCursor < len(m.workspaces)-1 {
					m.workspaceCursor++
				}
			case "enter":
				if len(m.workspaces) > 0 &&
					m.workspaceCursor < len(m.workspaces) {

					return m, switchWorkspaceCmd(
						m.runtime,
						m.workspaces[m.workspaceCursor],
					)
				}

			case "ctrl+r":
				return m, loadWorkspacesCmd(m.runtime)
			}
			return m, nil
		}

		if m.historyMode {
			if m.historyDetail {
				switch key {
				case "esc":
					m.historyDetail = false
				case "up", "k":
					if m.historyScroll > 0 {
						m.historyScroll--
					}
				case "down", "j":
					m.historyScroll++
				case "pgup":
					m.historyScroll = max(0, m.historyScroll-10)
				case "pgdown":
					m.historyScroll += 10
				case "home":
					m.historyScroll = 0
				}
				return m, nil
			}

			switch key {
			case "esc", "y", "Y":
				m.historyMode = false
			case "up", "k":
				if m.historyCursor > 0 {
					m.historyCursor--
				}
			case "down", "j":
				if m.historyCursor < len(m.historyEntries)-1 {
					m.historyCursor++
				}
			case "enter":
				if len(m.historyEntries) > 0 {
					m.historyDetail = true
					m.historyScroll = 0
				}
			}
			return m, nil
		}

		if m.confirmMode {
			switch key {
			case "esc", "q":
				m.confirmMode = false

			case "left", "h":
				m.confirmCursor = 0

			case "right", "l":
				m.confirmCursor = 1

			case "tab":
				if m.confirmCursor == 0 {
					m.confirmCursor = 1
				} else {
					m.confirmCursor = 0
				}

			case "enter":
				if m.confirmCursor == 0 {
					m.confirmMode = false
					return m, nil
				}

				m.confirmMode = false
				m.taskMode = true
				m.taskName = m.confirmAction
				m.taskDone = false
				m.taskScroll = 0
				m.taskLogs = []string{
					"Preparing execution...",
				}

				return m, startTaskCmd(m.runtime, m.confirmAction)
			}

			return m, nil
		}

		if m.detailMode {
			switch key {
			case "esc", "enter", "q":
				m.detailMode = false

			case "up", "k":
				if m.detailScroll > 0 {
					m.detailScroll--
				}

			case "down", "j":
				m.detailScroll++

			case "pgup":
				m.detailScroll = max(0, m.detailScroll-10)

			case "pgdown":
				m.detailScroll += 10

			case "home":
				m.detailScroll = 0
			}

			return m, nil
		}

		if m.actionMode {
			switch key {
			case "esc", "tab":
				m.actionMode = false

			case "up", "k":
				if m.actionCursor > 0 {
					m.actionCursor--
				}

			case "down", "j":
				if m.actionCursor < len(actions)-1 {
					m.actionCursor++
				}

			case "enter":

				action := actions[m.actionCursor]

				m.actionMode = false

				if action == "apply" || action == "destroy" {
					m.confirmMode = true
					m.confirmAction = action
					m.confirmCursor = 0
					return m, nil
				}

				m.taskMode = true
				m.taskName = action
				m.taskDone = false
				m.taskScroll = 0
				m.taskLogs = []string{
					"Preparing execution...",
				}

				return m, startTaskCmd(m.runtime, action)

			case "q", "ctrl+c":
				return m, tea.Quit
			}
			return m, nil
		}

		if m.filtering {
			switch key {
			case "esc", "enter":
				m.filtering = false
			case "backspace":
				if len(m.filter) > 0 {
					m.filter = m.filter[:len(m.filter)-1]
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			default:
				if len(key) == 1 {
					m.filter += key
					m.cursor = 0
					m.offset = 0
				}
			}
			return m, nil
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			rows := m.visibleRows()
			if m.cursor < len(rows)-1 {
				m.cursor++
			}

		case " ", "space":
			row := m.currentRow()
			if row != nil && row.Resource != nil {
				addr := row.Resource.Address
				m.selected[addr] = !m.selected[addr]
			}

		case "enter":
			row := m.currentRow()
			if row != nil && row.Resource != nil {
				m.detailMode = true
				m.detailScroll = 0
			}

		case "tab":
			m.actionMode = true
			m.actionCursor = 0

		case "/":
			m.filtering = true
			m.filter = ""
			m.resetCursor()

		case "h", "H":
			m.hideNoop = !m.hideNoop
			m.resetCursor()

		case "p", "P":
			m.providersMode = true

		case "a", "A":
			m.analyticsMode = true

		case "w", "W":
			m.workspaceMode = true
			return m, loadWorkspacesCmd(m.runtime)

		case "y", "Y":
			m.historyMode = true
			m.historyDetail = false
			return m, loadHistoryCmd()

		case "ctrl+r":
			m.loading = true
			m.err = nil
			return m, scanCmd(m.runtime)
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("")
	}

	if m.loading {
		return tea.NewView(lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			borderStyle.Render("Running terraform plan..."),
		))
	}

	if m.err != nil {
		return tea.NewView(lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			focusedBorderStyle.Render(errorStyle.Render("Scan failed")+"\n\n"+m.err.Error()+"\n\nPress Ctrl+r to retry | q to quit"),
		))
	}

	content := m.renderListView()

	view := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		content,
		lipgloss.WithWhitespaceChars(" "),
	)

	if m.detailMode {
		view = m.renderDetailOverlay(view)
	}

	if m.actionMode {
		view = m.renderActionModalOverlay(view)
	}

	if m.confirmMode {
		view = m.renderConfirmOverlay(view)
	}

	if m.taskMode {
		view = m.renderTaskOverlay(view)
	}

	if m.historyMode {
		view = m.renderHistoryOverlay(view)
	}

	if m.workspaceMode {
		view = m.renderWorkspaceOverlay(view)
	}

	if m.analyticsMode {
		view = m.renderAnalyticsOverlay(view)
	}

	return tea.NewView(view)
}

func (m Model) renderListView() string {
	var s strings.Builder

	fmt.Fprintln(&s, m.renderHeaderLine())
	fmt.Fprintln(&s)
	fmt.Fprint(&s, m.renderFilterBox())
	fmt.Fprintln(&s, m.renderResourcesBox())
	fmt.Fprintln(&s, m.renderInfoBar())
	s.WriteString("\n" + m.renderLegend() + "\n")
	s.WriteString(m.renderHelpBar() + "\n")

	return s.String()
}

func (m Model) renderFilterBox() string {
	filterContent := "⌕ Press '/' to filter..."

	if m.filtering {
		filterContent = "⌕ " + m.filter + "█"
	} else if m.filter != "" {
		filterContent = "⌕ " + m.filter
	}

	return borderStyle.Width(m.viewWidth).Render(filterContent) + "\n"
}

func (m Model) renderResourcesBox() string {
	rows := m.visibleRows()
	if len(rows) == 0 {
		rows = resources.DemoRows()
	}

	visible := m.visibleResourceRows()

	start := m.offset
	if start > len(rows) {
		start = 0
	}

	end := min(len(rows), start+visible)

	var b strings.Builder

	for i := start; i < end; i++ {
		switch rows[i].Kind {
		case resources.RowModule:
			fmt.Fprintln(&b, m.moduleRow(i, rows[i]))
		case resources.RowResource:
			fmt.Fprintln(&b, m.resourceRow(i, rows[i]))
		}
	}

	rendered := end - start
	for rendered < visible {
		fmt.Fprintln(&b)
		rendered++
	}

	header := ""
	if len(rows) > visible {
		header = dimStyle.Render(fmt.Sprintf(" showing %d-%d/%d", start+1, end, len(rows))) + "\n"
	}

	return resourceBorderStyle.
		Width(m.viewWidth).
		Render(strings.TrimSuffix(header+b.String(), "\n"))
}

func (m Model) resourceRow(idx int, row resources.Row) string {
	r := row.Resource
	if r == nil {
		return dimStyle.Render("No matching resources")
	}

	address := r.Address
	if r.Reason != "" {
		address += fmt.Sprintf(" (%s)", r.Reason)
	}

	prefix := treePrefixDefaultStyle.Render(row.TreePrefix)
	if row.TreePrefix != "" {
		prefix += " "
	}

	check := "[ ]"
	if m.selected[r.Address] {
		check = "[x]"
	}

	pointer := "  "
	if idx == m.cursor {
		pointer = "> "
	}

	line := strings.TrimSpace(fmt.Sprintf("%s%s %s %s", pointer, check, r.Action.Symbol(), address))

	switch {
	case idx == m.cursor:
		line = cursorStyle.Render(line)
	case m.selected[r.Address]:
		line = selectedStyle.Render(line)
	case r.Selected:
		line = selectedStyle.Render(line)
	}

	actionStyle := styleForAction(r.Action)

	if idx == m.cursor {
		line = cursorStyle.Render(line)
	} else {
		line = actionStyle.Render(line)
	}

	return prefix + line
}

func (m Model) moduleRow(idx int, row resources.Row) string {
	symbol := "▾"
	if !row.Expanded {
		symbol = "▸"
	}

	pointer := "  "
	if idx == m.cursor {
		pointer = "> "
	}

	line := fmt.Sprintf("%s%s %s", pointer, symbol, row.Address)

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

	selected := countSelected(m.selected)

	full := fmt.Sprintf("  Engine: %s", engineName)

	if m.currentWorkspace != "" {
		full += fmt.Sprintf(" | ws: %s", m.currentWorkspace)
	}

	if m.runtime.Dir != "" && m.viewWidth > 110 {
		full += fmt.Sprintf(" | dir: %s", m.runtime.Dir)
	}

	full += fmt.Sprintf(" | %d selected", selected)

	return " " + successStyle.Render("✓") + infoBarStyle.Render(full)
}

func renderKeyHint(key, desc string) string {
	key = "'" + key + "'"

	return helpKeyStyle.Render(key) +
		helpDescStyle.Render(" "+desc)
}

func (m Model) renderHelpBar() string {
	hideText := "hide unchanged"
	if m.hideNoop {
		hideText = "show unchanged"
	}

	hints := []string{
		renderKeyHint("/", "filter"),
		renderKeyHint("Space", "select"),
		renderKeyHint("Enter", "detail"),
		renderKeyHint("Tab", "action"),
		renderKeyHint("H", hideText),
		renderKeyHint("Ctrl+r", "refresh"),
		renderKeyHint("P", "providers"),
		renderKeyHint("A", "analytics"),
		renderKeyHint("W", "workspaces"),
		renderKeyHint("Y", "history"),
		renderKeyHint("q", "quit"),
	}

	line := " " + strings.Join(hints, "  ")
	if lipgloss.Width(line) <= m.viewWidth {
		return line
	}

	if m.viewWidth >= 90 {
		return " " + strings.Join(hints[:6], "  ") +
			"\n " + strings.Join(hints[6:], "  ")
	}

	short := []string{
		renderKeyHint("/", "filter"),
		renderKeyHint("Space", "select"),
		renderKeyHint("Tab", "action"),
		renderKeyHint("q", "quit"),
	}

	return " " + strings.Join(short, "  ")
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
