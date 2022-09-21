// cmd/testandbench/main.go
package main

import (
	"fmt"
	"testing"
)

// Unit testing is an important part of writing principled Go programs. The
// testing package provides the tools we need to write unit tests and the go
// test command runs tests.

// A test is created by writing a function with a name beginning with Test.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue executing the test.
		// t.Fatal* will report test failures and stop the test immediately.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Writing tests can be repetitive, so it's idiomatic to use a table-driven
// style, where test inputs and expected outputs are listed in a table and a
// single loop walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each table entry. These are
		// shown separately when executing go test -v.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in _test.go files and are named beginning with
// Benchmark. The testing runner executes each benchmark function several times,
// increasing b.N on each run until it collects a precise measurement.
func BenchmarkIntMin(b *testing.B) {
	// Typically the benchmark runs a function we're benchmarking in a loop b.N
	// times.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
