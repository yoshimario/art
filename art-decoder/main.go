package main

import (
	"art-decoder/functions"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// isInteractive checks if we're running in an interactive terminal
func isInteractive() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

func main() {
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	encodeMode := flag.Bool("encode", false, "Encode input text")
	flag.Parse()

	args := flag.Args()
	var input string

	// Read input from CLI arguments or stdin
	if len(args) > 0 {
		input = strings.Join(args, " ")
	} else {
		// If no arguments given, read either multi-line or single-line from stdin
		if *multiLine || *encodeMode {
			if isInteractive() {
				fmt.Println("Enter multi-line input. Press Ctrl+D (Linux/macOS) or Ctrl+Z (Windows) to finish:")
			}
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				os.Exit(1)
			}
			input = strings.Join(lines, "\n")
		} else {
			// Single-line mode from stdin
			reader := bufio.NewReader(os.Stdin)
			data, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				os.Exit(1)
			}
			input = strings.TrimSpace(data)
		}
	}

	// **Step 1: Validate Input Before Processing**
	err := functions.ValidateBrackets(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// **Step 2: Check If Encoding Mode Is Enabled**
	if *encodeMode {
		output, err := functions.Encode(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(output)
		return
	}

	// **Step 3: Decoding Mode**
	var output string
	if *multiLine {
		output, err = functions.DecodeMultiLine(input)
	} else {
		output, err = functions.DecodeSingleLine(input)
	}

	// **Step 4: If there's an error, print it and exit**
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Otherwise, print the decoded output
	fmt.Println(output)
}