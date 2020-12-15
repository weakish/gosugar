package gosugar

import "strings"

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
