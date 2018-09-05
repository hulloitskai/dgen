package throughput_test

import (
	"github.com/steven-xie/dgen/throughput"
	"os"
)

func Example() {
	const repstr = "test string "
	throughput.Dump(repstr, 3, throughput.RecommendedBufSize, os.Stdout)
	// Output: test string test string test string
}
