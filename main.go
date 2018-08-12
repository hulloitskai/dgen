package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check for arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Did not receive any arguments! Exiting.")
		os.Exit(1)
	}
	args = os.Args[1:]

	// Parse repeater string.
	str := args[0]
	repeats := 5000

	// Retreive repeats from input, if applicable.
	if len(args) > 1 {
		repeatstr := args[1]
		var err error
		if repeats, err = strconv.Atoi(repeatstr); err != nil {
			fmt.Printf(
				"Could not parse repeat argument (expected an int, got %v).\n",
				repeatstr,
			)
			os.Exit(2)
		}
	}

	// Write to output.
	for i := 0; i < repeats; i++ {
		os.Stdout.WriteString(str)
	}
}
