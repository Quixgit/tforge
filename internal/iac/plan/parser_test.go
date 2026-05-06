package plan

import (
	"os"
	"testing"
)

func TestParsePlan(t *testing.T) {
	data, err := os.ReadFile("../../../test/fixtures/plans/simple.json")
	if err != nil {
		t.Fatal(err)
	}

	p, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	if len(p.ResourceChanges) != 2 {
		t.Fatalf("expected 2 resource changes")
	}
}
