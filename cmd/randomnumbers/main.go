// cmd/randomnumbers/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go's math/rand package provides pseudorandom number generation.

func main() {
	p := fmt.Print
	pl := fmt.Println

	// For example, rand.Intn returns a random int n, 0 <= n < 100
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pl()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0
	pl(rand.Float64())

	// This can be used to generate random floats in other ranges, for example
	// 5.0 <= f' < 10.0
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()

	// The default number generator is deterministic, so it'll produce the same
	// sequence of numbers each time by default. To produce varying sequences,
	// give it a seed that changes. Note that this is not safe to use for
	// random numbers you intend to be secret, use crypto/rand for those.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call the resulting rand.Rand just like the functions on the rand package.
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	pl()

	// If you seed a source with the same number, it produces the same sequence
	// of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	pl()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
	pl()

	// See the math/rand package docs for references on other random quantities
	// that Go can provide: https://pkg.go.dev/math/rand
}
