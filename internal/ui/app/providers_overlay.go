package app

import (
	"fmt"
	"sort"
	"strings"
)

func (m Model) renderProvidersOverlay(background string) string {
	counts := map[string]int{}

	for _, row := range m.visibleRows() {
		if row.Resource == nil {
			continue
		}

		provider := row.Resource.Provider
		if provider == "" {
			provider = "unknown"
		}

		counts[provider]++
	}

	providers := make([]string, 0, len(counts))
	for provider := range counts {
		providers = append(providers, provider)
	}
	sort.Strings(providers)

	lines := []string{
		infoBarStyle.Render("Providers"),
		"",
	}

	if len(providers) == 0 {
		lines = append(lines, dimStyle.Render("No providers found"))
	}

	for _, provider := range providers {
		lines = append(lines, fmt.Sprintf("%-60s %d", provider, counts[provider]))
	}

	lines = append(lines, "")
	lines = append(lines, dimStyle.Render("Esc close"))

	box := focusedBorderStyle.
		Width(min(90, m.width-10)).
		Render(strings.Join(lines, "\n"))

	return centeredLayer(background, box, 4, m.width, m.height)
}
