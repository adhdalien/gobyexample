// cmd/sorting/main.go
package main

import (
	"fmt"
	"sort"
)

// Go's sort package implements sorting for builtins and user-defined types.
// We'll look at sorting for builtins first.

// Sometimes we'll want to sort a collection by something other than its natural
// order. For example, suppose we wanted to sort strings by their length instead
// of alphabetically. Here's an example of custom sorts in Go.
type byLength []string

// We'll implement sort.Interface - Len, Less and Swap on our type so we can use
// the sort package's generic Sort function. Len and Swap will usually be across
// types and Less will hold the actual custom sorting logic. In our case we want
// to sort in order of increasing string length, so we use len(s[i]) and
// len(s[j]) here.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// Sort methods are specific to the builtin type; here's an example for
	// strings. Note that sorting is in-place, so it changes the given slice and
	// doesn't return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	// An example of sorting ints.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	// We can also use sort to check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	// With the custom byLength type in place, we can now implement our custom
	// sort of converting this fruits slice byLength, and then use sort.Sort on
	// that typed slice.
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// Running our program shows all our lists sorted as desired.
}
