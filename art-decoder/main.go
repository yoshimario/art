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
			if input == "" {
				input = arg
			} else {
				fmt.Println("Error: Too many inputs or unknown argument:", arg)
				return
			}
		}
	}

	if isMultiLine {
		fmt.Println("Enter multi-line input (Ctrl+D to end):")
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
		encoded, err := functions.Encode(input)
		if err != nil {
			fmt.Println("Encoding Error:", err)
			return
		}
		fmt.Println(encoded)
	} else {
		decoded, err := functions.DecodeMultiLine(input)
		if err != nil {
			fmt.Println("Decoding Error:", err)
			return
		}
		fmt.Println(decoded)
	}
}
