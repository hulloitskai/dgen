package cmd_test

import (
	"bytes"
	"testing"

	"github.com/steven-xie/glip"
	"github.com/stevenxie/dgen/cmd"
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
	args := []string{in}
	cmd.Exec(args)

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
	args := []string{in, repstr}
	cmd.Exec(args)

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

	args := []string{"-c", in, repstr}
	cmd.Exec(args)

	out, err := glip.ReadString()
	if err != nil {
		t.Fatal("Ran into error while reading clipboard:", err)
	}
	if out != expect {
		t.Errorf("Expected clipboard to contain \"%s\", instead got \"%s\"",
			expect, out)
	}
}
