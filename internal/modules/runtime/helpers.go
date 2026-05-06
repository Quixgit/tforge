package runtime

import (
	"time"

	"github.com/quix/tforge/internal/iac/engine"
	iacruntime "github.com/quix/tforge/internal/iac/runtime"
)

func getRunner(eng engine.Engine) iacruntime.Runner {
	return iacruntime.NewRunner(30 * time.Minute)
}

func spec(
	engineName string,
	binary string,
	dir string,
	command string,
	args ...string,
) iacruntime.CommandSpec {

	return iacruntime.CommandSpec{
		Engine:  engineName,
		Binary:  binary,
		Dir:     dir,
		Command: command,
		Args:    args,
	}
}
