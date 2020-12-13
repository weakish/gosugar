package gosugar

import "github.com/weakish/goaround"

type StringSet map[string]UnitType

func NewStringSet() StringSet {
	return make(StringSet)
}

func (set StringSet) IsEmpty() bool {
	goaround.NotNil(set)
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
	goaround.NotNil(set)
	return set.contains(s)
}

func (set StringSet) ToSlice() []string {
	goaround.NotNil(set)
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
	goaround.NotNil(set)
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
	goaround.NotNil(set)
	return set.remove(s)
}

func (set StringSet) ContainsAll(strings []string) bool {
	goaround.NotNil(set)
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
	goaround.NotNil(set)
	return set.manipulateAll(strings, StringSet.add)
}
func (set StringSet) RemoveAll(strings []string) bool {
	goaround.NotNil(set)
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
	goaround.NotNil(set)
	var changed int = 0
	for s := range set {
		if !stringSliceContains(strings, s) {
			changed += Btoi(set.remove(s))
		}
	}
	return Itob(changed)
}

func (set StringSet) Clear() {
	goaround.NotNil(set)
	// Go compiler 1.11+ will optimize this automatically.
	for s := range set {
		delete(set, s)
	}
}

func (set StringSet) Equals(other StringSet) bool {
	goaround.NotNil(set)
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

type IntSet map[int]UnitType

func NewIntSet() IntSet {
	return make(IntSet)
}

func (set IntSet) IsEmpty() bool {
	goaround.NotNil(set)
	if len(set) == 0 {
		return true
	} else {
		return false
	}
}

func (set IntSet) contains(i int) bool {
	_, present := set[i]
	return present
}
func (set IntSet) Contains(i int) bool {
	goaround.NotNil(set)
	return set.contains(i)
}

func (set IntSet) ToSlice() []int {
	goaround.NotNil(set)
	ints := make([]int, len(set))
	for i := range set {
		ints = append(ints, i)
	}
	return ints
}

func (set IntSet) add(i int) bool {
	if set.contains(i) {
		return false
	} else {
		set[i] = Unit
		return true
	}
}
func (set IntSet) Add(i int) bool {
	goaround.NotNil(set)
	return set.add(i)
}

func (set IntSet) remove(i int) bool {
	if set.contains(i) {
		delete(set, i)
		return true
	} else {
		return false
	}
}
func (set IntSet) Remove(i int) bool {
	goaround.NotNil(set)
	return set.remove(i)
}

func (set IntSet) ContainsAll(ints []int) bool {
	goaround.NotNil(set)
	for _, i := range ints {
		if !set.contains(i) {
			return false
		}
	}
	return true
}

func (set IntSet) manipulateAll(ints []int, operation func(IntSet, int) bool) bool {
	var changed int = 0
	for _, i := range ints {
		changed += Btoi(operation(set, i))
	}
	return Itob(changed)
}
func (set IntSet) AddAll(ints []int) bool {
	goaround.NotNil(set)
	return set.manipulateAll(ints, IntSet.add)
}
func (set IntSet) RemoveAll(ints []int) bool {
	goaround.NotNil(set)
	return set.manipulateAll(ints, IntSet.remove)
}

func intSliceContains(ints []int, i int) bool {
	for _, integer := range ints {
		if i == integer {
			return true
		}
	}
	return false
}
func (set IntSet) RetainAll(ints []int) bool {
	goaround.NotNil(set)
	var changed int = 0
	for i := range set {
		if !intSliceContains(ints, i) {
			changed += Btoi(set.remove(i))
		}
	}
	return Itob(changed)
}

func (set IntSet) Clear() {
	goaround.NotNil(set)
	// Go compiler 1.11+ will optimize this automatically.
	for i := range set {
		delete(set, i)
	}
}

func (set IntSet) Equals(other IntSet) bool {
	goaround.NotNil(set)
	setSize := len(set)
	otherSize := len(other)
	if setSize != otherSize {
		return false
	} else {
		for i := range set {
			if !other.contains(i) {
				return false
			}
		}
		return true
	}
}

func (set IntSet) MigrateFrom(old map[int]bool) {
	if set.IsEmpty() {
		for i := range old {
			set[i] = Unit
		}
	} else {
		panic("The receiver of IntSet.MigrateFrom() is not empty!")
	}
}
