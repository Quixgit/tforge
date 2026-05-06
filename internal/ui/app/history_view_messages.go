package app

import "github.com/quix/tforge/internal/history"

type historyLoadedMsg struct {
	entries []history.Entry
	err     error
}
