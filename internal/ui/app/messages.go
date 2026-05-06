package app

import resources "github.com/quix/tforge/internal/modules/resources"

type scanFinishedMsg struct {
	rows []resources.Row
	err  error
}
