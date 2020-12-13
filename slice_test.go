package gosugar

import (
	"fmt"
	"testing"
)



func ExampleGetByte() {
	fmt.Println(GetByte("goaround", 2))
	// Output: 97
}

func ExampleGetByte_exceeding() {
	fmt.Println(GetByte("goaround", 100))
	// Output: 100
}

func ExampleGetByte_negative() {
fmt.Println(GetByte("goaround", -1))
	// Output: 100
}

func ExampleGetByte_leftExceeding() {
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
func ExampleGetBool_exceeding() {
	fmt.Println(GetBool([]bool{true, false}, 100))
	// Output: false
}
func ExampleGetBool_negative() {
	fmt.Println(GetBool([]bool{true, false}, -1))
	// Output: false
}
func ExampleGetBool_left_exceeding() {
	fmt.Println(GetBool([]bool{true, false}, -100)) // left exceeding
	// Output: true
}

func ExampleGetInt() {
	fmt.Println(GetInt([]int{1, 2, 3}, 0))
	// Output: 1
}
func ExampleGetInt_exceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, 100))
	// Output: 3
}
func ExampleGetInt_negative() {
	fmt.Println(GetInt([]int{1, 2, 3}, -1))
	// Output: 3
}
func ExampleGetInt_leftExceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, -100))
	// Output: 1
}

func ExampleGetString() {
	fmt.Println(GetString([]string{"go", "around"}, 0))
}
func ExampleGetString_exceeding() {
	fmt.Println(GetString([]string{"go", "exceeded"}, 100))
	// Output: exceeded
}
func ExampleGetString_negative() {
	fmt.Println(GetString([]string{"go", "negative"}, -1))
	// Output: negative
}
func ExampleGetString_leftExceeding() {
	fmt.Println(GetString([]string{"go", "left exceeding"}, -100))
	// Output: go
}

func ExampleSubString() {
	fmt.Println(SubString("goaround", 0, 2))
	// Output: go
	}

	func ExampleSubString_negative() {
		fmt.Println(SubString("negative", 2, -1))
		// Output: gativ
	}

	func ExampleSubString_exceeding() {
		fmt.Println(SubString("exceeding", 2, 100))
		// Output: ceeding
	}

	func ExampleSubString_equalIndices() {
		fmt.Println(SubString("equal indices", 2, 2))
		// Output:
	}

	func ExampleSubString_empty() {
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

func ExampleSubBoolSlice_empty() {
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
