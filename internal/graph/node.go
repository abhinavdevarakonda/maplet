package graph

type NodeType string

// to add functions etc later
const DirectoryNode	NodeType = "directory"
const FileNode		NodeType = "file"

type Node struct {
	ID		string
	Type	NodeType
	Name	string
	Path	string
}
