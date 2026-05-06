package tofu

import "github.com/quix/tforge/internal/iac/engine"

func New(binary string) engine.Engine {
	if binary == "" {
		binary = "tofu"
	}

	base := engine.NewBase("tofu", binary)
	return base
}
