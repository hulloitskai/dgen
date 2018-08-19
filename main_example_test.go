package main

import (
	"os"
)

func Example() {
	const in = "testing text "
	const repstr = "3"

	// Simulate program running with these arguments.
	// Equivalent to: ./dgen "testing text " 3
	os.Args = []string{"dgen", in, repstr}

	main()
	// Output: testing text testing text testing text
}

func Example_default() {
	const in = "a"

	// Equivalent to: ./dgen "a"
	os.Args = []string{"dgen", in}

	main()
	// Output: aaaaaaaaaa
}
