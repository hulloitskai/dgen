package cmd_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stevenxie/dgen/cmd"
)

func BenchmarkExec(b *testing.B) {
	// Simulate the program running with these arguments.
	os.Args = []string{"dgen", "benchmark test text ", strconv.Itoa(b.N)}
	cmd.Exec()
}
