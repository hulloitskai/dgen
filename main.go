package main

import (
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"os"
	"strings"
)

func main() {
	args := parseFlags()
	str, reps := parseArgs(args)

	n, err := throughput.Dump(str, reps, Bufsize, os.Stdout)

	// Ensure that if extra info is about to be produced, there are at least two
	// newlines before that info is printed.
	if err != nil || Opts.Stats {
		nlcount := strings.Count(str, "\n")
		switch nlcount {
		case 0:
			fmt.Print("\n\n")
		case 1:
			fmt.Print("\n")
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Encountered error while dumping source string:",
			err)
		os.Exit(7)
	}
	if Opts.Stats {
		fmt.Printf("Successfully printed %d bytes.\n", n)
	}
}
