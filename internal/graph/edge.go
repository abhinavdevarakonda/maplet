package graph

type EdgeType string

const ContainsEdge EdgeType = "contains"

type Edge struct {
	From	string
	To		string
	Type	EdgeType
}
