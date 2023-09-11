package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tree(path string, prefix string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, file := range files {
		isLast := i == len(files)-1

		if file.IsDir() {
			if isLast {
				fmt.Printf("%s└── %s\n", prefix, file.Name())
				tree(filepath.Join(path, file.Name()), prefix+"    ")
			} else {
				fmt.Printf("%s├── %s\n", prefix, file.Name())
				tree(filepath.Join(path, file.Name()), prefix+"│   ")
			}
		} else {
			if isLast {
				fmt.Printf("%s└── %s\n", prefix, file.Name())
			} else {
				fmt.Printf("%s├── %s\n", prefix, file.Name())
			}
		}
	}

	return nil
}

func main() {
	startDir := "."
	if len(os.Args) > 1 {
		startDir = os.Args[1]
	}

	absPath, err := filepath.Abs(startDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(absPath)
	tree(absPath, "")
}
