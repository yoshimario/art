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
		fmt.Println("Usage: go run . [<encoded_string>] [-ml] [--encode]")
		return
	}

	var input string
	isMultiLine := false
	isEncode := false

	// Parse flags and input
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-ml", "--multi-line":
			isMultiLine = true
		case "--encode":
			isEncode = true
		default:
			// Treat the first non-flag argument as the input string
			if input == "" {
				input = arg
			} else {
				fmt.Println("Error: Unknown argument or too many inputs:", arg)
				return
			}
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