// cmd/variables/main.go
package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infers types of initialized variables.
	var d = true
	fmt.Println(d)

	// Uninitialized variables are zero-valued.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable.
	// e.g. for var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
