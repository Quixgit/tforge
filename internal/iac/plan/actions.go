package plan

import "github.com/quix/tforge/internal/core/state"

func MapActions(actions []string) state.Action {
	if len(actions) == 0 {
		return state.ActionNoOp
	}

	if len(actions) == 2 {
		if contains(actions, "delete") && contains(actions, "create") {
			return state.ActionReplace
		}
	}

	switch actions[0] {
	case "create":
		return state.ActionCreate
	case "update":
		return state.ActionUpdate
	case "delete":
		return state.ActionDelete
	case "read":
		return state.ActionRead
	case "no-op":
		return state.ActionNoOp
	default:
		return state.ActionUncertain
	}
}

func contains(v []string, target string) bool {
	for _, s := range v {
		if s == target {
			return true
		}
	}

	return false
}
