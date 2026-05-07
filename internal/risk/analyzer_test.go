package risk

import (
	"testing"

	"github.com/quix/tforge/internal/core/state"
	resources "github.com/quix/tforge/internal/modules/resources"
)

func TestAnalyzeRows(t *testing.T) {
	findings := AnalyzeRows([]resources.Row{
		{
			Resource: &state.Resource{
				Address: "aws_db_instance.main",
				Action:  state.ActionReplace,
			},
		},
	})

	if len(findings) != 1 {
		t.Fatalf("expected finding")
	}

	if findings[0].Level != LevelHigh {
		t.Fatalf("expected high risk")
	}
}
