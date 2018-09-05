package main

import (
	"os"
	"strconv"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	// Simulate the program running with these arguments.
	os.Args = []string{"dgen", "benchmark test text ", strconv.Itoa(b.N)}
	main()
}
