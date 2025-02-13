package functions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ValidateBrackets enforces bracket structure AND correct "[count char]" format.
func ValidateBrackets(input string) error {
	// 1) Immediately reject any occurrence of "#["
	//    => "Error: Missing opening bracket"
	if strings.Contains(input, "#[") {
			return errors.New("Error: Missing opening bracket")
	}

	// 2) Stack-based check for unbalanced or extra brackets
	stack := 0
	for i := 0; i < len(input); i++ {
			ch := input[i]
			if ch == '[' {
					stack++
			} else if ch == ']' {
					if stack == 0 {
							// Stray ']' with no matching '['
							// If preceding char != ']', => "Missing opening bracket"
							// else => "Extra closing bracket found"
							if i == 0 || input[i-1] != ']' {
									return errors.New("Error: Missing opening bracket")
							}
							return errors.New("Error: Extra closing bracket found")
					}
					stack--
			}
	}
	if stack > 0 {
			// Unclosed '['
			return errors.New("Error: Missing closing bracket")
	}

	// 3) Check bracket contents: "[<digits> <non-empty-chars>]"
	//    e.g. "[5 -_]" is valid; "[5#]" or "[5  ]" is invalid.
	bracketRe := regexp.MustCompile(`$begin:math:display$[^]]*$end:math:display$`)
	brackets := bracketRe.FindAllString(input, -1)
	for _, br := range brackets {
			// content = everything between '[' and ']'
			content := br[1 : len(br)-1]

			// Expect exactly 2 parts: "count" and "chars", separated by a space
			parts := strings.SplitN(content, " ", 2)
			if len(parts) < 2 {
					return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
			}

			countPart, charsPart := parts[0], parts[1]

			// Check that countPart is a valid integer
			if _, err := strconv.Atoi(countPart); err != nil {
					return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
			}

			// Trim the charsPart to reject "[5  ]" or "[10    ]" with only spaces
			if len(strings.TrimSpace(charsPart)) == 0 {
					return errors.New("Error: Invalid format inside brackets (expected '[count char]')")
			}
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