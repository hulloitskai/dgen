package main

import (
	"log"
)

func init() {
	// Only write the log message to os.Stderr.
	log.SetFlags(0)
}

const (
	// defaultReps is the default number of repeats, used if not otherwise
	// specified.
	defaultReps = 10
)

var (
	// presets is a map of string identifiers to repeat counts. Identifiers
	// correspond to various messaging services.
	presets = map[string]uint{
		"fb":       5000,
		"rpost":    40000,
		"rcomment": 10000,
		"rmsg":     10000,
		"twitter":  280,
	}
)

func parsePreset(id string) int {
	p, ok := presets[id]
	if !ok {
		log.Fatalf("Could not find preset: %s", id)
	}
	return int(p)
}
