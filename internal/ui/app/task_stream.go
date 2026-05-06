package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
	"github.com/quix/tforge/internal/core/events"
	"github.com/quix/tforge/internal/security"
)

type taskStartedMsg struct {
	events <-chan events.Event
	err    error
}

type taskEventMsg struct {
	event events.Event
	ok    bool
}

func startTaskCmd(rt RuntimeInfo, action string) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return taskStartedMsg{err: err}
		}

		if err := security.NewPolicy(rt.AllowApply, rt.AllowDestroy).Check(action); err != nil {
			return taskStartedMsg{err: err}
		}

		ctx := context.Background()

		var ch <-chan events.Event

		switch action {
		case "plan":
			ch, err = runtime.Engine.Plan(ctx, rt.Dir)
		case "apply":
			ch, err = runtime.Engine.Apply(ctx, rt.Dir)
		case "destroy":
			ch, err = runtime.Engine.Destroy(ctx, rt.Dir)
		default:
			return taskStartedMsg{err: errUnsupportedAction(action)}
		}

		return taskStartedMsg{
			events: ch,
			err:    err,
		}
	}
}

func waitTaskEventCmd(ch <-chan events.Event) tea.Cmd {
	return func() tea.Msg {
		ev, ok := <-ch
		return taskEventMsg{
			event: ev,
			ok:    ok,
		}
	}
}
