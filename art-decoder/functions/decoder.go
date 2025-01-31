package functions

import (
	"errors"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
func Decode(encodedString string) (string, error) {
	var result strings.Builder
	lines := strings.Split(encodedString, "\n")

	for lineIndex, line := range lines {
		i := 0

		// Handle leading spaces if present
		if strings.HasPrefix(line, "[") && strings.Contains(line, "]") {
			j := strings.Index(line, " ")
			k := strings.Index(line, "]")

			if j == -1 || k == -1 || j >= k {
				return "", errors.New("invalid leading spaces format")
			}

			countStr := line[1:j]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("invalid leading spaces count")
			}

			// Add the correct number of leading spaces
			result.WriteString(strings.Repeat(" ", count))
			i = k + 1 // Move past the closing bracket
		}

		// Decode the rest of the line
		for i < len(line) {
			if line[i] == '[' {
				// Extract count
				j := i + 1
				for j < len(line) && line[j] != ' ' {
					j++
				}
				if j >= len(line) {
					return "", errors.New("invalid format")
				}
				countStr := line[i+1 : j]
				count, err := strconv.Atoi(countStr)
				if err != nil {
					return "", errors.New("invalid count")
				}

				// Extract character(s) to repeat
				k := j + 1
				for k < len(line) && line[k] != ']' {
					k++
				}
				if k >= len(line) {
					return "", errors.New("unbalanced brackets")
				}

				// Extract character(s) to repeat
				char := line[j+1 : k]

				// Properly decode special bracket cases
				if char == "\\]" {
					char = "]"
				} else if char == "\\[" {
					char = "["
				}

				// Append repeated characters to result
				result.WriteString(strings.Repeat(char, count))

				i = k + 1
			} else {
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
		trimmed := strings.TrimRight(decodedLine, " ") // Prevent trailing spaces

		result.WriteString(trimmed)

		// Only add newline if it's not the last line
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}