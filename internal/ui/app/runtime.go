package app

type RuntimeInfo struct {
	Dir    string
	Engine string
	Binary string
}

func NewWithRuntime(info RuntimeInfo) Model {
	m := New()
	m.runtime = info
	return m
}
