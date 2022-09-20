// cmd/range/main.go
package main

import "fmt"

// Range iterates over elements in a variety of data structures. Let's see how
// to use range with some of the data structures we've already learned.

func main() {
	// Here we use range to sum the numbers in a slice. Arrays work like this
	// too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and the value for
	// each entry. Above we didn't need the index so we ignored it with _.
	// Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("idx:", i)
		}
	}

	// range on maps iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. The first value is
	// the starting byte index of the rune and the second is the rune itself.
	// See Strings and Runes for more details:
	// https://gobyexample.com/strings-and-runes
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
