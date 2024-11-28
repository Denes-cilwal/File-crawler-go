package main

import (
	"os"
	"path/filepath"
)

func filterOut(path, ext string, minSize int64, fileInfo os.FileInfo) bool {
	// fileInfo contains metadata about the file
	if fileInfo.IsDir() || fileInfo.Size() < minSize {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false

}
