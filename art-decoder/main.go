package main

import (
	"art/art-decoder/functions"
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
	
	// Sound effect flags
	soundMode := flag.Bool("sound", false, "Enable sound effects")
	soundType := flag.String("sound-type", "typing", "Sound type: beep, success, error, typing, loading, banner")
	soundVolume := flag.Int("volume", 50, "Sound volume (0-100)")
	soundPitch := flag.Float64("pitch", 1.0, "Sound pitch (0.5-2.0)")
	
	flag.Parse()

	// Setup animation configuration
	animConfig := functions.NewDefaultAnimationConfig()
	animConfig.Enabled = *animateMode
	animConfig.Type = functions.Animation(*animationType)
	animConfig.Speed = time.Duration(*animationSpeed) * time.Millisecond
	animConfig.ColorOutput = !*noColor

	// Setup sound configuration
	soundConfig := functions.NewDefaultSoundConfig()
	soundConfig.Enabled = *soundMode
	soundConfig.Type = functions.SoundType(*soundType)
	soundConfig.Volume = *soundVolume
	soundConfig.Pitch = *soundPitch

	// Display title if animations are enabled
	if animConfig.Enabled && functions.IsInteractive() {
		functions.AnimateTitle("ART DECODER", animConfig)
		
		// Play banner sound if sound effects are enabled
		if soundConfig.Enabled {
			titleSound := soundConfig
			titleSound.Type = functions.SoundBanner
			functions.PlaySound(titleSound)
		}
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
				
				// Play typing sound for prompt
				if soundConfig.Enabled {
					promptSound := soundConfig
					promptSound.Type = functions.SoundTyping
					functions.PlaySound(promptSound)
				}
			}
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				functions.PrintRed("Error reading input: " + err.Error())
				
				// Play error sound
				if soundConfig.Enabled {
					errorSound := soundConfig
					errorSound.Type = functions.SoundError
					functions.PlaySound(errorSound)
				}
				
				os.Exit(1)
			}
			input = strings.Join(lines, "\n")
		} else {
			// Single-line mode from stdin
			reader := bufio.NewReader(os.Stdin)
			data, err := reader.ReadString('\n')
			if err != nil {
				functions.PrintRed("Error reading input: " + err.Error())
				
				// Play error sound
				if soundConfig.Enabled {
					errorSound := soundConfig
					errorSound.Type = functions.SoundError
					functions.PlaySound(errorSound)
				}
				
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
		
		// Play loading sound
		if soundConfig.Enabled {
			loadingSound := soundConfig
			loadingSound.Type = functions.SoundLoading
			loadingSound.RepeatCount = 3
			functions.PlaySound(loadingSound)
		}
	}

	// Step 1: Validate Input Before Processing
	err := functions.ValidateBrackets(input)
	if err != nil {
		functions.PrintRed(err.Error())
		
		// Play error sound
		if soundConfig.Enabled {
			errorSound := soundConfig
			errorSound.Type = functions.SoundError
			functions.PlaySound(errorSound)
		}
		
		os.Exit(1)
	}

	// Step 2: Check If Encoding Mode Is Enabled
	if *encodeMode {
		output, err := functions.Encode(input)
		if err != nil {
			functions.PrintRed(err.Error())
			
			// Play error sound
			if soundConfig.Enabled {
				errorSound := soundConfig
				errorSound.Type = functions.SoundError
				functions.PlaySound(errorSound)
			}
			
			os.Exit(1)
		}
		
		// Play success sound
		if soundConfig.Enabled {
			successSound := soundConfig
			successSound.Type = functions.SoundSuccess
			functions.PlaySound(successSound)
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
		
		// Play error sound
		if soundConfig.Enabled {
			errorSound := soundConfig
			errorSound.Type = functions.SoundError
			functions.PlaySound(errorSound)
		}
		
		os.Exit(1)
	}

	// Play success sound
	if soundConfig.Enabled {
		successSound := soundConfig
		successSound.Type = functions.SoundSuccess
		functions.PlaySound(successSound)
	}

	// Animate or print the output
	functions.AnimateOutput(output, animConfig)
	
	// Play final beep to indicate completion
	if soundConfig.Enabled && soundConfig.Type != functions.SoundSuccess {
		completeSound := soundConfig
		completeSound.Type = functions.SoundBeep
		completeSound.Pitch = 1.5  // Higher pitch for completion
		functions.PlaySound(completeSound)
	}
}