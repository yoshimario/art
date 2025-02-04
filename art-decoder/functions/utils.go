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
	foundOpening := false

	for _, char := range input {
		if char == '[' {
			stack++
			foundOpening = true
		} else if char == ']' {
			if stack == 0 {
				return errors.New("Error: Extra closing bracket found")
			}
			stack--
		}
	}

	if !foundOpening {
		return errors.New("Error: Missing opening bracket")
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

	invalidPattern := regexp.MustCompile(`\[[^]]*\]`)
	invalidMatches := invalidPattern.FindAllString(input, -1)

	// If there are brackets that don't match the valid pattern
	if len(invalidMatches) != len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		count := match[1]
		char := match[2]

		if _, err := strconv.Atoi(count); err != nil {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}

		if strings.ContainsAny(char, "[]") {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}
	}

	return nil
}