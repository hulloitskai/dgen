// Package throughput provides functionality for producing content at an
// extremely high rate.
package throughput

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	ess "github.com/unixpickle/essentials"
)

// RecommendedBufSize is the buffer size recommended for Dump's bufsize
// argument.
//
// It is roughly based on the speed of writing to os.Stdout. A different size
// should be considered for programs that do not write directly to the display.
const RecommendedBufSize = 1000

// Dump repeats s, and streams the result to out.
//
// Dump's bufsize argument indicates the maximum string size (in bytes) for
// which a buffering strategy should be used to reduce the frequency of writes
// to out.
//
// To disable buffering entirely, set bufsize to 0.
func Dump(s string, repeats, bufsize int, out io.Writer) (n int, err error) {
	// Validate repeats.
	if repeats < 0 {
		return n, fmt.Errorf("throughput: cannot repeat a string a negative (%d) "+
			"number of times", repeats)
	}

	var (
		strlen = len(s)
		// usebuf describes whether or not to use a buffering strategy.
		usebuf = strlen < bufsize

		// ntmp is a temporary byte counter.
		ntmp   int
		n64tmp int64
		buf    *bytes.Buffer
	)

	// Initialize buf if necessary.
	if usebuf {
		buf = bytes.NewBuffer(make([]byte, 0, bufsize))
	}

	for i := 1; i <= repeats; i++ {
		if !usebuf {
			ntmp, err = io.WriteString(out, s)
			// This counts for bytes since output is actually written.
			n += ntmp
			if err != nil {
				return n, ess.AddCtx("throughput: writing string to output",
					err)
			}
			// We are done, since we do not have to deal with the buffer.
			continue
		}

		// This does not count for bytes since we are internally recording the
		// result.
		if _, err = buf.WriteString(s); err != nil {
			return n, ess.AddCtx("throughput: writing string to internal buffer", err)
		}

		// If adding one more string to the buffer will make it larger than the
		// buffer's capacity, pipe the buffer to output, and reset the buffer.
		if buf.Len()+strlen > bufsize || i == repeats {
			n64tmp, err = buf.WriteTo(out)
			n += int(n64tmp)
			if err != nil {
				return n, ess.AddCtx("throughput: writing buffered string to output",
					err)
			}

			// Return if at the final iteration.
			if i == repeats {
				return n, nil
			}

			// Reset the buffer if we are not at the final iteration.
			buf.Reset()
		}
	}

	if usebuf {
		// When we have a buffer, we always want to terminate within the for
		// loop, so that we can empty the remaining contents of the buffer.
		return n, errors.New("throughput: failed to properly terminate " +
			"string-repeat loop")
	}
	return n, nil
}
