package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscover(t *testing.T) {
	root := t.TempDir()

	mustWrite(t, filepath.Join(root, "envs/dev/app/terragrunt.hcl"))
	mustWrite(t, filepath.Join(root, "modules/vpc/main.tf"))
	mustWrite(t, filepath.Join(root, ".terraform/ignored/main.tf"))

	targets, err := Discover(root)
	if err != nil {
		t.Fatal(err)
	}

	if len(targets) != 2 {
		t.Fatalf("expected 2 targets, got %d", len(targets))
	}
}

func mustWrite(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
}
