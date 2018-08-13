package throughput_test

import (
	"github.com/steven-xie/dgen/throughput"
	"os"
)

func ExampleDump() {
	const exampleIn = "test string "
	throughput.Dump(exampleIn, 3, throughput.RecommendedBufSize, os.Stdout)
	// Output: test string test string test string
}
