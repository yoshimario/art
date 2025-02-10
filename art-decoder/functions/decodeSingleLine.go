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
	i := 0
	for i < len(encodedString) {
		if encodedString[i] == '[' {
			// Handle [count char] sequence
			j := i + 1
			for j < len(encodedString) && encodedString[j] != ' ' && encodedString[j] != ']' {
				j++
			}
			if j >= len(encodedString) || encodedString[j] != ' ' {
				return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
			}
			count, err := strconv.Atoi(encodedString[i+1 : j])
			if err != nil {
				return "", errors.New("Error: Invalid number format inside brackets")
			}
			j++
			charStart := j
			for j < len(encodedString) && encodedString[j] != ']' {
				j++
			}
			if j >= len(encodedString) {
				return "", errors.New("Error: Missing closing bracket")
			}
			charSeq := encodedString[charStart:j]
			result.WriteString(strings.Repeat(charSeq, count))
			i = j + 1
		} else {
			// Handle characters outside of [count char] sequences
			result.WriteByte(encodedString[i])
			i++
		}
	}
	return result.String(), nil
}