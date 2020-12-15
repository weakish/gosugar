package gosugar

import "strings"

func IsUnixHiddenFile(name string) bool {
	if name == "." || name == ".." {
		return false
	} else if strings.HasPrefix(name, ".") {
		return true
	} else {
		return false
	}
}
