package functions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ValidateBrackets enforces bracket structure AND correct "[count char]" format.
func ValidateBrackets(input string) error {
	// Immediately reject any occurrence of "#[" => "Error: Missing opening bracket"
	if strings.Contains(input, "#[") {
		return errors.New("Error: Missing opening bracket")
	}

	// Stack-based check for unbalanced or extra brackets
	stack := 0
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch == '[' {
			stack++
		} else if ch == ']' {
			if stack == 0 {
				if i == 0 || input[i-1] != ']' {
					return errors.New("Error: Missing opening bracket")
				}
				return errors.New("Error: Extra closing bracket found")
			}
			stack--
		}
	}
	if stack > 0 {
		return errors.New("Error: Missing closing bracket")
	}

	// Regex to find all bracketed segments
	bracketRe := regexp.MustCompile(`$begin:math:display$[^$end:math:display$]*\]`)
	brackets := bracketRe.FindAllString(input, -1)
	for _, br := range brackets {
		content := br[1 : len(br)-1] // Extract content inside brackets

		// Ensure exactly two parts: "count" and "chars"
		parts := strings.SplitN(content, " ", 2)
		if len(parts) < 2 {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}

		countPart, charsPart := parts[0], parts[1]

		// Check that countPart is a valid integer
		if _, err := strconv.Atoi(countPart); err != nil {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}

		// Trim the charsPart to reject "[5  ]" or similar cases with only spaces
		if strings.TrimSpace(charsPart) == "" {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}
	}

	return nil
}