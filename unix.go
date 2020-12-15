package gosugar

import (
	"fmt"
	"github.com/weakish/goaround"
	"os"
	"path/filepath"
	"strings"
)

func Quit(exitCode int, v ...interface{}) {
	_, _ = fmt.Fprint(os.Stderr, v...)
	os.Exit(exitCode)
}

func Quitf(exitCode int, format string, v ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(exitCode)
}

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
