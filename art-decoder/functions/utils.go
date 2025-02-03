package functions

import (
	"errors"
	"regexp"
)

// ValidateArguments checks that every bracketed sequence has the proper format.
// This regex accepts sequences of the form [count char] where count is one or more
// digits and char is one or more characters (allowing escaped brackets).
func ValidateArguments(input string) error {
	pattern := regexp.MustCompile(`\[\s*(\d+)\s+((?:\\[\[\]]|[^$begin:math:display$$end:math:display$])+)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) < 3 {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}
		// Additional checks (for example on the count or the character) can be added here.
	}

	// Ensure that any bracketed section matches the valid pattern.
	invalidPattern := regexp.MustCompile(`$begin:math:display$[^$end:math:display$]*\]`)
	invalidMatches := invalidPattern.FindAllString(input, -1)
	if len(invalidMatches) > len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	return nil
}

// ValidateBrackets ensures that all '[' and ']' are balanced.
// (This version simply counts brackets and does not reject literal text following a closing bracket.)
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
	if stack != 0 {
		return errors.New("Error: Mismatched brackets")
	}
	return nil
}