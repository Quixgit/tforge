package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
)

func runBatchItemCmd(rt RuntimeInfo, item batchItem, index int, action string) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    item.Target.Dir,
			Engine: string(item.Target.Kind),
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return batchItemFinishedMsg{index: index, err: err}
		}

		ctx := context.Background()

		switch action {
		case "plan":
			ch, err := runtime.Engine.Plan(ctx, item.Target.Dir)
			if err == nil {
				for range ch {
				}
			}
			return batchItemFinishedMsg{index: index, err: err}

		case "apply":
			ch, err := runtime.Engine.Apply(ctx, item.Target.Dir)
			if err == nil {
				for range ch {
				}
			}
			return batchItemFinishedMsg{index: index, err: err}
		}

		return batchItemFinishedMsg{index: index, err: errUnsupportedAction(action)}
	}
}
