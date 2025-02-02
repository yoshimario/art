package main

import (
	"art-decoder/functions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		runTests()
		return
	}

	var input string
	isMultiLine := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-ml", "--multi-line":
			isMultiLine = true
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

	decoded, err := functions.DecodeMultiLine(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decoded)
}