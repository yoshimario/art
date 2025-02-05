package functions

import (
	"errors"
	"regexp"
	"strconv"
)

// ValidateBrackets ensures the encoded string has correctly balanced brackets.
func ValidateBrackets(input string) error {
	stack := 0
	for _, char := range input { // Removed unused 'i'
		if char == '[' {
			stack++
		} else if char == ']' {
			if stack == 0 {
				return errors.New("Error: Extra closing bracket found")
			}
			stack--
		}
	}
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