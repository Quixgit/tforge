package runtime

import (
	"context"
	"os"

	"github.com/quix/tforge/internal/core/state"
	"github.com/quix/tforge/internal/iac/engine"
	"github.com/quix/tforge/internal/iac/plan"
	resourcesmod "github.com/quix/tforge/internal/modules/resources"
)

func Scan(ctx context.Context, eng engine.Engine, dir string) ([]resourcesmod.Row, error) {
	tmp, err := os.CreateTemp("", "tforge-*.tfplan")
	if err != nil {
		return nil, err
	}
	tmpfile := tmp.Name()
	_ = tmp.Close()
	defer os.Remove(tmpfile)

	if err := planTerraform(ctx, eng, dir, tmpfile); err != nil {
		return nil, err
	}

	raw, err := showJSON(ctx, eng, dir, tmpfile)
	if err != nil {
		return nil, err
	}

	p, err := plan.Parse(raw)
	if err != nil {
		return nil, err
	}

	rows := resourcesmod.RowsFromPlan(p)
	if len(rows) == 0 {
		rows = []resourcesmod.Row{
			{
				Kind: resourcesmod.RowResource,
				Resource: &state.Resource{
					Address: "No changes. Infrastructure is up-to-date.",
					Action:  state.ActionNoOp,
				},
			},
		}
	}

	return rows, nil
}
