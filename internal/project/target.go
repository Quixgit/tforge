package project

type Kind string

const (
	KindTerraform  Kind = "terraform"
	KindTerragrunt Kind = "terragrunt"
	KindTofu       Kind = "tofu"
	KindHelm       Kind = "helm"
)

type Role string

const (
	RoleModule Role = "module"
	RoleStack  Role = "stack"
	RoleHelm   Role = "helm"
)

type Target struct {
	Name string
	Dir  string
	Kind Kind
	Role Role
}
