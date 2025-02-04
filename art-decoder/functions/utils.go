package functions

import (
	"errors"
	"regexp"
)

// ValidateBrackets ensures the encoded string has correctly balanced brackets.
func ValidateBrackets(input string) error {
	stack := 0

	for _, char := range input {
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
	// Regex to detect expected format '[count char]' with optional spaces
	pattern := regexp.MustCompile(`\[\s*(\d+)\s+([^][]*)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	// Find all bracket instances
	bracketPattern := regexp.MustCompile(`\[[^]]*\]`)
	bracketMatches := bracketPattern.FindAllString(input, -1)

	// If there are brackets that don't match the valid pattern
	if len(bracketMatches) != len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	return nil
}