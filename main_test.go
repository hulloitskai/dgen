package main

import (
	"bytes"
	"github.com/steven-xie/glip"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	const (
		in     = "a"
		expect = "aaaaaaaaaa"
	)

	// Configure dgen to output to a buffer.
	buf := new(bytes.Buffer)
	out = buf

	// Equivalent to: ./dgen "a"
	os.Args = []string{"dgen", in}
	main()

	// Check result.
	if res := buf.String(); res != expect {
		t.Fatalf("Expected glip to produce \"%s\", instead got \"%s\"", expect, out)
	}
}

func TestMain_customreps(t *testing.T) {
	const (
		in     = "testing text "
		repstr = "3"
		expect = "testing text testing text testing text "
	)

	// Set dgen to output to a buffer.
	buf := new(bytes.Buffer)
	out = buf

	// Simulate program running with these arguments.
	// Equivalent to: ./dgen "testing text " 3
	os.Args = []string{"dgen", in, repstr}
	main()

	// Check result.
	if res := buf.String(); res != expect {
		t.Fatalf("Expected glip to produce \"%s\", instead got \"%s\"", expect, out)
	}
}

func TestMain_copyopt(t *testing.T) {
	const (
		in     = "testing text "
		repstr = "2"
		expect = "testing text testing text "
	)

	os.Args = []string{"dgen", "-c", in, repstr}
	main()

	out, err := glip.ReadString()
	if err != nil {
		t.Fatal("Ran into error while reading clipboard:", err)
	}
	if out != expect {
		t.Errorf("Expected clipboard to contain \"%s\", instead got: \"%s\"",
			expect, out)
	}
}
