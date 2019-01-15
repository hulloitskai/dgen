package throughput_test

import (
	"os"

	"github.com/stevenxie/dgen/pkg/throughput"
)

func Example() {
	const repstr = "test string "
	throughput.Dump(repstr, 3, throughput.RecommendedBufSize, os.Stdout)
	// Output: test string test string test string
}
