package app

func countSelected(selected map[string]bool) int {
	count := 0

	for _, ok := range selected {
		if ok {
			count++
		}
	}

	return count
}
