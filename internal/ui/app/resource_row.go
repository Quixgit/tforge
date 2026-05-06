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

	r := row.Resource

	check := "[ ]"
	if m.selected[r.Address] {
		check = "[x]"
	}

	pointer := " "
	if focused {
		pointer = ">"
	}

	actionText := "no-op"
	style := dimStyle

	switch r.Action {
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

	line := fmt.Sprintf("%s %s  %-12s %s", pointer, check, actionText, r.Address)

	if focused {
		return cursorStyle.Render(line)
	}

	if m.selected[r.Address] {
		return selectedStyle.Render(line)
	}

	return style.Render(line)
}
