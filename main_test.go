package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      FileConfig
		expected string
	}{
		{
			name:     "NoFilter",
			root:     "testdata",
			cfg:      FileConfig{ext: "", size: 0, list: true},
			expected: "testdata/dir.log\ntestdata/dir2/debug.sh\n",

			/*
				output:
					testdata/dir.log
					testdata/dir2/debug.sh
			*/

		},
		{
			name:     "FilterExtensionMatch",
			root:     "testdata",
			cfg:      FileConfig{ext: ".log", size: 0, list: true},
			expected: "testdata/dir.log\n",

			/*
			 testdata/dir.log

			*/
		},
		{
			name:     "FilterExtensionSizeMatch",
			root:     "testdata",
			cfg:      FileConfig{ext: ".log", size: 10, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name:     "FilterExtensionSizeNoMatch",
			root:     "testdata",
			cfg:      FileConfig{ext: ".log", size: 20, list: true},
			expected: "",
		},
		{
			name:     "FilterExtensionNoMatch",
			root:     "testdata",
			cfg:      FileConfig{ext: ".gz", size: 0, list: true},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()

			if tc.expected != res {

				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
