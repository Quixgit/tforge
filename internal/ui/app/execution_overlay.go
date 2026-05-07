package app

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/quix/tforge/internal/execution"
	"github.com/quix/tforge/internal/execution/graph"
)

func (m Model) renderExecutionOverlay(background string) string {
	if m.execTracker == nil {
		return background
	}

	states := m.execTracker.List()

	sort.Slice(states, func(i, j int) bool {
		return states[i].Address < states[j].Address
	})

	completed := 0
	running := 0
	failed := 0

	for _, r := range states {
		switch r.Status {
		case execution.StatusComplete:
			completed++

		case execution.StatusRunning:
			running++

		case execution.StatusFailed:
			failed++
		}
	}

	summary := fmt.Sprintf(
		"running %d  complete %d  failed %d",
		running,
		completed,
		failed,
	)

	lines := []string{
		infoBarStyle.Render("Live Execution"),
		dimStyle.Render(summary),
		"",
		dimStyle.Render("Resource                                 Status        Duration"),
		dimStyle.Render("────────────────────────────────────────────────────────────────"),
	}

	if len(states) == 0 {
		lines = append(lines, dimStyle.Render("Waiting for terraform events..."))
	}

	nodes := graph.Build(states)

	graphText := graph.Render(nodes)

	if strings.TrimSpace(graphText) != "" {
		lines = append(lines, graphText)
		lines = append(lines, "")
	}

	lines = append(lines, dimStyle.Render("Detailed Runtime States"))
	lines = append(lines, dimStyle.Render("────────────────────────────────────────────────────────────────"))

	for _, r := range states {
		icon := "◌"
		style := dimStyle

		switch r.Status {
		case execution.StatusRunning:
			icon = "●"
			style = warningStyle

		case execution.StatusComplete:
			icon = "✔"
			style = successStyle

		case execution.StatusFailed:
			icon = "✖"
			style = errorStyle
		}

		duration := formatDuration(r.Duration())

		line := fmt.Sprintf(
			"%-40s %-12s %s",
			truncate(r.Address, 40),
			icon+" "+string(r.Status),
			duration,
		)

		lines = append(lines, style.Render(line))

		if r.Error != "" {
			lines = append(lines, errorStyle.Render("  "+r.Error))
		}
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Esc close"))

	box := focusedBorderStyle.
		Width(min(100, m.width-8)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 3, m.width, m.height)
}

func formatDuration(d time.Duration) string {
	if d < time.Second {
		return dimStyle.Render(d.String())
	}

	if d < 10*time.Second {
		return successStyle.Render(d.Round(time.Second).String())
	}

	if d < 30*time.Second {
		return warningStyle.Render(d.Round(time.Second).String())
	}

	return errorStyle.Render(d.Round(time.Second).String())
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}

	if n <= 3 {
		return s[:n]
	}

	return s[:n-3] + "..."
}
