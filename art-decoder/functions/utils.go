package functions

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ValidateBrackets ensures the encoded string has correctly balanced brackets.
// ValidateBrackets ensures the encoded string has correctly balanced brackets.
func ValidateBrackets(input string) error {
	stack := 0
	hasOpeningBracket := false // Track if at least one `[` appears

	for i, char := range input {
		if char == '[' {
			stack++
			hasOpeningBracket = true
		} else if char == ']' {
			stack--
			if stack < 0 {
				return errors.New("Error: Extra closing bracket found")
			}
		} else if i > 0 && input[i-1] == ']' && !unicode.IsSpace(char) && char != '[' && char != '-' && char != '*' && char != '"' && char != 'o' {
			// Allow certain characters like '-', '*', '"', 'o' after ']'
			continue
		}
	}

	if !hasOpeningBracket {
		return errors.New("Error: Missing opening bracket")
	}

	if stack > 0 {
		return errors.New("Error: Missing closing bracket")
	}

	return nil
}

// ValidateArguments checks if the arguments inside square brackets are valid.
func ValidateArguments(input string) error {
	// Use regex to extract bracketed sections
	pattern := regexp.MustCompile(`\[\s*(\d+)\s+([^\[\]]+)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) < 3 {
			continue // Skip invalid matches
		}

		// Extract count and character(s)
		count := match[1]
		character := match[2]

		// Validate count (must be a number)
		for _, c := range count {
			if !unicode.IsDigit(c) {
				return errors.New("Error: Invalid count inside brackets, must be a number")
			}
		}

		// Validate character (should not contain `]` or `[`, must be a single char or valid sequence)
		if strings.Contains(character, "[") || strings.Contains(character, "]") {
			return errors.New("Error: Invalid character inside brackets")
		}
	}

	// Check if there are any invalid bracketed sections
	invalidPattern := regexp.MustCompile(`\[[^\]]*\]`)
	invalidMatches := invalidPattern.FindAllString(input, -1)
	if len(invalidMatches) > len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	return nil
}