package cmd

import (
	"fmt"
	"strconv"

	"github.com/alecthomas/kingpin"
	ess "github.com/unixpickle/essentials"
)

var (
	// presets is a map of string identifiers to repeat counts. Identifiers
	// correspond to various messaging services.
	presets = map[string]int{
		"fb":       5000,
		"rpost":    40000,
		"rcomment": 10000,
		"rmsg":     10000,
		"twitter":  280,
	}
)

func parsePreset(id string) (int, error) {
	p, ok := presets[id]
	if !ok {
		return 0, fmt.Errorf("could not find preset '%s'", id)
	}
	return p, nil
}

// repsParser parses a value into Opts.Reps.
type repsParser struct{}

func (rp repsParser) Set(value string) error {
	var (
		fchar = value[0]
		err   error
	)

	if ('0' <= fchar) && (fchar <= '9') { // token is an int
		if Opts.Reps, err = strconv.Atoi(value); err != nil {
			return ess.AddCtx("failed to parse 'count' argument as an int", err)
		}
		return nil
	}

	// Token is a preset
	Opts.Reps, err = parsePreset(value)
	return err
}

func (rp repsParser) String() string {
	return strconv.Itoa(Opts.Reps)
}

func parseReps(s kingpin.Settings) {
	var target repsParser
	s.SetValue(target)
}
