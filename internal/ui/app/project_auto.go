package app

import (
	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/project"
)

type projectAutoMsg struct {
	targets []project.Target
	err     error
}

func autoProjectCmd(root string) tea.Cmd {
	return func() tea.Msg {
		targets, err := project.Discover(root)
		return projectAutoMsg{
			targets: targets,
			err:     err,
		}
	}
}
