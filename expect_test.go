package gosugar

import (
	"errors"
	"fmt"
	"testing"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by 0")
	} else {
		return x / y, nil
	}
}

func ExampleExpect() {
	result, err := divide(6, 3)
	Expect(err, "cannot divide by 0")
	fmt.Println(result)
	// Output: 2
}

func TestExpect(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("6 / 0 should panic")
		}
	}()

	_, err := divide(6, 0)
	Expect(err, "cannot divide by 0")
}

func ExampleExpectErr() {
	_, err := divide(6, 0)
	ExpectErr(err, "assuming there is an error")
}
