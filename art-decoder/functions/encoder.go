package functions

import (
	"fmt"
	"strings"
)

// Encode encodes text-based art into the compressed format.
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

			// Count consecutive occurrences of the current character
			for i+count < len(line) && line[i+count] == currentChar {
				count++
			}

			// Encode characters properly, ensuring brackets `]` are handled safely
			if currentChar == ']' {
				result.WriteString(fmt.Sprintf("[%d %s]", count, "\\]"))
			} else {
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