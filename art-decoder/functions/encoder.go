package functions

import (
	"fmt"
	"strings"
)

// Encode encodes text-based art into the compressed format.
func Encode(input string) (string, error) {
	var result strings.Builder
	i := 0
	n := len(input)

	for i < n {
		currentChar := input[i]
		count := 1

		// Count consecutive occurrences of the current character
		for i+count < n && input[i+count] == currentChar {
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

	return result.String(), nil
}