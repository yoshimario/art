package main

import (
	"art-decoder/functions"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

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

	if *multiLine && isInteractive() {
		fmt.Println("Enter multi-line input. Press Ctrl+D to finish:")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	var output string
	if *multiLine {
		output, err = functions.DecodeMultiLine(string(input))
	} else {
		output, err = functions.DecodeSingleLine(string(input))
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}