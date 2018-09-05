package main

import (
	"bytes"
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"github.com/steven-xie/glip"
	"io"
	"log"
	"os"
)

// out is the io.ReadWriter that dgen will write to.
var out io.ReadWriter = os.Stdout

func main() {
	// Set up flags and buffer size.
	var (
		args      = parseFlags()
		str, reps = parseArgs(args)
		bufsize   = throughput.RecommendedBufSize
	)

	if Opts.Copy {
		out = new(bytes.Buffer)
		bufsize = 0 // disallow buffering when using clipboard
	}

	n, err := throughput.Dump(str, reps, bufsize, out)

	if Opts.Copy {
		board, err := glip.NewBoard()
		if err != nil {
			log.Fatalln("Failed to open system clipboard:", err)
		}
		board.ReadFrom(out)
	}

	if !Opts.Copy { // if wrote to os.Stdout
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
		log.Fatalln("Encountered error while dumping source string:", err)
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
