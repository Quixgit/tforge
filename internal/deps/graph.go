package deps

import (
	"github.com/quix/tforge/internal/project"
)

func Build(targets []project.Target) (Graph, error) {
	g := Graph{
		Nodes: map[string]*Node{},
	}

	for _, t := range targets {
		dependencies, _ := ParseDependencies(t.Dir)

		g.Nodes[t.Name] = &Node{
			Name: t.Name,
			Path: t.Dir,
			Deps: dependencies,
		}
	}

	return g, nil
}
