package plan

type Plan struct {
	FormatVersion   string           `json:"format_version"`
	ResourceChanges []ResourceChange `json:"resource_changes"`
}

type ResourceChange struct {
	Address       string `json:"address"`
	ModuleAddress string `json:"module_address"`

	Type string `json:"type"`
	Name string `json:"name"`

	Change Change `json:"change"`

	ActionReason string `json:"action_reason,omitempty"`
}

type Change struct {
	Actions []string       `json:"actions"`
	Before  map[string]any `json:"before"`
	After   map[string]any `json:"after"`
}
