package gosugar

import (
	"os"
	"testing"
)

func TestIsUnixHiddenFile(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "current directory",
			input:  ".",
			output: false,
		},
		{
			name:   "parent directory",
			input:  "..",
			output: false,
		},
		{
			name:   "hidden file",
			input:  ".git",
			output: true,
		},
		{
			name:   "leading underscore",
			input:  "_config",
			output: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsUnixHiddenFile(test.input)
			if result != test.output {
				t.Errorf("Test %v, expect %v, got %v.", test.name, test.output, result)
			}
		})
	}
}

func TestSkipDirOrFile(t *testing.T) {
	f, _ := os.Stat("unix.go")
	if SkipDirOrFile(f) != nil {
		t.Errorf("Given a regular file, SkipDirOrFile should return nil.")
	}
}
