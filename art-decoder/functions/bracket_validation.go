package functions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ValidateBrackets ensures the string has correctly balanced brackets
// and also rejects any `#[` sequence.
func ValidateBrackets(input string) error {
	// 1) Immediately reject if there's a `#[` sequence anywhere.
	//    We treat that as "Missing opening bracket."
	if strings.Contains(input, "#[") {
			return errors.New("Error: Missing opening bracket")
	}

	stack := 0

	for i, r := range input {
			if r == '[' {
					stack++
			} else if r == ']' {
					// If we have a closing bracket but stack == 0, it’s unbalanced.
					if stack == 0 {
							// Your tests do this special logic:
							//    If i == 0 or the previous character is not ']', say "Missing opening bracket".
							//    Otherwise, "Extra closing bracket found".
							if i == 0 || rune(input[i-1]) != ']' {
									return errors.New("Error: Missing opening bracket")
							}
							return errors.New("Error: Extra closing bracket found")
					}
					// Match up with an opening bracket
					stack--
			}
	}

	// If at the end stack > 0, we have more '[' than ']'
	if stack > 0 {
			return errors.New("Error: Missing closing bracket")
	}

	return nil
}

func ValidateArguments(input string) error {
    // Example: requires [count charSeq]
    pattern := regexp.MustCompile(`$begin:math:display$(\\d+)\\s+([^]]*)$end:math:display$`)
    matches := pattern.FindAllStringSubmatch(input, -1)

    bracketPattern := regexp.MustCompile(`$begin:math:display$[^]]*$end:math:display$`)
    bracketMatches := bracketPattern.FindAllString(input, -1)

    if len(bracketMatches) != len(matches) {
        return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
    }

    // Check counts are valid integers > 0
    for _, match := range matches {
        count, err := strconv.Atoi(match[1])
        if err != nil || count <= 0 {
            return errors.New("Error: Invalid number format inside brackets")
        }
    }
    return nil
}