package functions

import "strings"

// DecodeMultiLine processes multiple lines of encoded input.
func DecodeMultiLine(encodedString string) (string, error) {
	lines := strings.Split(encodedString, "\n")
	var result strings.Builder

	for _, line := range lines {
		decoded, err := DecodeSingleLine(line)
		if err != nil {
			return "", err
		}
		result.WriteString(decoded + "\n")
	}

	return strings.TrimRight(result.String(), "\n"), nil
}