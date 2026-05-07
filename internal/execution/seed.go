package execution

import (
	resources "github.com/quix/tforge/internal/modules/resources"
)

func (t *Tracker) SeedRows(rows []resources.Row) {
	for _, row := range rows {
		if row.Resource == nil {
			continue
		}

		addr := row.Resource.Address
		if addr == "" {
			continue
		}

		if _, ok := t.resources[addr]; ok {
			continue
		}

		t.resources[addr] = &ResourceState{
			Address: addr,
			Action:  string(row.Resource.Action),
			Status:  StatusPlanned,
		}
	}
}
