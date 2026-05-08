package deps

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	Name string
	Path string
	Deps []string
}
