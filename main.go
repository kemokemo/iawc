package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

var ver bool

func init() {
	flag.BoolVar(&ver, "version", false, "display version")
	flag.BoolVar(&ver, "v", false, "display version")
	flag.Parse()
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if ver {
		fmt.Fprintf(os.Stdout, "%s version %s.%s\n", Name, Version, Revision)
		return 0
	}

	b, err := os.ReadFile("iawc.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read iawc.yaml file, %v", err)
		return 1
	}
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "please set directories to find specified words")
		return 2
	}

	words := Words{}
	err = yaml.Unmarshal(b, &words)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal Words, %v", err)
		return 1
	}

	wg := &sync.WaitGroup{}
	for _, root := range args[1:] {
		wg.Add(1)
		go walk(root, wg, os.Stdout, words)
	}
	wg.Wait()

	return 0
}

func walk(root string, wg *sync.WaitGroup, w io.Writer, words Words) {
	defer wg.Done()

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error was invoked by WalkDir: %v, %v", path, err)
		}
		if d.IsDir() {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file: %v, %v", path, err)
		}

		found := false
		for _, t := range words.Targets {
			if words.CaseSensitive {
				found = strings.Contains(string(b), t)
			} else {
				found = strings.Contains(strings.ToLower(string(b)), strings.ToLower(t))
			}

			if found {
				fmt.Fprintf(w, "%v: %v\n", path, t)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to walk files, %v", err)
	}
}
