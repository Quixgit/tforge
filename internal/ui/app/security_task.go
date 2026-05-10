package app

import (
	"context"
	"os/exec"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/core/events"
)

func startSecurityCmd(rt RuntimeInfo) tea.Cmd {
	return func() tea.Msg {
		ch := make(chan events.Event, 32)

		go func() {
			defer close(ch)

			ch <- events.Event{
				Type:    events.TypeStarted,
				Engine:  rt.Engine,
				Command: "security",
				Line:    "Starting security checks...",
			}

			if hasBinary("tflint") {
				runSecurityTool(context.Background(), ch, "tflint", rt.Dir, []string{"--chdir", rt.Dir})
			} else {
				ch <- events.Event{
					Type:    events.TypeStdout,
					Command: "security",
					Line:    "tflint not found, skipping",
				}
			}

			if hasBinary("tfsec") {
				runSecurityTool(context.Background(), ch, "tfsec", rt.Dir, []string{rt.Dir})
			} else {
				ch <- events.Event{
					Type:    events.TypeStdout,
					Command: "security",
					Line:    "tfsec not found, skipping",
				}
			}

			ch <- events.Event{
				Type:     events.TypeFinished,
				Engine:   rt.Engine,
				Command:  "security",
				ExitCode: 0,
				Line:     "Security checks finished",
			}
		}()

		return taskStartedMsg{events: ch}
	}
}

func runSecurityTool(ctx context.Context, ch chan<- events.Event, name string, dir string, args []string) {
	ch <- events.Event{
		Type:    events.TypeStdout,
		Command: "security",
		Line:    "Running " + name + "...",
	}

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		ch <- events.Event{
			Type:    events.TypeStdout,
			Command: name,
			Line:    line,
		}
	}

	if err != nil {
		ch <- events.Event{
			Type:    events.TypeStderr,
			Command: name,
			Line:    err.Error(),
		}
	}
}
