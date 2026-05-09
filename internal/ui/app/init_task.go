package app

import (
	"context"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
	"github.com/quix/tforge/internal/core/events"
)

func startInitCmd(rt RuntimeInfo) tea.Cmd {
	return func() tea.Msg {
		opts := appcore.Options{
			Dir:    rt.Dir,
			Engine: rt.Engine,
		}

		runtime, err := appcore.NewRuntime(opts)
		if err != nil {
			return taskStartedMsg{err: err}
		}

		err = runtime.Engine.Init(context.Background(), rt.Dir)
		if err != nil {
			return taskStartedMsg{err: err}
		}

		ch := make(chan events.Event, 3)
		ch <- events.Event{Type: events.TypeStarted, Engine: rt.Engine, Command: "init"}
		ch <- events.Event{Type: events.TypeStdout, Engine: rt.Engine, Command: "init", Line: "terraform init completed"}
		ch <- events.Event{Type: events.TypeFinished, Engine: rt.Engine, Command: "init", ExitCode: 0}
		close(ch)

		return taskStartedMsg{events: ch}
	}
}
