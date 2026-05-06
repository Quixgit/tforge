package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
	runtimemod "github.com/quix/tforge/internal/modules/runtime"
)

func scanCmd(rt RuntimeInfo) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return scanFinishedMsg{err: err}
		}

		rows, err := runtimemod.Scan(
			context.Background(),
			runtime.Engine,
			rt.Dir,
		)

		return scanFinishedMsg{
			rows: rows,
			err:  err,
		}
	}
}
