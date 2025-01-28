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

	for _, line := range lines {
		i := 0
		// Handle leading spaces
		if strings.HasPrefix(line, "[") {
			// Extract the number of leading spaces
			j := strings.Index(line, " ")
			if j == -1 {
				return "", errors.New("invalid leading spaces format")
			}
			countStr := line[1:j]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("invalid leading spaces count")
			}

			// Add the leading spaces to the result
			result.WriteString(strings.Repeat(" ", count))
			i = j + 2 // Skip past the "[N  ]" part
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
				char := line[j+1 : k]

				// Append repeated characters to result
				result.WriteString(strings.Repeat(char, count))
				i = k + 1
			} else {
				result.WriteByte(line[i])
				i++
			}
		}

		// Add a newline character (except for the last line)
		if line != lines[len(lines)-1] {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
// DecodeMultiLine decodes a multi-line encoded string into text-based art.
func DecodeMultiLine(encodedString string) (string, error) {
	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for _, line := range lines {
		decodedLine, err := Decode(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decodedLine + "\n")
	}

	return result.String(), nil
}