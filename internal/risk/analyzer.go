package risk

import (
	"github.com/quix/tforge/internal/core/state"
	resources "github.com/quix/tforge/internal/modules/resources"
)

func AnalyzeRows(rows []resources.Row) []Finding {
	findings := []Finding{}

	for _, row := range rows {
		if row.Resource == nil {
			continue
		}

		r := row.Resource

		switch r.Action {
		case state.ActionDelete:
			findings = append(findings, Finding{
				Level:   LevelHigh,
				Address: r.Address,
				Action:  string(r.Action),
				Message: "resource will be deleted",
			})

		case state.ActionReplace:
			findings = append(findings, Finding{
				Level:   LevelHigh,
				Address: r.Address,
				Action:  string(r.Action),
				Message: "resource will be replaced",
			})

		case state.ActionUpdate:
			findings = append(findings, Finding{
				Level:   LevelMedium,
				Address: r.Address,
				Action:  string(r.Action),
				Message: "resource will be updated",
			})

		case state.ActionCreate:
			findings = append(findings, Finding{
				Level:   LevelLow,
				Address: r.Address,
				Action:  string(r.Action),
				Message: "resource will be created",
			})
		}
	}

	return findings
}
