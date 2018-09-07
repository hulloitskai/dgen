package main

import (
	"bytes"
	"github.com/steven-xie/glip"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMain_copyflag(t *testing.T) {
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
		t.Errorf("Expected clipboard to contain \"%s\", instead got \"%s\"",
			expect, out)
	}
}

func TestMain_preserveflag(t *testing.T) {
	const (
		in     = "a"
		repstr = "3"
		expect = "aaa"
	)

	// Configure dgen to output to buf.
	buf := new(bytes.Buffer)
	out = buf

	// Run dgen with the preserve flag.
	os.Args = []string{"dgen", "-p", in, repstr}
	main()

	// Check output.
	if res := buf.String(); res != expect {
		t.Errorf("Expected output to be \"%s\", instead got \"%s\"",
			expect, out)
	}
}

func TestMain_statsflag(t *testing.T) {
	const (
		in          = "testing "
		repstr      = "3"
		expectstr   = "testing testing testing "
		expectbytes = len(in) * 3
	)

	// Configure dgen to output to buf.
	buf := new(bytes.Buffer)
	out = buf

	// Run dgen with the preserve flag.
	os.Args = []string{"dgen", "-s", in, repstr}
	main()

	// Check output.
	res := buf.String()
	if !strings.Contains(res, expectstr) {
		t.Errorf("Expected output to contain \"%s\", instead got \"%s\"",
			expectstr, res)
	}

	expectedStatStr := strconv.Itoa(expectbytes) + " bytes"
	if !strings.Contains(res, expectedStatStr) {
		t.Errorf("Expected output to contain \"%s\", instead got \"%s\"",
			expectedStatStr, res)
	}
}
