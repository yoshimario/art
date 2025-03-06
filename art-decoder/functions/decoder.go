package functions

import (
	"errors"
	"strconv"
	"strings"
)

// DecodeMultiLine processes multiple lines of encoded input.
func DecodeMultiLine(encodedString string) (string, error) {
	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for _, line := range lines {
		decoded, err := DecodeSingleLine(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decoded + "\n")
	}

	return strings.TrimRight(result.String(), "\n"), nil
}

// DecodeSingleLine decodes a single line of compressed input.
func DecodeSingleLine(encoded string) (string, error) {
	var result strings.Builder
	i := 0

	for i < len(encoded) {
		if encoded[i] == '[' {
			// Start of a compressed pattern
			j := i + 1

			// Find the closing bracket
			for j < len(encoded) && encoded[j] != ']' {
				j++
			}

			if j >= len(encoded) {
				return "", errors.New("Error: Missing closing bracket")
			}

			// Extract the content inside the brackets
			content := encoded[i+1 : j]

			// Split the content into count and pattern
			// The count is the first part, and the pattern is everything after the first space
			spaceIndex := strings.Index(content, " ")
			if spaceIndex == -1 {
				return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
			}

			// Parse the count
			countStr := content[:spaceIndex]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("Error: Invalid count inside brackets")
			}

			// Extract the pattern (everything after the first space)
			pattern := content[spaceIndex+1:]

			// Repeat the pattern and add it to the result
			result.WriteString(strings.Repeat(pattern, count))

			// Move past the closing bracket
			i = j + 1
		} else {
			// Normal character, add it to the result
			result.WriteByte(encoded[i])
			i++
		}
	}

	return result.String(), nil
}