package graph

import (
	"sort"
	"strings"

	"github.com/quix/tforge/internal/execution"
)

func Build(states []execution.ResourceState) []*Node {
	groups := map[string][]execution.ResourceState{}

	for _, s := range states {
		group := detectGroup(s.Address)
		groups[group] = append(groups[group], s)
	}

	groupNames := make([]string, 0, len(groups))

	for g := range groups {
		groupNames = append(groupNames, g)
	}

	sort.Strings(groupNames)

	var out []*Node

	for _, g := range groupNames {
		parent := &Node{
			ID:    g,
			Label: g,
			Group: g,
		}

		for _, r := range groups[g] {
			parent.Children = append(parent.Children, &Node{
				ID:     r.Address,
				Label:  r.Address,
				Group:  g,
				Status: string(r.Status),
			})
		}

		out = append(out, parent)
	}

	return out
}

func detectGroup(addr string) string {
	parts := strings.Split(addr, ".")

	if len(parts) == 0 {
		return "root"
	}

	if strings.HasPrefix(addr, "module.") && len(parts) > 1 {
		return parts[1]
	}

	if strings.Contains(addr, "_") {
		p := strings.Split(parts[0], "_")
		if len(p) > 0 {
			return p[0]
		}
	}

	return "root"
}
