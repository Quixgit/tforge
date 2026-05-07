package project

type Kind string

const (
	KindTerraform  Kind = "terraform"
	KindTerragrunt Kind = "terragrunt"
	KindTofu       Kind = "tofu"
)

type Target struct {
	Name string
	Dir  string
	Kind Kind
}
