package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	appcore "github.com/quix/tforge/internal/app"
	ui "github.com/quix/tforge/internal/ui/app"
)

func main() {
	opts, err := appcore.ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "tforge: %v\n", err)
		os.Exit(2)
	}

	rt, err := appcore.NewRuntime(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tforge: %v\n", err)
		os.Exit(2)
	}

	model := ui.NewWithRuntime(ui.RuntimeInfo{
		Dir:    rt.Options.Dir,
		Engine: rt.Engine.Name(),
		Binary: rt.Engine.Binary(),

		AllowApply:   rt.Config.Security.AllowApply,
		AllowDestroy: rt.Config.Security.AllowDestroy,
	})

	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "tforge failed: %v\n", err)
		os.Exit(1)
	}
}
