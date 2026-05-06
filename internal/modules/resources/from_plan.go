package resources

import (
	"sort"

	"github.com/quix/tforge/internal/core/state"
	"github.com/quix/tforge/internal/iac/plan"
)

func RowsFromPlan(p plan.Plan) []Row {
	var rows []Row

	moduleSeen := map[string]bool{}

	sort.Slice(p.ResourceChanges, func(i, j int) bool {
		return p.ResourceChanges[i].Address < p.ResourceChanges[j].Address
	})

	for _, rc := range p.ResourceChanges {

		if rc.ModuleAddress != "" && !moduleSeen[rc.ModuleAddress] {
			moduleSeen[rc.ModuleAddress] = true

			rows = append(rows, Row{
				Kind:     RowModule,
				Address:  rc.ModuleAddress,
				Expanded: true,
			})
		}

		prefix := ""

		if rc.ModuleAddress != "" {
			prefix = "├──"
		}

		r := &state.Resource{
			Address: rc.Address,
			Module:  rc.ModuleAddress,
			Type:    rc.Type,
			Name:    rc.Name,
			Action:  plan.MapActions(rc.Change.Actions),
			Reason:  rc.ActionReason,
		}

		rows = append(rows, Row{
			Kind:       RowResource,
			Address:    rc.Address,
			Parent:     rc.ModuleAddress,
			TreePrefix: prefix,
			Resource:   r,
		})
	}

	return rows
}
