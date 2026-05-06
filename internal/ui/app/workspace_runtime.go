package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
)

func loadWorkspacesCmd(rt RuntimeInfo) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return workspacesLoadedMsg{err: err}
		}

		items, err := runtime.Engine.Workspaces(context.Background(), rt.Dir)
		return workspacesLoadedMsg{
			workspaces: items,
			err:        err,
		}
	}
}
