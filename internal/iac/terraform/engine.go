package terraform

import "github.com/quix/tforge/internal/iac/engine"

func New(binary string) engine.Engine {
	if binary == "" {
		binary = "terraform"
	}

	base := engine.NewBase("terraform", binary)
	return base
}
