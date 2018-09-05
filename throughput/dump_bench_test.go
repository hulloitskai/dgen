package throughput

import (
	"os"
	"testing"
)

func BenchmarkDump(b *testing.B) {
	const in = "throughput test string "
	Dump(in, b.N, RecommendedBufSize, os.Stdout)
}
