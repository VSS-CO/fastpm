package resolver

type Node struct {
	Name     string
	Version  string
	Children []*Node
}
