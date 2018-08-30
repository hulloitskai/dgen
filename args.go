package main

import (
	"log"
	"os"
	"strconv"
)

// parseArgs validates program arguments, and parses them into the string
// to be repeated (str), and repeat count (reps).
func parseArgs(args []string) (str string, reps int) {
	nargs := len(args)

	if nargs == 0 {
		// Quit if no arguments were received.
		log.Print("Warning: Did not receive any arguments!\n\n")
		showHelp()

		os.Exit(2)
	}

	if nargs > 2 {
		// Warn the user if more arguments were received than expected.
		log.Printf("Warning: Received more than 2 arguments. Ignoring the " +
			"following arguments:")

		for _, arg := range args[2:] {
			log.Printf("  – \"%s\"", arg)
		}
	}

	str = args[0]
	if nargs == 1 {
		return str, DefaultReps
	}

	repstr := args[1]
	var err error
	if reps, err = strconv.Atoi(repstr); err != nil {
		reps = parsePreset(repstr)
	}
	return str, reps
}
