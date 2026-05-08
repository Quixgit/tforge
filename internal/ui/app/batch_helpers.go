package app

func (m Model) selectedProjectItems() []batchItem {
	items := []batchItem{}

	for _, target := range m.projectTargets {
		if m.selectedProjects[target.Dir] {
			items = append(items, batchItem{
				Target: target,
				Status: batchPending,
			})
		}
	}

	return items
}

func (m Model) nextBatchIndex() int {
	for i, item := range m.batchItems {
		if item.Status == batchPending {
			return i
		}
	}

	return -1
}
