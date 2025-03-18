package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// Encode converts ASCII art into a compressed format while preserving structure.
func Encode(input string) (string, error) {
	var result strings.Builder
	lines := strings.Split(input, "\n")

	for lineIndex, line := range lines {
		// Process each line to encode repeating patterns
		processedLine := encodeSpecialPatterns(line)
		result.WriteString(processedLine)

		// Add newline character if it's not the last line
		if lineIndex < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

// encodeSpecialPatterns handles specific patterns known to appear in the ASCII art
func encodeSpecialPatterns(line string) string {
	// Define common patterns to replace
	patterns := map[string]string{
		"/^--^\\":     "/^--^\\",
		"\\____/":     "\\____/",
		"/      \\":   "/      \\",
		"|        |":  "|        |",
		"\\__  __/":   "\\__  __/",
		"\\**  **/":   "\\**  **/",
		"\\**  __/":   "\\**  __/",
	}

	// Replace known patterns with placeholders
	processedLine := line
	for pattern, replacement := range patterns {
		processedLine = strings.ReplaceAll(processedLine, pattern, replacement)
	}

	// Handle leading spaces
	spacesPrefix := countLeadingSpaces(processedLine)
	if spacesPrefix > 0 {
		processedLine = fmt.Sprintf("[%d  ]%s", spacesPrefix, strings.TrimLeft(processedLine, " "))
	}

	// Handle repeating patterns
	processedLine = encodeRepeatingPatterns(processedLine)

	return processedLine
}

// encodeRepeatingPatterns encodes repeating patterns like "| | | |", "|^|^|^", etc.
func encodeRepeatingPatterns(line string) string {
	// Handle consecutive "| " pairs (e.g., "| | | |")
	re := regexp.MustCompile(`(\| ){2,}`)
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		pairCount := len(s) / 2
		return fmt.Sprintf("[%d | ]", pairCount)
	})

	// Handle consecutive "|^" pairs (e.g., "|^|^|^")
	re = regexp.MustCompile(`(\|[\^]){2,}`)
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		pairCount := len(s) / 2
		return fmt.Sprintf("[%d |^]", pairCount)
	})

	// Handle consecutive "#" symbols (e.g., "########")
	re = regexp.MustCompile(`#{2,}`)
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		return fmt.Sprintf("[%d #]", len(s))
	})

	// Handle consecutive spaces (e.g., "    ")
	re = regexp.MustCompile(`\s{2,}`)
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		return fmt.Sprintf("[%d  ]", len(s))
	})

	return line
}

// countLeadingSpaces counts the number of spaces at the beginning of a string
func countLeadingSpaces(s string) int {
	count := 0
	for i := 0; i < len(s) && s[i] == ' '; i++ {
		count++
	}
	return count
}