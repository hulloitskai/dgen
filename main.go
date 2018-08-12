package main

import (
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"os"
)

func main() {
	args := parseFlags()
	str, reps := parseArgs(args)

	n, err := throughput.Dump(str, reps, Bufsize, os.Stdout)

	hasnl := hasTrailingNewline(str)
	// Ensure that if extra info is about to be produced, there are at least two
	// newlines before that info is printed.
	if err != nil || Opts.Stats {
		if hasnl {
			fmt.Print("\n")
		} else {
			fmt.Print("\n\n")
		}
	} else if !Opts.Preserve && !hasnl {
		fmt.Print("\n")
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

func hasTrailingNewline(s string) bool {
	strlen := len(s)
	if strlen < 1 {
		return false
	}
	return s[strlen-1] == '\n'
}
