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
	hasOpeningBracket := false

	for _, char := range input {
		if char == '[' {
			stack++
			hasOpeningBracket = true
		} else if char == ']' {
			stack--
			if stack < 0 {
				return errors.New("Error: Extra closing bracket found")
			}
		}
	}

	if !hasOpeningBracket && strings.Contains(input, "]") {
		return errors.New("Error: Missing opening bracket")
	}

	if stack > 0 {
		return errors.New("Error: Missing closing bracket")
	}

	return nil
}

// ValidateArguments ensures the `[count char]` format is valid.
func ValidateArguments(input string) error {
	pattern := regexp.MustCompile(`\[\s*(\d+)\s+([^\[\]]+)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 && strings.Contains(input, "[") {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		count := match[1]
		character := match[2]

		for _, c := range count {
			if !unicode.IsDigit(c) {
				return errors.New("Error: Invalid count inside brackets, must be a number")
			}
		}

		if strings.TrimSpace(character) == "" {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}
	}

	return nil
}