package functions

import (
	"fmt"
	"strings"
)

// Encode encodes text-based art into the compressed format.
func Encode(input string) (string, error) {
	var result strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
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

			// If the character repeats more than once, encode it
			if count > 1 {
				result.WriteString(fmt.Sprintf("[%d %c]", count, currentChar))
			} else {
				result.WriteByte(currentChar)
			}

			i += count
		}

		// Add a newline character (except for the last line)
		if line != lines[len(lines)-1] {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}