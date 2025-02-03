package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
func Decode(encodedString string) (string, error) {
	if err := ValidateBrackets(encodedString); err != nil {
			return "", err
	}

	var result strings.Builder
	lines := strings.Split(encodedString, "\n")

	for lineIndex, line := range lines {
			i := 0
			for i < len(line) {
					if line[i] == '[' {
							j := i + 1
							for j < len(line) && line[j] != ' ' {
									j++
							}
							if j >= len(line) {
									return "", errors.New("Error: Missing space after count")
							}
							countStr := line[i+1 : j]
							count, err := strconv.Atoi(countStr)
							if err != nil {
									return "", fmt.Errorf("Error: Invalid count format (%s)", countStr)
							}

							k := j + 1
							for k < len(line) && line[k] != ']' {
									k++
							}
							if k >= len(line) {
									return "", errors.New("Error: Missing closing bracket")
							}

							char := line[j+1 : k]

							// Handle escaped special characters
							if char == "\\]" {
									char = "]"
							} else if char == "\\[" {
									char = "["
							}

							result.WriteString(strings.Repeat(char, count))

							i = k + 1
					} else {
							result.WriteByte(line[i])
							i++
					}
			}
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

		result.WriteString(decodedLine)

		// Only add newline if it's not the last line
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}