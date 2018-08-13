package main

import (
	"fmt"
	"github.com/steven-xie/dgen/throughput"
	"os"
)

/////////////////////////////
// Configurable constants
/////////////////////////////

const (
	// DefaultReps is the default number of repeats, used if not otherwise
	// specified.
	DefaultReps = 10
	// Bufsize is the buffer size passed to throughput.Dump to configure its
	// buffering strategy.
	Bufsize = throughput.RecommendedBufSize
)

var (
	// Presets is a map of string identifiers to repeat counts. Identifiers
	// correspond to various messaging services.
	Presets = map[string]uint{
		"fb": 5000, "rpost": 40000, "rcomment": 10000, "rmsg": 10000,
		"twitter": 280,
	}
)

func parsePreset(id string) int {
	p, ok := Presets[id]
	if !ok {
		fmt.Fprintf(os.Stderr, "Could not find preset: %s\n", id)
		os.Exit(6)
	}
	return int(p)
}
