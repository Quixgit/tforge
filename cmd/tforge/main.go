package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	ui "github.com/quix/tforge/internal/ui/app"
)

func main() {
	p := tea.NewProgram(ui.New())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "tforge failed: %v\n", err)
		os.Exit(1)
	}
}
