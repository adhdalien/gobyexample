// cmd/commandline/main.go
package main

import (
	"flag"
	"fmt"
	"os"
)

// Command-line arguments are a common way to parameterize execution of
// programs. For example, go run hello.go uses run and hello.go arguments to the
// go program.

func main() {
	// Some command-line tools, like the go tool or git have many subcommands,
	// each with its own set of flags. For example, go build and go get are two
	// different subcommands of the go tool. The flag package lets us easily
	// define simple subcommands that have their own flags.

	// We declare a subcommand using the NewFlagSet function, and proceed to
	// define new flags specific for this subcommand.
	fooCommand := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCommand.Bool("enable", false, "enable")
	fooName := fooCommand.String("name", "", "name")

	// For a different subcommand we can define different supported flags.
	barCommand := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCommand.Int("level", 0, "level")

	// The subcommand is expected as the first argument to the program.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	switch os.Args[1] {
	case "foo":
		fooCommand.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCommand.Args())
	case "bar":
		barCommand.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCommand.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

func cliArguments() {
	// os.Args provides access to raw command-line arguments. Note that the
	// first value in this slice is the path to the program, and the os.Args[1:]
	// holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[1]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// Command-line flags are a common way to specify options for command-line
// programs. For example, in wc -l the -l is a command-line flag.

// Go provides a flag package supporting basic command-line flag parsing
// We'll use this package to implement our example command-line program.
func cliFlags() {
	// Basic flag declarations are available for string, integer and boolean
	// options. Here we declare a string flag word with a default value "foo"
	// and a short description. This flag.String function returns a string
	// pointer (not a string value); we'll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// This declares numb and fork flags, using a similar approach.
	numbPtr := flag.Int("numb", 42, "an integer")
	forkPtr := flag.Bool("fork", false, "a boolean")

	// It's also possible to declare an option that uses an existing var
	// declared elsewhere in the program. Note that we need to pass in a pointer
	// to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-
	// line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options any any trailing positional
	// arguments. Note that we need to dereference the pointers with e.g.
	// *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
