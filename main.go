package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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
		go walk(root, wg, os.Stdout)
	}
	wg.Wait()

	return 0
}

func walk(root string, wg *sync.WaitGroup, w io.Writer) {
	// Current implementation doesn't seem to generate error.
	// So, ignore returned error.
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		fmt.Fprintf(w, "path: %v\n", path)
		return nil
	})
	wg.Done()
}
