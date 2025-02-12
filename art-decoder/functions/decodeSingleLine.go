package functions

import (
    "errors"
    "strconv"
    "strings"
)

// DecodeSingleLine expands bracket patterns [count chars]
func DecodeSingleLine(encodedString string) (string, error) {
    if encodedString == "" {
        return "", nil
    }

    // 1) Validate brackets first
    if err := ValidateBrackets(encodedString); err != nil {
        return "", err
    }

    // 2) Then decode
    var result strings.Builder
    i := 0
    for i < len(encodedString) {
        ch := encodedString[i]

        if ch == '[' {
            // parse "[count char... ]"
            j := i + 1
            for j < len(encodedString) && encodedString[j] != ' ' && encodedString[j] != ']' {
                j++
            }
            if j >= len(encodedString) || encodedString[j] != ' ' {
                return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
            }

            countStr := encodedString[i+1 : j]
            count, err := strconv.Atoi(countStr)
            if err != nil {
                return "", errors.New("Error: Invalid number format inside brackets")
            }

            j++ // move past the space
            charStart := j
            for j < len(encodedString) && encodedString[j] != ']' {
                j++
            }
            if j >= len(encodedString) {
                return "", errors.New("Error: Missing closing bracket")
            }

            charSeq := encodedString[charStart:j]
            // Expand the bracketed substring count times
            result.WriteString(strings.Repeat(charSeq, count))

            // Advance index beyond the closing bracket
            i = j + 1
        } else {
            // normal character
            result.WriteByte(ch)
            i++
        }
    }

    return result.String(), nil
}