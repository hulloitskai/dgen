package main

import (
	"bytes"
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"github.com/steven-xie/glip"
	"io"
	"log"
	"os"
	"time"
)

// out is the io.ReadWriter that dgen will write to.
var out io.ReadWriter = os.Stdout

func main() {
	// Read flags, initialize variables.
	var (
		args      = parseFlags()
		str, reps = parseArgs(args)
		bufsize   = throughput.RecommendedBufSize
		start     time.Time
		duration  time.Duration
	)

	// When using clipboard, output to a temporary buffer and disable
	// throughput.Dump's buffering strategy.
	if Opts.Copy {
		out = new(bytes.Buffer)
		bufsize = 0
	}

	// If the Stats opt is enabled, recording the start time.
	if Opts.Stats {
		start = time.Now()
	}

	n, err := throughput.Dump(str, reps, bufsize, out)

	// If the Copy opt is enabled, pipe data from the temporary buffer into the
	// clipboard.
	if Opts.Copy {
		board, err := glip.NewBoard()
		if err != nil {
			log.Fatalln("Failed to open system clipboard:", err)
		}
		board.ReadFrom(out)
	}

	// Record duration if the Stats opt is enabled.
	if Opts.Stats {
		duration = time.Since(start)
	}

	if !Opts.Copy { // if wrote to os.Stdout
		hasnl := hasTrailingNewline(str)

		// Ensure that if extra info is about to be produced, there are at least two
		// newlines before that info is printed.
		if err != nil || Opts.Stats { // extra info will be printed
			if hasnl {
				out.Write([]byte{'\n'})
			} else {
				io.WriteString(out, "\n\n")
			}
		} else if !Opts.Preserve && !hasnl {
			out.Write([]byte{'\n'})
		}
	}

	if err != nil {
		log.Fatalln("Encountered error while dumping source string:", err)
	}

	if Opts.Stats {
		fmt.Fprintf(out, "Successfully printed %d bytes in %s.\n", n,
			duration.String())
	}
}

func hasTrailingNewline(s string) bool {
	strlen := len(s)
	if strlen < 1 {
		return false
	}
	return s[strlen-1] == '\n'
}
