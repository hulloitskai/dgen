package cmd_test

import (
	"strconv"
	"testing"

	"github.com/stevenxie/dgen/cmd"
)

func BenchmarkExec(b *testing.B) {
	// Simulate the program running with these arguments.
	args := []string{"benchmark test text ", strconv.Itoa(b.N)}
	cmd.Exec(args)
}
