package main

import (
	"art-decoder/functions"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// isInteractive checks if we're running in an interactive terminal
func isInteractive() bool {
    info, err := os.Stdin.Stat()
    if err != nil {
        return false
    }
    return (info.Mode() & os.ModeCharDevice) != 0
}

func main() {
    multiLine := flag.Bool("ml", false, "Decode multi-line input")
    flag.Parse()

    args := flag.Args()
    var input string

    // 1) If there are command-line arguments, join them into one string
    //    e.g. go run . "[5 _]" becomes "[5 _]"
    //         go run . 5 #[5 -_]-5 # becomes "5 #[5 -_]-5 #"
    if len(args) > 0 {
        input = strings.Join(args, " ")
    } else {
        // 2) If no arguments given, read either multi-line or single-line from stdin
        if *multiLine {
            // Multi-line mode
            if isInteractive() {
                fmt.Println("Enter multi-line input. Press Ctrl+D (Linux/macOS) or Ctrl+Z (Windows) to finish:")
            }
            scanner := bufio.NewScanner(os.Stdin)
            var lines []string
            for scanner.Scan() {
                lines = append(lines, scanner.Text())
            }
            if err := scanner.Err(); err != nil {
                fmt.Fprintln(os.Stderr, "Error reading input:", err)
                os.Exit(1)
            }
            input = strings.Join(lines, "\n")
        } else {
            // Single-line mode from stdin
            reader := bufio.NewReader(os.Stdin)
            data, err := reader.ReadString('\n')
            if err != nil {
                fmt.Fprintln(os.Stderr, "Error reading input:", err)
                os.Exit(1)
            }
            input = strings.TrimSpace(data)
        }
    }

    // 3) Call the appropriate decoder
    var output string
    var err error
    if *multiLine {
        // Multi-line decode
        output, err = functions.DecodeMultiLine(input)
    } else {
        // Single-line decode
        output, err = functions.DecodeSingleLine(input)
    }

    // 4) If there's an error from decoding, print it and exit
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // 5) Otherwise, print the decoded output
    fmt.Println(output)
}