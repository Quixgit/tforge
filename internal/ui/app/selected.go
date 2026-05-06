package app

func (m Model) selectedAddresses() []string {
	rows := m.visibleRows()

	out := []string{}

	for _, row := range rows {
		if row.Resource == nil {
			continue
		}

		if m.selected[row.Resource.Address] {
			out = append(out, row.Resource.Address)
		}
	}

	return out
}
