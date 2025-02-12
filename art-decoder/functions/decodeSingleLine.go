package functions

import (
    "strconv"
    "strings"
)

func DecodeSingleLine(encodedString string) (string, error) {
    // 1) Validate bracket structure/format
    if err := ValidateBrackets(encodedString); err != nil {
        return "", err
    }

    // 2) Expand bracket patterns [count chars]
    var result strings.Builder
    i := 0
    for i < len(encodedString) {
        if encodedString[i] == '[' {
            j := i + 1
            // Find first space or ']' after the count
            for j < len(encodedString) && encodedString[j] != ' ' && encodedString[j] != ']' {
                j++
            }
            // No need to check errors here, because bracket validation
            // guaranteed we have "[<digits> <non-empty>]"
            countStr := encodedString[i+1 : j]
            count, _ := strconv.Atoi(countStr) // safe, we validated it

            j++ // move past the space
            charStart := j
            for j < len(encodedString) && encodedString[j] != ']' {
                j++
            }
            // guaranteed to find the ']' by bracket validation
            charSeq := encodedString[charStart:j]
            // Expand it
            result.WriteString(strings.Repeat(charSeq, count))

            i = j + 1 // move beyond ']'
        } else {
            result.WriteByte(encodedString[i])
            i++
        }
    }

    return result.String(), nil
}