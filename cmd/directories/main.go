// cmd/directories/main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Go has several functions for working with directories in the file system.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// When creating temporary directories, it's good practice to defer their
	// removal. os.Removeall will delete a whole directory tree (like rm -rf).
	defer os.RemoveAll("subdir")

	// Helper function to create empty files.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")

	// We can create a hierarchy of directories, including parents with MkDirAll
	// . This is similar to the command-line mkdir -p.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objs.
	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir lets us change the current working directory, like cd.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Now we'll see the contents of subdir/parent/child when listing the
	// current directory.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started
	err = os.Chdir("../../..")
	check(err)

	// We can also visit a directory recursively, including all its sub-dirs.
	// Walk accepts a callback function to handle every file or directory
	// visited.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	check(err)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
