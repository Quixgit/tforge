package app

import (
	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/history"
)

func loadHistoryCmd() tea.Cmd {
	return func() tea.Msg {
		entries, err := history.List(50)
		return historyLoadedMsg{
			entries: entries,
			err:     err,
		}
	}
}
