// cmd/files/main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading and writing files are basic tasks needed for many Go programs. First
// we'll look at some examples of reading files, then we'll look at writing.

// Most file operations require checking calls for errors. This helper will
// streamline our error checks in the main() function.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Perhaps the most basic file reading task is slurping a file's entire
	// contents into memory.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// You'll often want more control over how and what parts of a file are
	// read. For these tasks, start by Opening a file to obtain an os.File val.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file. Allow up to 5 to be read
	// but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// The io package provides some functions that may be helpful for file
	// reading. For example, reads like the ones above can be more robustly
	// implemented with ReadAtLeast.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but Seek(0, 0) accomplishes this.
	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package implements a buffered reader that may be useful for
	// both its efficiency with many small reads and because of the additional
	// reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you're done (usually this would be scheduled
	// immediately after Opening with defer).
	f.Close()

	// Writing files in Go follows similar patterns to the ones we see above.
	// To start, here's how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err = os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err = os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	// You can Write byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err = f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err = f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we
	// saw above.
	w := bufio.NewWriter(f)
	n4, err4 := w.WriteString("buffered\n")
	check(err4)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the
	// underlying writer.
	w.Flush()
}
