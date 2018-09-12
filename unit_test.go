package gosugar

import "testing"

func TestUnit(t *testing.T) {
	emptyStruct := Unit
	if emptyStruct != struct{}{} {
		t.Errorf("Unit is not an empty struct but %v", emptyStruct)
	}
}
