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

	command := os.Args[1]

	switch command {
	case "analyze":
		path := "."
		if len(os.Args) > 2 {
			path = os.Args[2]
		}

		fmt.Println(analyzer.Analyze(path))
	default:
		fmt.Println("Uknown command:",command)
	}
}


