package gosugar

import (
	"github.com/weakish/goaround"
)

type StringSet map[string]UnitType

func NewStringSet() StringSet {
	return make(StringSet)
}

func (set StringSet) IsEmpty() bool {
	goaround.RequireNonNull(set)
	if len(set) == 0 {
		return true
	} else {
		return false
	}
}

func (set StringSet) contains(s string) bool {
	_, present := set[s]
	return present
}
func (set StringSet) Contains(s string) bool {
	goaround.RequireNonNull(set)
	return set.contains(s)
}

func (set StringSet) ToSlice() []string {
	goaround.RequireNonNull(set)
	strings := make([]string, len(set))
	for s := range set {
		strings = append(strings, s)
	}
	return strings
}

func (set StringSet) add(s string) bool {
	if set.contains(s) {
		return false
	} else {
		set[s] = Unit
		return true
	}
}
func (set StringSet) Add(s string) bool {
	goaround.RequireNonNull(set)
	return set.add(s)
}

func (set StringSet) remove(s string) bool {
	if set.contains(s) {
		delete(set, s)
		return true
	} else {
		return false
	}
}
func (set StringSet) Remove(s string) bool {
	goaround.RequireNonNull(set)
	return set.remove(s)
}

func (set StringSet) ContainsAll(strings []string) bool {
	goaround.RequireNonNull(set)
	for _, s := range strings {
		if !set.contains(s) {
			return false
		}
	}
	return true
}

func (set StringSet) manipulateAll(strings []string, operation func(StringSet, string) bool) bool {
	var changed int = 0
	for _, s := range strings {
		changed += Btoi(operation(set, s))
	}
	return Itob(changed)
}
func (set StringSet) AddAll(strings []string) bool {
	goaround.RequireNonNull(set)
	return set.manipulateAll(strings, StringSet.add)
}
func (set StringSet) RemoveAll(strings []string) bool {
	goaround.RequireNonNull(set)
	return set.manipulateAll(strings, StringSet.remove)
}

func stringSliceContains(strings []string, s string) bool {
	for _, str := range strings {
		if s == str {
			return true
		}
	}
	return false
}
func (set StringSet) RetainAll(strings []string) bool {
	goaround.RequireNonNull(set)
	var changed int = 0
	for s := range set {
		if !stringSliceContains(strings, s) {
			changed += Btoi(set.remove(s))
		}
	}
	return Itob(changed)
}

func (set StringSet) Clear() {
	goaround.RequireNonNull(set)
	// Go compiler 1.11+ will optimize this automatically.
	for s := range set {
		delete(set, s)
	}
}

func (set StringSet) Equals (other StringSet) bool {
	goaround.RequireNonNull(set)
	setSize := len(set)
	otherSize := len(other)
	if setSize != otherSize {
		return false
	} else {
		for s := range set {
			if !other.contains(s) {
				return false
			}
		}
		return true
	}
}

func (set StringSet) MigrateFrom(old map[string]bool) {
	if set.IsEmpty() {
		for s := range old {
			set[s] = Unit
		}
	} else {
		panic("The receiver of StringSet.MigrateFrom() is not empty!")
	}
}