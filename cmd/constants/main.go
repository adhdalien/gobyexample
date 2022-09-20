// cmd/constants/main.go
package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean and numeric values.

// const declares a constant value.
const s string = "constant string"

func main() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 5_000_000

	// Constant expressions can perform arithmetic with arbitary precision.
	const d = 3e20 / n

	// A numeric constant has no type until it's given one, such as by an
	// explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call. For example, here
	// math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
