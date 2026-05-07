package graph

import (
	"testing"

	"github.com/quix/tforge/internal/execution"
)

func TestBuild(t *testing.T) {
	nodes := Build([]execution.ResourceState{
		{
			Address: "aws_vpc.main",
			Status:  execution.StatusRunning,
		},
	})

	if len(nodes) == 0 {
		t.Fatalf("expected graph nodes")
	}
}
