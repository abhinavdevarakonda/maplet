package main

import (
	"fmt"
	"os"
	"github.com/abhinavdevarakonda/maplet/internal/analyzer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: maplet analyze <project-path>")
		return
	}

	switch os.Args[1] {
	case "analyze":
		path := "."
		if (len(os.Args) > 2) {
			path = os.Args[2]
		}

		g := analyzer.Analyze(path)

		fmt.Printf(
			"graph built: %d nodes, %d edges\n",
			len(g.Nodes),
			len(g.Edges),
		)

	default:
		fmt.Println("uknown command:", os.Args[1])
	}
}
