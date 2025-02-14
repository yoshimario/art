package functions

import (
	"fmt"
	"strings"
)

// Encode converts text-based art into a compressed format.
func Encode(input string) (string, error) {
	var result strings.Builder
	lines := strings.Split(input, "\n")

	for lineIndex, line := range lines {
		// Count leading spaces
		leadingSpaces := 0
		for leadingSpaces < len(line) && line[leadingSpaces] == ' ' {
			leadingSpaces++
		}

		// Encode leading spaces
		if leadingSpaces > 0 {
			result.WriteString(fmt.Sprintf("[%d  ]", leadingSpaces))
		}

		// Encode the rest of the line
		i := leadingSpaces
		for i < len(line) {
			currentChar := line[i]
			count := 1

			// **Fix:** Group repeated character sequences correctly (handles `- _` properly)
			for i+count < len(line) && line[i+count] == currentChar {
				count++
			}

			// Handle escaping brackets `[]`
			switch currentChar {
			case ']':
				result.WriteString(fmt.Sprintf("[%d \\]]", count))
			case '[':
				result.WriteString(fmt.Sprintf("[%d \\[]", count))
			default:
				result.WriteString(fmt.Sprintf("[%d %c]", count, currentChar))
			}

			i += count
		}

		// Add a newline character if it's not the last line
		if lineIndex < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}