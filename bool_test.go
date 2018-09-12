package gosugar

import "fmt"

func ExampleBtoi() {
	fmt.Println(Btoi(true))
	fmt.Println(Btoi(false))
	// Output:
	// 1
	// 0
}

func ExampleItob() {
	fmt.Println(Itob(0))
	fmt.Println(Itob(1))
	fmt.Println(Itob(16))
	fmt.Println(Itob(-1))
	// Output:
	// false
	// true
	// true
	// true
}
