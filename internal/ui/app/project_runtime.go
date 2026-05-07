package app

import (
	tea "charm.land/bubbletea/v2"

	"github.com/quix/tforge/internal/project"
)

func loadProjectTargetsCmd(root string) tea.Cmd {
	return func() tea.Msg {
		targets, err := project.Discover(root)
		return projectTargetsLoadedMsg{
			targets: targets,
			err:     err,
		}
	}
}
