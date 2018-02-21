package gosugar

// Btoi converts boolean {true, false} to {1, 0}.
func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
