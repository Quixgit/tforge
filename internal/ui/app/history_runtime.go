package app

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/history"
)

func saveHistoryCmd(rt RuntimeInfo, action string, logs []string, success bool) tea.Cmd {
	return func() tea.Msg {
		err := history.Save(history.Entry{
			Time:    time.Now(),
			Dir:     rt.Dir,
			Engine:  rt.Engine,
			Action:  action,
			Success: success,
			Logs:    logs,
		})

		return historySavedMsg{err: err}
	}
}
