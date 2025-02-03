package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DecodeMultiLine processes multi-line encoded input
func DecodeMultiLine(encodedString string) (string, error) {
	if !strings.Contains(encodedString, "[") {
			return "", errors.New("Error: Missing opening bracket") // ✅ Fix now flags missing `[`
	}

	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for i, line := range lines {
			decodedLine, err := Decode(line)
			if err != nil {
					return "", err
			}

			result.WriteString(decodedLine)

			if i < len(lines)-1 {
					result.WriteString("\n")
			}
	}
	return result.String(), nil
}

// Decode processes a single encoded line
func Decode(encodedString string) (string, error) {
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
					j := i + 1
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

					k := j + 1
					for k < len(encodedString) && encodedString[k] != ']' {
							k++
					}
					if k >= len(encodedString) {
							return "", errors.New("Error: Missing closing bracket")
					}

					char := encodedString[j+1 : k]
					result.WriteString(strings.Repeat(char, count)) // ✅ Fix ensures no extra character

					i = k + 1
			} else {
					result.WriteByte(encodedString[i])
					i++
			}
	}
	return result.String(), nil
}