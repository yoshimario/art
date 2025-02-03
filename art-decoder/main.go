package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"art-decoder/functions"
)

func main() {
	// Define a flag for multi-line decoding.
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	flag.Parse()

	// Check whether standard input is coming from a terminal.
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error checking stdin:", err)
		os.Exit(1)
	}

	// Print the prompt only if input is from a terminal (i.e., interactive mode).
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Enter multi-line input (Ctrl+D to end):")
	}

	// Read all input from standard input.
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Decode the input using the appropriate function.
	var output string
	if *multiLine {
		output, err = functions.DecodeMultiLine(string(input)) // Use multi-line decoder
	} else {
		output, err = functions.Decode(string(input)) // Use single-line decoder
	}

	// Handle decoding errors.
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Print the decoded output.
	fmt.Println(output)
}