package execution

import (
	"testing"

	"github.com/quix/tforge/internal/core/state"
	resources "github.com/quix/tforge/internal/modules/resources"
)

func TestSeedRows(t *testing.T) {
	tracker := NewTracker()

	tracker.SeedRows([]resources.Row{
		{
			Resource: &state.Resource{
				Address: "aws_s3_bucket.main",
				Action:  state.ActionCreate,
			},
		},
	})

	items := tracker.List()
	if len(items) != 1 {
		t.Fatalf("expected 1 item")
	}

	if items[0].Status != StatusPlanned {
		t.Fatalf("expected planned status")
	}
}
