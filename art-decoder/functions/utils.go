package functions

import "errors"

// ValidateBrackets checks if the square brackets in the input are balanced.
func ValidateBrackets(input string) bool {
	stack := 0
	for _, char := range input {
		if char == '[' {
			stack++
		} else if char == ']' {
			stack--
			if stack < 0 {
				return false
			}
		}
	}
	return stack == 0
}

// ValidateArguments checks if the arguments inside square brackets are valid.
func ValidateArguments(input string) error {
	if !ValidateBrackets(input) {
		return errors.New("unbalanced brackets")
	}

	// Additional validation logic can be added here if needed
	return nil
}