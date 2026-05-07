package graph

type Node struct {
	ID       string
	Label    string
	Group    string
	Status   string
	Children []*Node
}
