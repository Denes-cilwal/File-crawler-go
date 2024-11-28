package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

//func TestRun(t *testing.T) {
//	testCases := []struct {
//		name     string
//		root     string
//		cfg      FileConfig
//		expected string
//	}{
//		{
//			name:     "NoFilter",
//			root:     "testdata",
//			cfg:      FileConfig{ext: "", size: 0, list: true},
//			expected: "testdata/dir.log\ntestdata/dir2/debug.sh\n",
//
//			/*
//				output:
//					testdata/dir.log
//					testdata/dir2/debug.sh
//			*/
//
//		},
//		{
//			name:     "FilterExtensionMatch",
//			root:     "testdata",
//			cfg:      FileConfig{ext: ".log", size: 0, list: true},
//			expected: "testdata/dir.log\n",
//
//			/*
//			 testdata/dir.log
//
//			*/
//		},
//		{
//			name:     "FilterExtensionSizeMatch",
//			root:     "testdata",
//			cfg:      FileConfig{ext: ".log", size: 10, list: true},
//			expected: "testdata/dir.log\n",
//		},
//		{
//			name:     "FilterExtensionSizeNoMatch",
//			root:     "testdata",
//			cfg:      FileConfig{ext: ".log", size: 20, list: true},
//			expected: "",
//		},
//		{
//			name:     "FilterExtensionNoMatch",
//			root:     "testdata",
//			cfg:      FileConfig{ext: ".gz", size: 0, list: true},
//			expected: "",
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			var buffer bytes.Buffer
//
//			if err := run(tc.root, &buffer, tc.cfg); err != nil {
//				t.Fatal(err)
//			}
//
//			res := buffer.String()
//
//			if tc.expected != res {
//
//				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
//			}
//		})
//	}
//}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         FileConfig
		extNoDelete string
		nDelete     int // Number of files with the extension that should be deleted.
		nNoDelete   int // Number of files with the extension that should not be deleted
		expected    string
	}{
		{name: "DeleteExtensionMixed",
			cfg:         FileConfig{ext: ".log", del: true},
			extNoDelete: ".gz", nDelete: 5, nNoDelete: 5,
			expected: ""},
	}

	// Execute RunDel test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			// Create a temporary directory with the specified number of files
			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,   // Create `nDelete` files with `.log` extension.
				tc.extNoDelete: tc.nNoDelete, // Create `nNoDelete` files with `.gz` extension.
			})
			defer cleanup() // Ensure cleanup is called after the test.
			// Run the function being tested
			if err := run(tempDir, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()

			// Check the function output
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}

			// Check remaining files in the directory
			filesLeft, err := os.ReadDir(tempDir) // Use ReadDir to list files in the directory
			if err != nil {
				t.Error(err)
			}

			fmt.Println("filesLeft: ", filesLeft)
			fmt.Println("len(filesLeft): ", len(filesLeft))
			// Assert the number of remaining files
			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("Expected %d files left, got %d instead\n",
					tc.nNoDelete, len(filesLeft))
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "crawlTest") // Create a temporary directory
	if err != nil {
		t.Fatal(err)
	}

	// Create files in the directory as specified in the `files` map
	/*
		Outer loop: ext = ".log", count = 2.
		Inner loop:
				i = 1: Creates file1.log.
				i = 2: Creates file2.log.
	*/
	for ext, count := range files {

		for i := 1; i <= count; i++ {
			fileName := fmt.Sprintf("file%d%s", i, ext) // e.g., "file1.log"
			filePath := filepath.Join(tempDir, fileName)
			if err := os.WriteFile(filePath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}

	// Return the temporary directory and a cleanup function
	return tempDir, func() {
		os.RemoveAll(tempDir) // Remove the temporary directory and its contents
	}
}
