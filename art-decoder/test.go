package main

import (
	"art-decoder/functions"
	"fmt"
)

func runTests() {
	fmt.Println("Running tests...")

	tests := []struct {
		input    string
		expected string
	}{
		{"[5 #][5 -_]-[5 #]", "#####-_-_-_-_-#####"},
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