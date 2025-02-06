package main

import (
	"art-decoder/functions"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func isInteractive() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

func main() {
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	flag.Parse()

	args := flag.Args()
	var input string

	// Use command-line arguments if provided
	if len(args) > 0 {
		input = strings.Join(args, " ")
	} else {
		// Multi-line input mode
		if *multiLine {
			// Only show prompt if running interactively (not in test mode)
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
			// Read single-line input from stdin
			reader := bufio.NewReader(os.Stdin)
			data, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				os.Exit(1)
			}
			input = strings.TrimSpace(data)
		}
	}

	// Decode using the appropriate function
	var output string
	var err error

	if *multiLine {
		output, err = functions.DecodeMultiLine(input)
	} else {
		output, err = functions.DecodeSingleLine(input)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}