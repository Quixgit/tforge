package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClassifyTerraformModule(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "modules", "vpc")

	mustWriteFile(t, filepath.Join(dir, "variables.tf"), `variable "name" {}`)
	mustWriteFile(t, filepath.Join(dir, "outputs.tf"), `output "id" { value = "x" }`)
	mustWriteFile(t, filepath.Join(dir, "main.tf"), `resource "null_resource" "x" {}`)

	role := Classify(dir, KindTerraform)
	if role != RoleModule {
		t.Fatalf("expected module, got %s", role)
	}
}

func TestClassifyTerraformStackByTfvars(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "live", "prod")

	mustWriteFile(t, filepath.Join(dir, "main.tf"), `resource "null_resource" "x" {}`)
	mustWriteFile(t, filepath.Join(dir, "terraform.tfvars"), `name = "prod"`)

	role := Classify(dir, KindTerraform)
	if role != RoleStack {
		t.Fatalf("expected stack, got %s", role)
	}
}

func TestClassifyTerragruntAsStack(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "envs", "dev", "vpc")

	mustWriteFile(t, filepath.Join(dir, "terragrunt.hcl"), ``)

	role := Classify(dir, KindTerragrunt)
	if role != RoleStack {
		t.Fatalf("expected stack, got %s", role)
	}
}

func mustWriteFile(t *testing.T, path string, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
