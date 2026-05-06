package app

import (
	"fmt"

	"github.com/quix/tforge/internal/core/state"
	resources "github.com/quix/tforge/internal/modules/resources"
)

func (m Model) renderResourceRow(row resources.Row, focused bool) string {
	if row.Resource == nil {
		return ""
	}

	resource := row.Resource

	check := "[ ]"
	if m.selected[resource.Address] {
		check = "[x]"
	}

	prefix := " "
	if focused {
		prefix = ">"
	}

	actionText := "no-op"
	style := dimStyle

	switch resource.Action {
	case state.ActionCreate:
		actionText = "+ create"
		style = successStyle
	case state.ActionUpdate:
		actionText = "~ update"
		style = warningStyle
	case state.ActionDelete:
		actionText = "- delete"
		style = errorStyle
	case state.ActionReplace:
		actionText = "+/- replace"
		style = warningStyle
	}

	line := fmt.Sprintf(
		"%s %s  %-12s %s",
		prefix,
		check,
		actionText,
		resource.Address,
	)

	if focused {
		return cursorStyle.Render(line)
	}

	return style.Render(line)
}
