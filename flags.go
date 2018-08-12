package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"os"
)

var (
	// Opts are flag-enabled options for dgen.
	Opts struct {
		// Stats will show statistics at the end of the string dump.
		Stats    bool `short:"s" long:"stats" description:"Show statistics after string dump."`
		Preserve bool `short:"p" long:"preserve" description:"Preserve whitespacing; do not add terminating newlines."`
	}

	fparser = flags.NewParser(&Opts, flags.Default)
)

func showHelp() {
	fparser.WriteHelp(os.Stdout)
}

func parseFlags() (args []string) {
	args, err := fparser.Parse()

	if err != nil {
		if flagerr, ok := err.(*flags.Error); ok {
			switch flagerr.Type {
			// Ignore minor parsing errors.
			case flags.ErrDuplicatedFlag, flags.ErrInvalidChoice:
				fmt.Println("Warning: Caught a discrepancy while parsing flags; " +
					"proceeding anyways...")
				return args

				// Friendly exit if help flag was triggered.
			case flags.ErrHelp:
				os.Exit(0)

			case flags.ErrUnknownFlag:
				fmt.Print("\n")
				showHelp()
				os.Exit(4)

			default:
				fmt.Fprintln(os.Stderr, "Encountered flag parsing error of type:",
					flagerr.Type)
				os.Exit(4)
			}
		}

		fmt.Fprintln(os.Stderr, "Failed to parse given flags (unknown error type):",
			err)
		os.Exit(5)
	}
	return args
}
