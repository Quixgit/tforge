package app

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/history"
	"github.com/quix/tforge/internal/security"
)

func saveHistoryCmd(rt RuntimeInfo, action string, logs []string, success bool) tea.Cmd {
	return func() tea.Msg {
		safeLogs := make([]string, 0, len(logs))
		for _, line := range logs {
			safeLogs = append(safeLogs, security.MaskLine(line))
		}

		err := history.Save(history.Entry{
			Time:    time.Now(),
			Dir:     rt.Dir,
			Engine:  rt.Engine,
			Action:  action,
			Success: success,
			Logs:    safeLogs,
		})

		return historySavedMsg{err: err}
	}
}
