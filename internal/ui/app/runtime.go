package app

type RuntimeInfo struct {
	Dir    string
	Engine string
	Binary string

	AllowApply   bool
	AllowDestroy bool
}

func NewWithRuntime(info RuntimeInfo) Model {
	m := New()

	m.runtime = info
	m.loading = true
	m.selected = map[string]bool{}

	return m
}
