package moduleparser

type Variable struct {
	Name        string
	Type        string
	Description string
	Default     string
	Required    bool
}

type Output struct {
	Name string
}

type Module struct {
	Variables []Variable
	Outputs   []Output
	Providers []string
	Resources []string
}
