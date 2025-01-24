package functions

import (
	"errors"
	"strconv"
	"strings"
)

// Decode decodes a single-line encoded string into text-based art.
func Decode(encodedString string) (string, error) {
	var result strings.Builder
	i := 0
	for i < len(encodedString) {
		if encodedString[i] == '[' {
			// Extract count
			j := i + 1
			for j < len(encodedString) && encodedString[j] != ' ' {
				j++
			}
			if j >= len(encodedString) {
				return "", errors.New("invalid format")
			}
			countStr := encodedString[i+1 : j]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("invalid count")
			}

			// Extract character(s) to repeat
			k := j + 1
			for k < len(encodedString) && encodedString[k] != ']' {
				k++
			}
			if k >= len(encodedString) {
				return "", errors.New("unbalanced brackets")
			}
			char := encodedString[j+1 : k]

			// Append repeated characters to result
			result.WriteString(strings.Repeat(char, count))
			i = k + 1
		} else {
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

	for _, line := range lines {
		decodedLine, err := Decode(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decodedLine + "\n")
	}

	return result.String(), nil
}