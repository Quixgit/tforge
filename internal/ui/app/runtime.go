package app

import "github.com/quix/tforge/internal/project"

type RuntimeInfo struct {
	Root   string
	Dir    string
	Engine string
	Binary string

	AllowApply   bool
	AllowDestroy bool
}

func NewWithRuntime(info RuntimeInfo) Model {
	m := New()

	if info.Root == "" {
		info.Root = info.Dir
	}

	m.runtime = info
	m.selected = map[string]bool{}

	if project.HasConfig(info.Dir) {
		m.loading = true
	} else {
		m.loading = false
		m.projectMode = true
	}

	return m
}
