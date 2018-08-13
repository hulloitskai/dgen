package throughput

import (
	"os"
	"testing"
)

const In = "throughput test string "

func BenchmarkDump(b *testing.B) {
	Dump(In, b.N, RecommendedBufSize, os.Stdout)
}
