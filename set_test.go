package gosugar

import (
	"fmt"
	"testing"
)

func ExampleNewStringSet() {
	fmt.Println(NewStringSet())
	// Output:
	// map[]
}
func ExampleNewIntSet() {
	fmt.Println(NewIntSet())
	// Output:
	// map[]
}

func TestStringSet_Add(t *testing.T) {
	set := NewStringSet()

	if !set.Add("goaround") {
		t.Errorf("An empty set adding a string should return true, not false!")
	}
	if !set.Contains("goaround") {
		t.Errorf("%v should contain added string 'goaround', but it does not!", set)
	}
	if set.Add("goaround") {
		t.Errorf("%v already contains 'around', adding it again should return false, not true!", set)
	}
}

func TestStringSet_AddAll(t *testing.T) {
	set := NewStringSet()

	if !set.AddAll([]string{"go", "around"}) {
		t.Errorf("An empty set adding two strings should return true, not false!")
	}

	if !set.Contains("go") {
		t.Errorf("%v should contain 'go', but it does not!", set)
	}
	if !set.Contains("around") {
		t.Errorf("%v should contain 'around', but it does not!", set)
	}

	if set.AddAll([]string{}) {
		t.Errorf("Adding an empty slice to %v should return false, not true!", set)
	}
	if set.AddAll([]string{"go", "around"}) {
		t.Errorf("Adding a slice with old strings to %v should return false, not true!", set)
	}
	if !set.AddAll([]string{"go", "sugar"}) {
		t.Errorf("Adding a slice with old and new strings to %v should return true, not false!", set)
	}
}

func TestStringSet_Clear(t *testing.T) {
	set := NewStringSet()
	set.Add("goaround")
	set.Clear()
	if !set.IsEmpty() {
		t.Errorf("After clearing, %v is still not empty!", set)
	}
}

func TestStringSet_Contains(t *testing.T) {
	set := NewStringSet()
	set.Add("goaround")
	if set["goaround"] != Unit {
		t.Errorf("%v should contain 'goaround', but it does not!", set)
	}
}

func TestStringSet_ContainsAll(t *testing.T) {
	set := NewStringSet()
	stringSlice := []string{"go", "around"}
	set.AddAll(stringSlice)
	for _, s := range stringSlice {
		if !set.Contains(s) {
			t.Errorf("%v should contain %v, but it does not!", set, s)
		}
	}
}

func TestStringSet_Equals(t *testing.T) {
	set := NewStringSet()
	stringSlice := []string{"go", "around"}
	set.AddAll(stringSlice)

	other := NewStringSet()
	other.Add(stringSlice[0])
	other.Add(stringSlice[1])

	if !set.Equals(other) {
		t.Errorf("%v and %v should be equal, but they are not!", set, other)
	}

	set.Clear()
	emptySet := NewStringSet()
	if !set.Equals(emptySet) {
		t.Errorf("Both %v and %v should be empty and equal, but they are not!", set, emptySet)
	}
	if emptySet.Equals(other) {
		t.Errorf("Empty set %v and non-empty set %v should not equal to each other, but they do!", emptySet, other)
	}
}

func TestStringSet_IsEmpty(t *testing.T) {
	set := NewStringSet()
	if !set.IsEmpty() {
		t.Errorf("%v should be empty, but it is not!", set)
	}
	set.Add("goaround")
	if set.IsEmpty() {
		t.Errorf("%v should not be empty, but it is!", set)
	}
}

func TestStringSet_MigrateFrom(t *testing.T) {
	oldSet := make(map[string]bool)
	oldSet["go"] = true
	oldSet["around"] = false

	newSet := NewStringSet()
	newSet.MigrateFrom(oldSet)

	set := NewStringSet()
	set.AddAll([]string{"go", "around"})

	if !newSet.Equals(set) {
		t.Errorf("%v and %v should equal to each other, but they don't!", newSet, set)
	}
}

func TestStringSet_Remove(t *testing.T) {
	set := NewStringSet()
	set.AddAll([]string{"go", "around"})
	if !set.Remove("go") {
		t.Errorf("Removing existing string 'go' should return true, but it returns false!")
	}
	if set.Contains("go") {
		t.Errorf("%v should not contain removed string, but it does!", set)
	}
	if !set.Contains("around") {
		t.Errorf("%v should contain 'around', but it does!", set)
	}
}

func TestStringSet_RemoveAll(t *testing.T) {
	set := NewStringSet()
	set.AddAll([]string{"go", "around"})
	if !set.RemoveAll([]string{"go", "around"}) {
		t.Errorf("Removing all existing strings for the first time should return true.")
	}
	if set.RemoveAll([]string{"go", "around"}) {
		t.Errorf("Removing them again should return false, but it returns true.")
	}
	if !set.IsEmpty() {
		t.Errorf("After removing all elements, %v should be empty, but it is not!", set)
	}
}

func TestStringSet_RetainAll(t *testing.T) {
	set := NewStringSet()
	set.AddAll([]string{"go", "around"})
	if !set.RetainAll([]string{"go", "sugar"}) {
		t.Errorf("RetainAll() changes %v, so it should return true, not false!", set)
	}
	if !set.Contains("go") {
		t.Errorf("%v should contain 'go', but it does not!", set)
	}
	if set.Contains("around") {
		t.Errorf("%v should have removed 'around', but it did not.", set)
	}
}

func TestStringSet_ToSlice(t *testing.T) {
	set := NewStringSet()
	set.AddAll([]string{"go", "around"})

	var stringSlice []string = set.ToSlice()
	if length := len(stringSlice); length < 2 {
		t.Errorf("Slice %v should have length >= 2, but it has length %v", stringSlice, length)
	}
}