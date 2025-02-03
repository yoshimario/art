package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"art-decoder/functions"
)

// runTests is your built-in test runner.
func runTests() {
	fmt.Println("Running tests...")

	tests := []struct {
		input    string
		expected string
	}{
		{"[5 #][5 -_]-[5 #]", "#####-_-_-_-_-_-#####"},
		{"[3 @][2 !]", "@@@!!"},
		{"[5 #][5 -_]-[5 #]]", "Error: Extra closing bracket found"},
		{"[5 #]5 -_]-[5 #]", "Error: Missing opening bracket"},
		{"[5#][5 -_]-[5 #]", "Error: Invalid format inside brackets (expected '[count char]')"},
		{"5 #[5 -_]-5 #", "Error: Missing opening bracket"},
	}

	for _, test := range tests {
		output, err := functions.DecodeMultiLine(test.input)
		if err != nil {
			output = err.Error()
		}

		if output == test.expected {
			fmt.Printf("✅ Test passed: %s\n", test.input)
		} else {
			fmt.Printf("❌ Test failed: %s\nExpected: %s\nGot: %s\n", test.input, test.expected, output)
		}
	}
}

func main() {
	// Define a flag for multi-line decoding.
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	flag.Parse()

	// If the first non-flag argument is "test", run the tests and exit.
	if len(flag.Args()) > 0 && flag.Arg(0) == "test" {
		runTests()
		return
	}

	// Otherwise, proceed in interactive mode.
	// Print the prompt only if standard input is a terminal.
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error checking stdin:", err)
		os.Exit(1)
	}
	// os.ModeCharDevice is set when the input is coming from a terminal.
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Enter multi-line input (Ctrl+D to end):")
	}

	// Read the entire input from standard input.
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Decode using the appropriate function.
	var output string
	if *multiLine {
		output, err = functions.DecodeMultiLine(string(input))
	} else {
		output, err = functions.Decode(string(input))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print(output)
}