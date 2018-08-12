package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"os"
)

// Opts are flag-enabled options for dgen.
var Opts struct {
	// Stats will show statistics at the end of the string dump.
	Stats bool `short:"s" long:"stats" description:"Show statistics after string dump."`
}

func parseFlags() (args []string) {
	args, err := flags.Parse(&Opts)

	if err != nil {
		if flagerr, ok := err.(*flags.Error); ok {
			switch flagerr.Type {
			case flags.ErrDuplicatedFlag, flags.ErrInvalidChoice:
				fmt.Printf("Warning: Caught a discrepancy while parsing flags: %v. "+
					"Proceeding anyways...\n", err)
				return args

			// Friendly exit if help flag was triggered.
			case flags.ErrHelp:
				os.Exit(0)

			default:
				fmt.Fprintf(os.Stderr, "Failed to parse flags: %v (error type: %s)\n",
					err, flagerr.Type)
				os.Exit(4)
			}
		}

		fmt.Fprintln(os.Stderr, "Failed to parse given flags (unknown error type):",
			err)
		os.Exit(5)
	}
	return args
}
