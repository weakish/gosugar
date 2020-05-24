package gosugar

import (
	"fmt"
	"testing"
)



func ExampleGetByte() {
	fmt.Println(GetByte("goaround", 2))
	// Output: 97
}
func ExampleGetByteExceeding() {
	fmt.Println(GetByte("goaround", 100))
	// Output: 100
}
func ExampleGetByteNegative() {
	fmt.Println(GetByte("goaround", -1))
	// Output: 100
}
func ExampleGetByteLeftExceeding() {
	fmt.Println(GetByte("goaround", -100))
	// Output: 103
}
func TestGetByteFromEmptyStringShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Invoking GetByte on an empty string should panic!")
		}
	}()

	GetByte("", 0)
}

func ExampleGetBool() {
	fmt.Println(GetBool([]bool{true, false}, 0))
	// Output: true
}
func ExampleGetBoolExceeding() {
	fmt.Println(GetBool([]bool{true, false}, 100))
	// Output: false
}
func ExampleGetBoolNegative() {
	fmt.Println(GetBool([]bool{true, false}, -1))
	// Output: false
}
func ExampleGetBoolLeftExceeding() {
	fmt.Println(GetBool([]bool{true, false}, -100))
	// Output: true
}

func ExampleGetInt() {
	fmt.Println(GetInt([]int{1, 2, 3}, 0))
	// Output: 1
}
func ExampleGetIntExceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, 100))
	// Output: 3
}
func ExampleGetIntNegative() {
	fmt.Println(GetInt([]int{1, 2, 3}, -1))
	// Output: 3
}
func ExampleGetIntLeftExceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, -100))
	// Output: 1
}

func ExampleGetString() {
	fmt.Println(GetString([]string{"go", "around"}, 0))
	// Output: go
}
func ExampleGetStringExceeding() {
	fmt.Println(GetString([]string{"go", "around"}, 100))
	// Output: around
}
func ExampleGetStringNegative() {
	fmt.Println(GetString([]string{"go", "around"}, -1))
	// Output: around
}
func ExampleGetStringLeftExceeding() {
	fmt.Println(GetString([]string{"go", "around"}, -100))
	// Output: go
}

func ExampleSubString() {
	fmt.Println(SubString("goaround", 0, 2))
	// Output: go
}
func ExampleSubStringNegative() {
	fmt.Println(SubString("goaround", 2, -1))
	// Output: aroun
}
func ExampleSubStringExceeding() {
	fmt.Println(SubString("goaround", 2, 100))
	// Output: around
}
func ExampleSubStringEqualIndices() {
	fmt.Println(SubString("goaround", 2, 2))
	// Output:
}
func ExampleSubEmptyString() {
	fmt.Println(SubString("", 0, 0))
	// Output:
}
func TestSubStringUpsideDownShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Subslicing a string with fromIndex > toIndex should panic!")
		}
	}()

	SubString("goaround", 2, 0)
}

func ExampleSubBoolSlice() {
	fmt.Println(SubBoolSlice([]bool{true, false}, 0, 1))
	// Output: [true]
}
func ExampleSubEmptyBoolSlice() {
	fmt.Println(SubBoolSlice([]bool{}, 0, 0))
	// Output: []
}

func ExampleSubStringSlice() {
	fmt.Println(SubStringSlice([]string{"go", "around"}, 1, 2))
	// Output: [around]
}

func ExampleSubIntSlice() {
	fmt.Println(SubIntSlice([]int{1, 2, 3}, 0, 2))
	// Output: [1 2]
}
