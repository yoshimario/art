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
		// First handle the special patterns in each line
		processedLine := encodeSpecialPatterns(line)
		
		// Then encode the processed line
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
	// First handle the cat face and box patterns
	catFacePatterns := map[string]string{
		"/^--^\\":     "/^--^\\",
		"\\____/":     "\\____/",
		"/      \\":   "/      \\",
		"|        |":  "|        |",
		"\\__  **/":   "\\__  **/",
		"\\**  **/":   "\\**  **/",
		"\\**  __/":   "\\**  __/",
	}

	// Pre-process the line by replacing the cat face patterns with placeholders
	processedLine := line
	for pattern, replacement := range catFacePatterns {
		processedLine = strings.ReplaceAll(processedLine, pattern, replacement)
	}

	// Handle spaces at the beginning of a line
	spacesPrefix := countLeadingSpaces(processedLine)
	if spacesPrefix > 0 {
		processedLine = fmt.Sprintf("[%d  ]%s", spacesPrefix, processedLine[spacesPrefix:])
	}

	// Handle consecutive hash (#) symbols
	re := regexp.MustCompile(`#{4,}`)
	processedLine = re.ReplaceAllStringFunc(processedLine, func(s string) string {
		return fmt.Sprintf("[%d #]", len(s))
	})

	// Handle the most common pattern: consecutive "| " pairs
	// This must capture patterns like "| | | | | |"
	re = regexp.MustCompile(`(\| ){4,}`)
	processedLine = re.ReplaceAllStringFunc(processedLine, func(s string) string {
		pairCount := len(s) / 2
		return fmt.Sprintf("[%d | ]", pairCount)
	})

	// Special pattern for "|^|^|^|^|" sequences
	re = regexp.MustCompile(`(\|[\^]){4,}`)
	processedLine = re.ReplaceAllStringFunc(processedLine, func(s string) string {
		pairCount := len(s) / 2
		return fmt.Sprintf("[%d |^]", pairCount)
	})

	// Special pattern for "\^" and "/^" sequences
	processedLine = strings.ReplaceAll(processedLine, "\\^", "\\^")
	processedLine = strings.ReplaceAll(processedLine, "/^", "/^")

	// Handle spaces between patterns
	re = regexp.MustCompile(`\s{3,}`)
	processedLine = re.ReplaceAllStringFunc(processedLine, func(s string) string {
		return fmt.Sprintf("[%d  ]", len(s))
	})

	return processedLine
}

// countLeadingSpaces counts the number of spaces at the beginning of a string
func countLeadingSpaces(s string) int {
	count := 0
	for i := 0; i < len(s) && s[i] == ' '; i++ {
		count++
	}
	return count
}