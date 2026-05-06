package app

import (
	"strings"

	resources "github.com/quix/tforge/internal/modules/resources"
)

func (m Model) visibleRows() []resources.Row {
	rows := make([]resources.Row, 0, len(m.rows))

	for _, row := range m.rows {

		if row.Kind == resources.RowResource && row.Resource != nil {

			if m.hideNoop && row.Resource.Action.Symbol() == "" {
				continue
			}

			if m.filter != "" {
				value := strings.ToLower(row.Resource.Address)

				if !strings.Contains(value, strings.ToLower(m.filter)) {
					continue
				}
			}
		}

		rows = append(rows, row)
	}

	return rows
}
