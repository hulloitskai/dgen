package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/stevenxie/dgen/internal/info"
)

var app = kingpin.New(
	"dgen",
	"A CLI tool for repeating a string an excessive number of times.",
).Version(info.Version)

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
