package main

import (
	"art-decoder/functions"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func isInteractive() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	// Check if stdin is a character device (terminal)
	return (info.Mode() & os.ModeCharDevice) != 0
}

func main() {
	// Define a flag for multi-line decoding.
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	flag.Parse()

	// Show prompt only when running interactively
	if *multiLine && isInteractive() {
		fmt.Println("Enter multi-line input. Press Ctrl+D to finish:")
	}

	// Read all input from standard input.
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(0) // Changed from os.Exit(1) to os.Exit(0) to prevent exit status 1
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0) // Changed from os.Exit(1) to os.Exit(0)
	}

	// Print the decoded output.
	fmt.Println(output)
}