// +build darwin

package main

import (
	"github.com/steven-xie/glip"
	"os"
	"testing"
)

func TestMain_copymode(t *testing.T) {
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
