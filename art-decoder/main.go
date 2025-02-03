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