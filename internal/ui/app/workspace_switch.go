package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
)

type workspaceSwitchedMsg struct {
	workspace string
	err       error
}

func switchWorkspaceCmd(rt RuntimeInfo, workspace string) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return workspaceSwitchedMsg{err: err}
		}

		err = runtime.Engine.SelectWorkspace(
			context.Background(),
			rt.Dir,
			workspace,
		)

		return workspaceSwitchedMsg{
			workspace: workspace,
			err:       err,
		}
	}
}
