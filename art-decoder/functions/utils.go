package functions

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ValidateBrackets ensures the encoded string has correctly balanced brackets.
func ValidateBrackets(input string) error {
	stack := 0

	for _, char := range input {
		if char == '[' {
			stack++
		} else if char == ']' {
			stack--
			if stack < 0 {
				return errors.New("Error: Extra closing bracket found")
			}
		}
	}

	if stack > 0 {
		return errors.New("Error: Missing closing bracket")
	}

	return nil
}

// ValidateArguments checks if the arguments inside square brackets are valid.
func ValidateArguments(input string) error {
	pattern := regexp.MustCompile(`\[\s*(\d+)\s+([^\[\]]+)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	invalidPattern := regexp.MustCompile(`\[[^\]]*\]`)
	invalidMatches := invalidPattern.FindAllString(input, -1)

	if len(invalidMatches) != len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		count := match[1]
		for _, c := range count {
			if !unicode.IsDigit(c) {
				return errors.New("Error: Invalid count inside brackets, must be a number")
			}
		}

		character := match[2]
		if strings.ContainsAny(character, "[]") {
			return errors.New("Error: Invalid character inside brackets")
		}
	}

	return nil
}