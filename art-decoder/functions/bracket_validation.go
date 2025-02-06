package functions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ValidateBrackets ensures the encoded string has correctly balanced brackets.
func ValidateBrackets(input string) error {
	stack := 0
	firstOpenIndex := strings.Index(input, "[")
	firstCloseIndex := strings.Index(input, "]")

	// Detect misplaced closing bracket at the wrong position
	if firstCloseIndex != -1 && (firstOpenIndex == -1 || firstCloseIndex < firstOpenIndex) {
		return errors.New("Error: Missing opening bracket")
	}

	for i, char := range input {
		if char == '[' {
			stack++
		} else if char == ']' {
			if stack == 0 {
				// Check if this `]` is in a position where a `[` should have been
				if i == 0 || input[i-1] != ']' {
					return errors.New("Error: Missing opening bracket")
				}
				return errors.New("Error: Extra closing bracket found")
			}
			stack--
		}
	}

	// If there are unmatched opening brackets, return "Missing closing bracket"
	if stack > 0 {
		return errors.New("Error: Missing closing bracket")
	}

	return nil
}
// ValidateArguments checks if the arguments inside square brackets are valid.
func ValidateArguments(input string) error {
	pattern := regexp.MustCompile(`$begin:math:display$\\s*(\\d+)\\s+([^][]*)\\s*$end:math:display$`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	bracketPattern := regexp.MustCompile(`$begin:math:display$[^]]*$end:math:display$`)
	bracketMatches := bracketPattern.FindAllString(input, -1)

	if len(bracketMatches) != len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}
	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil || count <= 0 {
			return errors.New("Error: Invalid number format inside brackets")
		}
	}
	return nil
}