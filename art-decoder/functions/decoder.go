package functions

import (
	"errors"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
func Decode(encodedString string) (string, error) {
	// Allow empty (or all whitespace) lines.
	if strings.TrimSpace(encodedString) == "" {
		return "", nil
	}

	// First validate that all bracketed sequences follow the proper format.
	if err := ValidateArguments(encodedString); err != nil {
		return "", err
	}

	var result strings.Builder
	lines := strings.Split(encodedString, "\n")

	// Process each line separately.
	for _, line := range lines {
		line = strings.TrimSpace(line) // Trim whitespace

		openBracketCount := 0
		closeBracketCount := 0

		i := 0
		for i < len(line) {
			if line[i] == '[' {
				openBracketCount++
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
					return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
				}

				k := j + 1
				for k < len(line) && line[k] != ']' {
					k++
				}
				if k >= len(line) {
					return "", errors.New("Error: Missing closing bracket")
				}

				charSeq := line[j+1 : k]
				if charSeq == `\]` {
					charSeq = "]"
				} else if charSeq == `\[` {
					charSeq = "["
				}

				result.WriteString(strings.Repeat(charSeq, count))
				i = k + 1
			} else if line[i] == ']' {
				closeBracketCount++
				if closeBracketCount > openBracketCount {
					return "", errors.New("Error: Extra closing bracket found")
				}
			} else if line[i] == '-' || line[i] == '|' || line[i] == ' ' {
				result.WriteByte(line[i])
				i++
				continue
			} else {
				return "", errors.New("Error: Missing opening bracket")
			}
		}

		if openBracketCount > closeBracketCount {
			return "", errors.New("Error: Missing closing bracket")
		}

		result.WriteString("\n")
	}

	s := result.String()
	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return s, nil
}

// DecodeMultiLine decodes a multi-line encoded string into text-based art.
func DecodeMultiLine(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil // Allow empty input
	}

	lines := strings.Split(input, "\n") // Split input into lines
	var decodedLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}

		decoded, err := Decode(line) // Call the existing Decode function
		if err != nil {
			return "", err // Return error immediately if any line fails
		}

		decodedLines = append(decodedLines, decoded)
	}

	return strings.Join(decodedLines, "\n"), nil
}