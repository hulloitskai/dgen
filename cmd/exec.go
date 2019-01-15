package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/steven-xie/glip"
	"github.com/stevenxie/dgen/pkg/throughput"
)

// Exec is the program entrypoint.
func Exec(args []string) {
	// Parse options.
	kingpin.MustParse(app.Parse(args))

	// Initialize variables.
	var (
		bufsize int
		start   time.Time
	)

	// When using clipboard, disable throughput.Dump's buffering strategy.
	if !Opts.Copy {
		bufsize = throughput.RecommendedBufSize
	}
	// If the Stats opt is enabled, recording the start time.
	if Opts.Stats {
		start = time.Now()
	}

	n, err := throughput.Dump(Opts.Msg, Opts.Reps, bufsize, Opts.Out)
	if Opts.Out == os.Stdout {
		hasnl := hasTrailingNewline(Opts.Msg)

		// Ensure that if extra info is about to be produced, there are at least two
		// newlines before that info is printed.
		if err != nil || Opts.Stats { // extra info will be printed
			if hasnl {
				fmt.Println()
				io.WriteString(os.Stdout, "\n")
			} else {
				fmt.Print("\n\n")
			}
		} else if !Opts.Preserve && !hasnl {
			fmt.Println()
		}
	}
	if err != nil {
		kingpin.Errorf("Error while dumping message: %v", err)
	}

	// If the Copy opt is enabled, pipe data from the temporary buffer into the
	// clipboard.
	if Opts.Copy {
		board, err := glip.NewBoard()
		if err != nil {
			kingpin.Errorf("Error while opening clipboard: %v", err)
		}
		if _, err = board.ReadFrom(Opts.Out); err != nil {
			kingpin.Errorf("Error while reading from buffer to clipboard: %v", err)
		}
	}

	if Opts.Stats {
		var (
			duration = time.Since(start)
			verb     = "printed"
		)
		if Opts.Copy {
			verb = "copied"
		}
		fmt.Printf("Successfully %s %d bytes in %s.\n", verb, n, duration.String())
	}
}

func hasTrailingNewline(s string) bool {
	strlen := len(s)
	if strlen < 1 {
		return false
	}
	return s[strlen-1] == '\n'
}
