package main

import (
	"os"
	"strconv"
	"testing"
)

func Example() {
	// Simulate program running with these arguments.
	os.Args = []string{"dgen", "testing text ", "3"}

	main() // Equivalent of ./dgen "testing text " 3
	// Output: testing text testing text testing text
}

func BenchmarkMain(b *testing.B) {
	// Simulate the program running with these arguments.
	os.Args = []string{"dgen", "benchmark test text ", strconv.Itoa(b.N)}
	main()
}
