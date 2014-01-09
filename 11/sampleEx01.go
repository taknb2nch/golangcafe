package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	root := "C:/golang/go/src/pkg/"

	listFiles(root, root)
}

func listFiles(rootPath, searchPath string) {
	fis, err := ioutil.ReadDir(searchPath)

	if err != nil {
		panic(err)
	}

	for _, fi := range fis {
		fullPath := filepath.Join(searchPath, fi.Name())

		if fi.IsDir() {
			listFiles(rootPath, fullPath)
		} else {
			rel, err := filepath.Rel(rootPath, fullPath)

			if err != nil {
				panic(err)
			}

			fmt.Println(rel)
		}
	}
}

