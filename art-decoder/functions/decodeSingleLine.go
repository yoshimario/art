package functions

import (
	"errors"
	"strconv"
	"strings"
)

// DecodeSingleLine decodes a single-line encoded string into text-based art.
func DecodeSingleLine(encodedString string) (string, error) {
	if encodedString == "" {
		return "", nil
	}

	if err := ValidateBrackets(encodedString); err != nil {
		return "", err
	}

	if err := ValidateArguments(encodedString); err != nil {
		return "", err
	}

	var result strings.Builder
	lines := strings.Split(encodedString, "\n")

	for _, line := range lines {
		i := 0
		for i < len(line) {
			if line[i] == '[' {
				j := i + 1
				for j < len(line) && line[j] != ' ' && line[j] != ']' {
					j++
				}
				if j >= len(line) || line[j] != ' ' {
					return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
				}
				count, err := strconv.Atoi(line[i+1 : j])
				if err != nil {
					return "", errors.New("Error: Invalid number format inside brackets")
				}
				j++
				charStart := j
				for j < len(line) && line[j] != ']' {
					j++
				}
				if j >= len(line) {
					return "", errors.New("Error: Missing closing bracket")
				}
				charSeq := line[charStart:j]
				result.WriteString(strings.Repeat(charSeq, count))
				i = j + 1
			} else {
				result.WriteByte(line[i])
				i++
			}
		}
	}
	return result.String(), nil
}