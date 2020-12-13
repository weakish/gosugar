package gosugar

// Expect is inspired by Rust's expect
func Expect(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

// ExpectErr is inspired by Rust's expect_err
func ExpectErr(err error, msg string) {
	if err == nil {
		panic(msg)
	}
}
