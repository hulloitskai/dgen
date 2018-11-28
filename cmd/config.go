package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	// Version is the program version. To be generated upon compilation by:
	//   -ldflags "-X github.com/stevenxie/dgen/cmd.Version=$(VERSION)"
	//
	// It should match the output of the following command:
	//   git describe --tags | cut -c 2-
	Version string

	app = kingpin.New(
		"dgen",
		"A CLI tool for repeating a string an excessive number of times.",
	).Version(Version)
)

func init() {
	// Customize help, version flag.
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	// Args:
	app.Arg("message", "The message to repeat.").Required().StringVar(&Opts.Msg)
	parseReps(app.Arg(
		"count | preset",
		"The number of times to repeat the message.",
	))

	// Flags:
	app.Flag("stats", "Show statistics after string dump.").Short('s').
		BoolVar(&Opts.Stats)
	app.Flag(
		"preserve",
		"Preserve whitespacing; do not add terminating newlines.",
	).Short('p').BoolVar(&Opts.Preserve)
	app.Flag(
		"copy",
		"Write to clipboard, rather than stdout. Implies --preserve.",
	).Short('c').Action(func(*kingpin.ParseContext) error {
		// When using clipboard, output to a temporary buffer.
		Opts.Out = new(bytes.Buffer)
		return nil
	}).BoolVar(&Opts.Copy)
}

// Opts are dgen's configurable options.
var Opts = struct {
	Stats    bool
	Preserve bool
	Copy     bool

	Msg  string
	Reps int

	Out io.ReadWriter
}{
	Out:  os.Stdout,
	Reps: 10,
}
