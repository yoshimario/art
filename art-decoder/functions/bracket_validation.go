package functions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func ValidateBrackets(input string) error {
	// Reject "#[" as an incorrect opening bracket
	if strings.Contains(input, "#[") {
		return errors.New("Error: Missing opening bracket")
	}

	// Stack-based validation for bracket pairs
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

	// **Step 1: Match ALL bracket expressions**
	bracketRe := regexp.MustCompile(`$begin:math:display$\\s*\\d+\\s+[^$end:math:display$]*\s*\]`)
	allBrackets := bracketRe.FindAllString(input, -1)

	// **Step 2: Allow normal text (for decoding), but reject `[5  ]`**
	if len(allBrackets) == 0 {
		return nil // Allow decoder to process non-bracketed text
	}

	// **Step 3: Validate all bracket expressions**
	for _, br := range allBrackets {
		content := br[1 : len(br)-1] // Extract content inside brackets
		parts := strings.Fields(content)

		// Ensure exactly two parts: "count" and "chars"
		if len(parts) != 2 {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
		}

		countPart, charsPart := parts[0], parts[1]

		// Ensure countPart is a valid integer
		if _, err := strconv.Atoi(countPart); err != nil {
			return errors.New("Error: Invalid number format inside brackets")
		}

		// **Strict check: Reject `[5  ]` (empty character sequence)**
		if len(charsPart) == 0 || strings.TrimSpace(charsPart) == "" {
			return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
	}
	}

	return nil
}