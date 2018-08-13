package main

import (
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"github.com/steven-xie/glip"
	"io"
	"os"
)

func main() {
	args := parseFlags()
	str, reps := parseArgs(args)

	// Set up output device.
	var out io.Writer = os.Stdout
	if Opts.Copy {
		board, err := glip.NewBoard()
		if err != nil {
			errln("Failed to open system clipboard:", err)
			os.Exit(1)
		}
		out = board
	}

	n, err := throughput.Dump(str, reps, Bufsize, out)

	// If wrote to os.Sdout...
	if !Opts.Copy {
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
	}

	if err != nil {
		errln("Encountered error while dumping source string:", err)
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
