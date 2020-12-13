package gosugar

import (
	"github.com/weakish/goaround"
)

func requireNonEmpty(length int) {
	if length <= 0 {
		panic("Empty slice is not accepted!")
	}
}

func normalizeIndex(index int, length int) int {
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return 0
		} else {
			return negatived
		}
	} else {
		return index
	}
}

func normalizeInclusiveIndex(index int, length int) int {
	var normalizedIndex int = normalizeIndex(index, length)
	if normalizedIndex < length {
		return normalizedIndex
	} else {
		return length - 1
	}
}

func normalizeExclusiveIndex(index int, length int) int {
	var normalizedIndex int = normalizeIndex(index, length)
	if normalizedIndex < length {
		return normalizedIndex
	} else {
		return length
	}
}

// GetByte indexes string with bound safety.
func GetByte(s string, index int) byte {
	goaround.NotNil(s)
	var length int = len(s)
	requireNonEmpty(length)
	var normalizedIndex int = normalizeInclusiveIndex(index, length)
	return s[normalizedIndex]
}

// GetBool indexes slices of booleans with bound safety.
func GetBool(s []bool, index int) bool {
	goaround.NotNil(s)
	var length int = len(s)
	requireNonEmpty(length)
	var normalizedIndex int = normalizeInclusiveIndex(index, length)
	return s[normalizedIndex]
}

// GetString indexes slices of strings with bound safety.
func GetString(s []string, index int) string {
	goaround.NotNil(s)
	var length int = len(s)
	requireNonEmpty(length)
	var normalizedIndex int = normalizeInclusiveIndex(index, length)
	return s[normalizedIndex]
}

// GetInt indexes slices of integers with bound safety.
func GetInt(s []int, index int) int {
	goaround.NotNil(s)
	var length int = len(s)
	requireNonEmpty(length)
	var normalizedIndex int = normalizeInclusiveIndex(index, length)
	return s[normalizedIndex]
}

func normalizeFromTo(fromIndex int, toIndex int, length int) (int, int) {
	var normalizedFromIndex int = normalizeInclusiveIndex(fromIndex, length)
	var normalizedToIndex int = normalizeExclusiveIndex(toIndex, length)
	if normalizedFromIndex > normalizedToIndex {
		panic("fromIndex of subslice exceeds toIndex")
	} else {
		return normalizedFromIndex, normalizedToIndex
	}
}

// SubString returns a substring with bound safety.
func SubString(s string, fromIndex int, toIndex int) string {
	goaround.NotNil(s)
	var length int = len(s)
	if length == 0 {
		return ""
	} else {
		from, to := normalizeFromTo(fromIndex, toIndex, length)
		return s[from:to]
	}
}

// SubBoolSlice returns a subslice with bound safety.
func SubBoolSlice(s []bool, fromIndex int, toIndex int) []bool {
	goaround.NotNil(s)
	var length int = len(s)
	if length == 0 {
		return []bool{}
	} else {
		from, to := normalizeFromTo(fromIndex, toIndex, length)
		return s[from:to]
	}
}

// SubStringSlice returns a subslice with bound safety.
func SubStringSlice(s []string, fromIndex int, toIndex int) []string {
	goaround.NotNil(s)
	var length int = len(s)
	if length == 0 {
		return []string{}
	} else {
		from, to := normalizeFromTo(fromIndex, toIndex, length)
		return s[from:to]
	}
}

// SubIntSlice returns a subslice with bound safety.
func SubIntSlice(s []int, fromIndex int, toIndex int) []int {
	goaround.NotNil(s)
	var length int = len(s)
	if length == 0 {
		return []int{}
	} else {
		from, to := normalizeFromTo(fromIndex, toIndex, length)
		return s[from:to]
	}
}
