package main

import (
	"fmt"
	"os"
)

func errln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func errf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
