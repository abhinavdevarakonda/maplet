package analyzer

import (
	"path/filepath"
	"os"

	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

func Analyze(root string) *graph.Graph {
	g := graph.New()

	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}

			g.AddNode(&graph.Node{
				ID: p,
				Type: graph.DirectoryNode,
				Name: info.Name(),
				Path: p,
			})

			if p != root {
				parent := filepath.Dir(p)
				g.AddEdge(parent, p, graph.ContainsEdge)
			}

			return nil
		}

		if filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		g.AddNode(&graph.Node{
			ID:		p,
			Type:	graph.FileNode,
			Name:	info.Name(),
			Path:	p,
		})

		parent := filepath.Dir(p)
		g.AddEdge(parent, p, graph.ContainsEdge)

		return nil
	})

	return g
}
