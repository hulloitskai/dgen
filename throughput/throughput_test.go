package throughput

import (
	"os"
	"testing"
)

////////////////////////////////////
// Standardized testing constants
////////////////////////////////////

const Bufsize = RecommendedBufSize

var In = "throughput test string "

////////////////////////////////////
// Tests and benchmarks
////////////////////////////////////

func ExampleDump() {
	const exampleIn = "test string "
	Dump(exampleIn, 3, Bufsize, os.Stdout)
	// Output: test string test string test string
}

func BenchmarkDump(b *testing.B) {
	Dump(In, b.N, Bufsize, os.Stdout)
}
