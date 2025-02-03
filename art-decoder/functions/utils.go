package functions

import (
	"errors"
	"regexp"
)

// ValidateArguments checks that every bracketed sequence is of the form
// [count char] where count is one or more digits and char is one or more characters.
// Escaped brackets (i.e. "\[" or "\]") are allowed.
func ValidateArguments(input string) error {
	// This regex matches valid bracketed sequences.
	pattern := regexp.MustCompile(`$begin:math:display$\\s*(\\d+)\\s+((?:\\\\[\\[$end:math:display$]|[^$begin:math:display$$end:math:display$])+)\s*\]`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) < 3 {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}
	}

	// The invalidPattern finds any bracketed section.
	invalidPattern := regexp.MustCompile(`$begin:math:display$[^$end:math:display$]*\]`)
	invalidMatches := invalidPattern.FindAllString(input, -1)
	// If there are more bracketed sections than valid ones, at least one is invalid.
	if len(invalidMatches) > len(matches) {
		return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}

	return nil
}