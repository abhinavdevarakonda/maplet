package graph

type Graph struct {
	// here, string is just the ID/Name key in Node struct from node.go
	Nodes map[string]*Node
	Edges []*Edge
}

func New() *Graph {
	return &Graph {
		Nodes: make(map[string]*Node),
		Edges: []*Edge{},
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes[n.ID] = n
}

func (g *Graph) AddEdge(from, to string, t EdgeType) {
	g.Edges = append(g.Edges, &Edge{
		From:	from,
		To:		to,
		Type:	t,
	})
}
