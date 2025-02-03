package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
func Decode(encodedString string) (string, error) {
	// Allow empty (or all whitespace) lines.
	if strings.TrimSpace(encodedString) == "" {
		return "", nil
	}

	// First validate the encoded sequences and bracket balance.
	if err := ValidateArguments(encodedString); err != nil {
		return "", err
	}
	if err := ValidateBrackets(encodedString); err != nil {
		return "", err
	}

	var result strings.Builder
	i := 0
	for i < len(encodedString) {
		if encodedString[i] == '[' {
			// Process an encoded sequence.
			j := i + 1
			// Find the space that separates the count from the character(s).
			for j < len(encodedString) && encodedString[j] != ' ' {
				j++
			}
			if j >= len(encodedString) {
				return "", errors.New("Error: Missing space after count")
			}

			countStr := encodedString[i+1 : j]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", fmt.Errorf("Error: Invalid count format (%s)", countStr)
			}

			// Find the closing bracket.
			k := j + 1
			for k < len(encodedString) && encodedString[k] != ']' {
				k++
			}
			if k >= len(encodedString) {
				return "", errors.New("Error: Missing closing bracket")
			}

			charSeq := encodedString[j+1 : k]
			// Handle escaped characters.
			if charSeq == `\]` {
				charSeq = "]"
			} else if charSeq == `\[` {
				charSeq = "["
			}

			result.WriteString(strings.Repeat(charSeq, count))
			i = k + 1
		} else {
			// For any character outside an encoded sequence, simply append it.
			result.WriteByte(encodedString[i])
			i++
		}
	}
	return result.String(), nil
}

// DecodeMultiLine decodes a multi-line encoded string into text-based art.
func DecodeMultiLine(encodedString string) (string, error) {
	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for idx, line := range lines {
		decodedLine, err := Decode(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decodedLine)
		// Append newline if this isn’t the last line.
		if idx < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}