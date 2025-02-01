package functions

import (
	"errors"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
// Decode decodes a single-line encoded string into text-based art.
// It takes an encoded string as input and returns the decoded string and an error.
// The encoded string is expected to have a specific format with bracketed counts and characters.
// This function handles the decoding logic, including handling leading spaces and escaped special characters.
func Decode(encodedString string) (string, error) {
	var result strings.Builder
	lines := strings.Split(encodedString, "\n")

	for lineIndex, line := range lines {
		i := 0

		// Decode the line
		for i < len(line) {
			if line[i] == '[' {
				// Extract count
				j := i + 1
				for j < len(line) && line[j] != ' ' {
					j++
				}
				if j >= len(line) {
					return "", errors.New("invalid format: missing space after count")
				}
				countStr := line[i+1 : j]
				count, err := strconv.Atoi(countStr)
				if err != nil {
					return "", errors.New("invalid count: " + countStr)
				}

				// Extract character(s) to repeat
				k := j + 1
				for k < len(line) && line[k] != ']' {
					k++
				}
				if k >= len(line) {
					return "", errors.New("unbalanced brackets: missing closing bracket")
				}

				char := line[j+1 : k]

				// Handle escaped special characters
				if char == "\\]" {
					char = "]"
				} else if char == "\\[" {
					char = "["
				}

				// Append repeated characters to result
				result.WriteString(strings.Repeat(char, count))

				i = k + 1 // Move past the closing bracket
			} else {
				// Append regular characters to result
				result.WriteByte(line[i])
				i++
			}
		}

		// Add a newline character if it's not the last line
		if lineIndex < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

// DecodeMultiLine decodes a multi-line encoded string into text-based art.
func DecodeMultiLine(encodedString string) (string, error) {
	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for i, line := range lines {
		decodedLine, err := Decode(line)
		if err != nil {
			return "", err
		}

		// Ensure spaces and characters are not trimmed incorrectly
		result.WriteString(decodedLine)

		// Only add newline if it's not the last line
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}