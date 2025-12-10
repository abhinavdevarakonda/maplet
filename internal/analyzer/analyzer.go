package analyzer

import (
	"fmt"
	"path/filepath"
	"os"
)

func ListFiles(files []string) {
	for i := range files {
		fmt.Println(files[i])
	}
}

func Analyze(path string) string {
	files := []string{}

	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		count := 0
		if err != nil {
			return nil
		}

		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(info.Name()) == ".go" {
			count++
			files = append(files,info.Name())
		}

		return nil
	})

	ListFiles(files)
	return fmt.Sprintf("Found %d files", len(files))
}
