// cmd/pointers/main.go
package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records
// within your program.

// We'll show how pointers work in contrast to values with 2 functions: zeroval
// and zeroptr.

// zeroval has an int parameter, so arguments will be passed to it
// by value. zeroval will get a copy of ival distinct from the one in the
// calling function.
func zeroval(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int
// pointer. The *iptr code in the function body then deferences the pointer from
// its memory address to the current value at that address. Assigning a value to
// a deferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// zeroval doesn't affect the original i variable.
	zeroval(i)
	fmt.Println("zeroval:", i)

	// The & syntax gives the memory address of i, i.e. a pointer to i. This
	// means it will change the original i variable.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too
	fmt.Println("pointer:", &i)
}
