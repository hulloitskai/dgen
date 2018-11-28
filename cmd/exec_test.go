package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stevenxie/dgen/cmd"

	"github.com/steven-xie/glip"
)

func TestExec(t *testing.T) {
	const (
		in     = "a"
		expect = "aaaaaaaaaa"
	)

	// Configure dgen to output to a buffer.
	buf := new(bytes.Buffer)
	cmd.Opts.Out = buf

	// Equivalent to: ./dgen "a"
	os.Args = []string{"dgen", in}
	cmd.Exec()

	// Check result.
	if res := buf.String(); res != expect {
		t.Fatalf("Expected output to be \"%s\", instead got \"%s\"", expect, res)
	}
}

func TestExec_customreps(t *testing.T) {
	const (
		in     = "testing text "
		repstr = "3"
		expect = "testing text testing text testing text "
	)

	// Set dgen to output to a buffer.
	buf := new(bytes.Buffer)
	cmd.Opts.Out = buf

	// Simulate program running with these arguments.
	// Equivalent to: ./dgen "testing text " 3
	os.Args = []string{"dgen", in, repstr}
	cmd.Exec()

	// Check result.
	if res := buf.String(); res != expect {
		t.Fatalf("Expected output to be \"%s\", instead got \"%s\"", expect, res)
	}
}

func TestExec_copyflag(t *testing.T) {
	const (
		in     = "testing text "
		repstr = "2"
		expect = "testing text testing text "
	)

	os.Args = []string{"dgen", "-c", in, repstr}
	cmd.Exec()

	out, err := glip.ReadString()
	if err != nil {
		t.Fatal("Ran into error while reading clipboard:", err)
	}
	if out != expect {
		t.Errorf("Expected clipboard to contain \"%s\", instead got \"%s\"",
			expect, out)
	}
}
