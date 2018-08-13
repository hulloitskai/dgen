// Package throughput provides functionality for producing content at an
// extremely high rate.
package throughput

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// RecommendedBufSize is an arbitrary size recommended for Dump's "bufsize"
// argument.
const RecommendedBufSize = 1000

// Dump repeats the string "in", and streams the output to "out". Takes in a
// "bufsize" arguments that indicates the maximum string size (in bytes) for
// which a buffering strategy should be used to reduce the frequency of writes
// to "out".
func Dump(in string, repeats, bufsize int, out io.Writer) (n int, err error) {
	// Validate "repeats" argument.
	if repeats < 0 {
		return n, fmt.Errorf("cannot repeat a string a negative (%d) number of"+
			"times", repeats)
	}

	var (
		strlen = len(in)
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
		if usebuf {
			// This does not count for bytes since we are internally recording the
			// result.
			if _, err = buf.WriteString(in); err != nil {
				return n, fmt.Errorf("failed to write string to internal buffer: %v",
					err)
			}
		} else {
			ntmp, err = io.WriteString(out, in)
			// This counts for bytes since output is actually written.
			n += ntmp
			if err != nil {
				return n, fmt.Errorf("failed to write string to output: %v", err)
			}
			// We are done, since we do not have to deal with the buffer.
			continue
		}

		// If adding one more string to the buffer will make it larger than the
		// buffer's capacity, pipe the buffer to output, and reset the buffer.
		//
		// If this is the final iteration, just pipe the buffer to output.
		if buf.Len()+strlen > bufsize || i == repeats {
			n64tmp, err = buf.WriteTo(out)
			n += int(n64tmp)
			if err != nil {
				return n, fmt.Errorf("failed to write buffered string to output: %v",
					err)
			}

			// Return if at the final iteration.
			if i == repeats {
				return n, nil
			}

			// Reset the buffer if we are not at the final iteration.
			buf = bytes.NewBuffer(make([]byte, 0, bufsize))
		}
	}

	return n, errors.New("Dump unexpectedly reached the end of the function; " +
		"expected to terminate within the string-repeat loop")
}
