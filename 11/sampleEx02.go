package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "C:/golang/go/src/pkg/"

	err := filepath.Walk(root, 
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// 特定のディレクトリ以下を無視する場合は
				// return filepath.SkipDir
				return nil
			}

			rel, err := filepath.Rel(root, path)

			fmt.Println(rel)

			return nil
		})

	if err != nil {
		fmt.Println(1, err)
	}
}
