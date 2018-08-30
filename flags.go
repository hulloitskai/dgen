package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"
)

var (
	// Opts are flag-enabled options for dgen.
	Opts struct {
		Stats    bool `short:"s" long:"stats" description:"Show statistics after string dump."`
		Preserve bool `short:"p" long:"preserve" description:"Preserve whitespacing; do not add terminating newlines."`
		Copy     bool `short:"c" long:"copy" description:"Write dump to clipboard, rather than to standard output."`
	}

	fparser = makeParser()
)

func makeParser() (p *flags.Parser) {
	p = flags.NewParser(&Opts, flags.Default)
	p.Usage = "[OPTIONS] <string> [<repeat count> | <preset name>]"
	return p
}

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
				log.Println("Warning: Caught a discrepancy while parsing flags; " +
					"proceeding anyways...")
				return args

				// Friendly exit if help flag was triggered.
			case flags.ErrHelp:
				os.Exit(0)

			case flags.ErrUnknownFlag:
				fmt.Print("\n")
				showHelp()
				os.Exit(1)

			default:
				log.Fatalln("Encountered flag parsing error of type:", flagerr.Type)
			}
		}

		log.Fatalln("Failed to parse given flags (unknown error type):", err)
	}
	return args
}
