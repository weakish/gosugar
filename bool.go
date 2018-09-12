package gosugar

// Btoi converts boolean {true, false} to {1, 0}.
func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// Itob converts integers to boolean, 0 is false, non-zeros are true.
func Itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}
