package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Please set directories to find specified words")
		return 1
	}

	wg := &sync.WaitGroup{}
	for _, root := range args[1:] {
		wg.Add(1)
		go walk(root, wg, os.Stdout, []string{"名前", "What"})
	}
	wg.Wait()

	return 0
}

func walk(root string, wg *sync.WaitGroup, w io.Writer, targets []string) {
	defer wg.Done()

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file: %v, %v", path, err)
		}

		for _, t := range targets {
			if strings.Contains(string(b), t) {
				fmt.Fprintf(w, "%v: %v\n", path, t)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to walk files, %v", err)
	}
}
