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
		go walk(root, wg, os.Stdout)
	}
	wg.Wait()

	return 0
}

func walk(root string, wg *sync.WaitGroup, w io.Writer) {
	defer wg.Done()

	// Current implementation doesn't seem to generate error.
	// So, ignore returned error.
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		fmt.Fprintf(w, "path: %v\n", path)
		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file: %v, %v", path, err)
		}
		i := strings.Index(string(b), "名前")
		if i != -1 {
			fmt.Fprintf(w, "%v: %v\n", path, i)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to walk files, %v", err)
	}
}
