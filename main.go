package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileConfig struct {
	ext  string
	size int64
	list bool
}

func main() {
	root := flag.String("root", ".", "Root directory to start scanning")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to search for")
	size := flag.Int64("size", 0, "Minimum file size to search for")
	flag.Parse()

	c := FileConfig{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, c FileConfig) error {
	// WalkFunc is the type of the function called by Walk to visit each file or directory
	// go as first class function which means you can pass function as argument to another function
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// walk unable to access the file or directory
			return err
		}

		if filterOut(path, c.ext, c.size, info) {
			return nil
		}

		if c.list {
			return listFile(path, out)
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}

func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}
