package app

import (
	"context"
	"fmt"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
)

func runTaskCmd(rt RuntimeInfo, action string) tea.Cmd {
	return func() tea.Msg {

		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return taskFinishedMsg{err: err}
		}

		ctx := context.Background()

		switch action {

		case "plan":
			_, err = runtime.Engine.Plan(ctx, rt.Dir)

		case "apply":
			_, err = runtime.Engine.Apply(ctx, rt.Dir)

		case "destroy":
			_, err = runtime.Engine.Destroy(ctx, rt.Dir)

		default:
			err = fmt.Errorf("unsupported action: %s", action)
		}

		return taskFinishedMsg{
			err: err,
		}
	}
}
