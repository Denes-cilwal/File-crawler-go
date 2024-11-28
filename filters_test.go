package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		size     int64
		expected bool
	}{
		{
			name:     "FilterOutDir",
			file:     "testdata",
			ext:      "",
			size:     0,
			expected: true,
		},
		{
			name:     "FilterOutSize",
			file:     "testdata/dir.log",
			ext:      "",
			size:     20,
			expected: true,
		},

		{
			name:     "FilterOutExt",
			file:     "testdata/dir.log",
			ext:      ".gz",
			size:     0,
			expected: true,
		},
	}

	for _, tc := range testCases {
		info, err := os.Stat(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		actual := filterOut(tc.file, tc.ext, tc.size, info)
		if actual != tc.expected {
			t.Errorf("expected %t, got %t", tc.expected, actual)
		}
	}
}
