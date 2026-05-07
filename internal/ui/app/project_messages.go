package app

import "github.com/quix/tforge/internal/project"

type projectTargetsLoadedMsg struct {
	targets []project.Target
	err     error
}
