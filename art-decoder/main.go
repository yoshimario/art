package main

import (
	"art-decoderg/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <encoded_string> [--multi-line] [--encode]")
		return
	}

	input := os.Args[1]
	isMultiLine := false
	isEncode := false

	// Check for additional flags
	for _, arg := range os.Args[2:] {
		switch arg {
		case "--multi-line":
			isMultiLine = true
		case "--encode":
			isEncode = true
		default:
			fmt.Println("Error: Unknown flag", arg)
			return
		}
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