package main

import (
	"art-decoder/functions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <encoded_string> [-ml] [--encode]")
		return
	}

	input := os.Args[1]
	isMultiLine := false
	isEncode := false

	// Check for additional flags
	for _, arg := range os.Args[2:] {
		switch arg {
		case "-ml", "--multi-line": // Support both -ml and --multi-line
			isMultiLine = true
		case "--encode":
			isEncode = true
		default:
			fmt.Println("Error: Unknown flag", arg)
			return
		}
	}

	if isMultiLine {
		// Read multi-line input from stdin
		fmt.Println("Enter your multi-line input (Ctrl+D to end):")
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.Join(lines, "\n")
	}

	if isEncode {
		// Encode mode
		encoded, err := functions.Encode(input)
		if err != nil {
			fmt.Println("Error")
			return
		}
		fmt.Println(encoded)
	} else {
		// Decode mode
		var decoded string
		var err error

		if isMultiLine {
			decoded, err = functions.DecodeMultiLine(input)
		} else {
			decoded, err = functions.Decode(input)
		}

		if err != nil {
			fmt.Println("Error")
			return
		}
		fmt.Println(decoded)
	}
}