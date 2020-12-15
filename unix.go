package gosugar

import (
	"github.com/weakish/goaround"
	"os"
	"path/filepath"
	"strings"
)

// IsUnixHiddenFile detects if a file name indicating it is a unix hidden file.
func IsUnixHiddenFile(name string) bool {
	if name == "." || name == ".." {
		return false
	} else if strings.HasPrefix(name, ".") {
		return true
	} else {
		return false
	}
}

// SkipDirOrFile skips the directory or regular file when walking a file tree.
func SkipDirOrFile(fileInfo os.FileInfo) error {
	goaround.NotNil(fileInfo)
	if fileInfo.IsDir() {
		return filepath.SkipDir
	} else {
		return nil
	}
}
