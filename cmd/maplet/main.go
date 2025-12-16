package main

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/abhinavdevarakonda/maplet/internal/analyzer"
	"github.com/abhinavdevarakonda/maplet/internal/export"
	"github.com/abhinavdevarakonda/maplet/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("	maplet analyze <path>")
		fmt.Println("	maplet export <path>")
		return
	}

	command := os.Args[1]

	path := "."
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	switch command {
	case "analyze":
		g := analyzer.Analyze(path)

		fmt.Printf(
			"graph built: %d nodes, %d edges\n",
			len(g.Nodes),
			len(g.Edges),
		)

	case "export":
		g := analyzer.Analyze(path)
		eg := export.FromGraph(g)

		data, err := json.MarshalIndent(eg, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))

	case "serve":
		g := analyzer.Analyze(path)
		eg := export.FromGraph(g)

		srv := server.New(eg)
		if err := srv.Start("localhost:6767"); err != nil  {
			panic(err)
		}

	default:
		fmt.Println("uknown command:", os.Args[1])
	}
}
