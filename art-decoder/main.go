package main

import (
	"art-decoder/functions"
	"bufio"
	"flag"
	"os"
	"strings"
	"time"
)

func main() {
	// Define command line flags
	multiLine := flag.Bool("ml", false, "Decode multi-line input")
	encodeMode := flag.Bool("encode", false, "Encode input text")
	
	// Animation flags
	animateMode := flag.Bool("animate", false, "Enable animation effects")
	animationType := flag.String("animation-type", "typing", "Animation type: typing, rainbow, banner, loading")
	animationSpeed := flag.Int("speed", 30, "Animation speed (lower is faster)")
	noColor := flag.Bool("no-color", false, "Disable color output")
	
	flag.Parse()

	// Setup animation configuration
	animConfig := functions.NewDefaultAnimationConfig()
	animConfig.Enabled = *animateMode
	animConfig.Type = functions.Animation(*animationType)
	animConfig.Speed = time.Duration(*animationSpeed) * time.Millisecond
	animConfig.ColorOutput = !*noColor

	// Display title if animations are enabled
	if animConfig.Enabled && functions.IsInteractive() {
		functions.AnimateTitle("ART DECODER", animConfig)
	}

	args := flag.Args()
	var input string

	// Read input from CLI arguments or stdin
	if len(args) > 0 {
		input = strings.Join(args, " ")
	} else {
		// If no arguments given, read either multi-line or single-line from stdin
		if *multiLine || *encodeMode {
			if functions.IsInteractive() {
				functions.PrintCyan("Enter multi-line input. Press Ctrl+D (Linux/macOS) or Ctrl+Z (Windows) to finish:")
			}
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				functions.PrintRed("Error reading input: " + err.Error())
				os.Exit(1)
			}
			input = strings.Join(lines, "\n")
		} else {
			// Single-line mode from stdin
			reader := bufio.NewReader(os.Stdin)
			data, err := reader.ReadString('\n')
			if err != nil {
				functions.PrintRed("Error reading input: " + err.Error())
				os.Exit(1)
			}
			input = strings.TrimSpace(data)
		}
	}

	// Show loading animation if enabled
	if animConfig.Enabled && animConfig.Type == functions.AnimationLoading && functions.IsInteractive() {
		loadingConfig := animConfig
		loadingConfig.Type = functions.AnimationLoading
		functions.AnimateOutput("", loadingConfig)
	}

	// Step 1: Validate Input Before Processing
	err := functions.ValidateBrackets(input)
	if err != nil {
		functions.PrintRed(err.Error())
		os.Exit(1)
	}

	// Step 2: Check If Encoding Mode Is Enabled
	if *encodeMode {
		output, err := functions.Encode(input)
		if err != nil {
			functions.PrintRed(err.Error())
			os.Exit(1)
		}
		
		// Animate or print the output
		functions.AnimateOutput(output, animConfig)
		return
	}

	// Step 3: Decoding Mode
	var output string
	if *multiLine {
		output, err = functions.DecodeMultiLine(input)
	} else {
		output, err = functions.DecodeSingleLine(input)
	}

	// Step 4: If there's an error, print it and exit
	if err != nil {
		functions.PrintRed(err.Error())
		os.Exit(1)
	}

	// Animate or print the output
	functions.AnimateOutput(output, animConfig)
}