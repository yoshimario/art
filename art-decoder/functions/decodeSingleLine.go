package functions

import (
    "errors"
    "strconv"
    "strings"
)

func DecodeSingleLine(encoded string) (string, error) {
    // 1) Validate bracket structure/format first
    if err := ValidateBrackets(encoded); err != nil {
        return "", err
    }

    // 2) Now decode expansions [count chars]
    var result strings.Builder
    i := 0
    for i < len(encoded) {
        if encoded[i] == '[' {
            j := i + 1
            // Find either space or closing bracket to separate count
            for j < len(encoded) && encoded[j] != ' ' && encoded[j] != ']' {
                j++
            }
            if j >= len(encoded) || encoded[j] != ' ' {
                // Shouldn't happen if ValidateBrackets was correct, but just in case:
                return "", errors.New("Error: Invalid format inside brackets (expected '[count char]')")
            }

            // parse the integer
            countStr := encoded[i+1 : j]
            // safe to ignore the error because ValidateBrackets ensures it's valid integer
            count, _ := strconv.Atoi(countStr)

            j++ // skip the space
            charStart := j
            for j < len(encoded) && encoded[j] != ']' {
                j++
            }
            // guaranteed to find ']' because bracket validation ensures balanced
            charSeq := encoded[charStart:j]

            // Expand the substring
            result.WriteString(strings.Repeat(charSeq, count))
            i = j + 1 // move past ']'
        } else {
            // Normal character
            result.WriteByte(encoded[i])
            i++
        }
    }

    return result.String(), nil
}