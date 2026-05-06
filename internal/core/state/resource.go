package state

type Action string

const (
	ActionNoOp      Action = "no-op"
	ActionCreate    Action = "create"
	ActionUpdate    Action = "update"
	ActionDelete    Action = "delete"
	ActionReplace   Action = "replace"
	ActionRead      Action = "read"
	ActionMove      Action = "move"
	ActionImport    Action = "import"
	ActionUncertain Action = "uncertain"
)

func (a Action) Symbol() string {
	switch a {
	case ActionCreate:
		return "+"
	case ActionUpdate:
		return "~"
	case ActionDelete:
		return "-"
	case ActionReplace:
		return "+/-"
	case ActionRead:
		return "<="
	case ActionMove:
		return "→"
	case ActionImport:
		return "⇢"
	default:
		return ""
	}
}

type Resource struct {
	Address string
	Module  string
	Type    string
	Name    string
	Action  Action
	Reason  string

	Before map[string]any
	After  map[string]any

	Selected bool
}
