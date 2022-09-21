// cmd/spawnproc/main.go
package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes.

func main() {
	// We'll start with a simple command that takes no arguments or input and
	// just prints something to stdout. The exec.Command helper creates an
	// object to represent this external process.
	dateCommand := exec.Command("date")

	// The Output method runs the command, waits for it to finish and collects
	// its standard output. If there were no errors, dateOut will hold bytes
	// with the date info.
	dateOut, err := dateCommand.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Output and other methods of Command will return *exec.Error if there was
	// a problem executing the command (e.g. wrong path) and *exec.ExitOnError
	// if the command ran but exited with a non-zero return code.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed to execute", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Next we'll look at a slightly more involved case where we pipe data to
	// the external process on its stdin and collect the results from its
	// stdout.
	grepCommand := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start the process, write some
	// input to it, read the resulting output, and finally wait for the process
	// to exit.
	grepIn, _ := grepCommand.StdinPipe()
	grepOut, _ := grepCommand.StdoutPipe()
	grepCommand.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCommand.Wait()

	// We omitted error checks in the above example, but you could use the usual
	// if err != nil pattern for all of them. We also only collect the
	// StdoutPipe results, but you could collect the StderrPipe in exactly the
	// same way.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly
	// delineated command and argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full command with a string,
	// you can use bash's -c option:
	lsCommand := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCommand.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
