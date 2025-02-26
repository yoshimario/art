package functions

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Animation represents the type of animation to perform
type Animation string

const (
	// Animation types
	AnimationTyping   Animation = "typing"
	AnimationRainbow  Animation = "rainbow"
	AnimationBanner   Animation = "banner"
	AnimationLoading  Animation = "loading"
	AnimationNone     Animation = "none"
)

// AnimationConfig stores configuration for animations
type AnimationConfig struct {
	Enabled     bool
	Type        Animation
	Speed       time.Duration
	Cycles      int
	Width       int
	ColorOutput bool
}

// NewDefaultAnimationConfig creates a default animation configuration
func NewDefaultAnimationConfig() AnimationConfig {
	return AnimationConfig{
		Enabled:     false,
		Type:        AnimationNone,
		Speed:       30 * time.Millisecond,
		Cycles:      2,
		Width:       20,
		ColorOutput: true,
	}
}

// IsInteractive checks if the program is running in an interactive terminal
func IsInteractive() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

// AnimateOutput displays the output with the specified animation effect
func AnimateOutput(output string, config AnimationConfig) {
	// Skip animation if not enabled or not in interactive mode
	if !config.Enabled || !IsInteractive() {
		if config.ColorOutput {
			PrintGreen(output)
		} else {
			fmt.Println(output)
		}
		return
	}

	// Apply the selected animation effect
	switch config.Type {
	case AnimationTyping:
		animateTyping(output, config)
	case AnimationRainbow:
		animateRainbow(output, config)
	case AnimationBanner:
		animateBanner(output, config)
	case AnimationLoading:
		animateLoading(config)
		if config.ColorOutput {
			PrintGreen(output)
		} else {
			fmt.Println(output)
		}
	default:
		if config.ColorOutput {
			PrintGreen(output)
		} else {
			fmt.Println(output)
		}
	}
}

// animateTyping displays text with a typing animation effect
func animateTyping(text string, config AnimationConfig) {
	lines := strings.Split(text, "\n")
	
	for i, line := range lines {
		for _, char := range line {
			if config.ColorOutput {
				fmt.Print(ColorGreen + string(char) + ColorReset)
			} else {
				fmt.Print(string(char))
			}
			time.Sleep(config.Speed)
		}
		
		// Don't add newline after the last line if it already ends with one
		if i < len(lines)-1 || !strings.HasSuffix(text, "\n") {
			fmt.Println()
		}
	}
}

// animateRainbow displays text with changing colors
func animateRainbow(text string, config AnimationConfig) {
	colorCodes := []string{
		ColorRed,
		ColorYellow,
		ColorGreen,
		ColorCyan,
		ColorBlue,
		ColorPurple,
	}

	lines := strings.Split(text, "\n")
	
	for _, line := range lines {
		for c := 0; c < config.Cycles; c++ {
			fmt.Print("\r")
			for i, char := range line {
				if config.ColorOutput {
					colorIndex := (i + c) % len(colorCodes)
					fmt.Print(colorCodes[colorIndex] + string(char) + ColorReset)
				} else {
					fmt.Print(string(char))
				}
				time.Sleep(config.Speed / 2)
			}
			time.Sleep(config.Speed * 2)
		}
		
		// Final display in green or plain
		fmt.Print("\r")
		if config.ColorOutput {
			fmt.Println(ColorGreen + line + ColorReset)
		} else {
			fmt.Println(line)
		}
	}
}

// animateBanner displays a scrolling banner of text
func animateBanner(text string, config AnimationConfig) {
	// For multi-line text, only animate the first line
	firstLine := strings.Split(text, "\n")[0]
	
	// If the text is shorter than the width, pad it
	if len(firstLine) < config.Width {
		firstLine = firstLine + strings.Repeat(" ", config.Width-len(firstLine))
	}
	
	// Pad the text for scrolling
	paddedText := firstLine + "   " + firstLine
	
	for c := 0; c < config.Cycles; c++ {
		for i := 0; i < len(firstLine)+3; i++ {
			if i+config.Width <= len(paddedText) {
				segment := paddedText[i : i+config.Width]
				if config.ColorOutput {
					fmt.Printf("\r%s%s%s", ColorCyan, segment, ColorReset)
				} else {
					fmt.Printf("\r%s", segment)
				}
				time.Sleep(config.Speed)
			}
		}
	}
	
	fmt.Println()
	
	// Display the rest of the text normally
	restOfText := strings.Join(strings.Split(text, "\n")[1:], "\n")
	if restOfText != "" {
		if config.ColorOutput {
			PrintGreen(restOfText)
		} else {
			fmt.Println(restOfText)
		}
	}
}

// animateLoading displays a loading bar animation
func animateLoading(config AnimationConfig) {
	fmt.Print("Decoding art ")
	width := config.Width
	
	for i := 0; i <= width; i++ {
		progress := float64(i) / float64(width) * 100
		bar := ""
		
		for j := 0; j < width; j++ {
			if j < i {
				bar += "█"
			} else {
				bar += "░"
			}
		}
		
		if config.ColorOutput {
			fmt.Printf("\r%s[%s]%s %.1f%%", ColorYellow, bar, ColorReset, progress)
		} else {
			fmt.Printf("\r[%s] %.1f%%", bar, progress)
		}
		
		time.Sleep(config.Speed)
	}
	
	fmt.Println("\nDone!")
}

// AnimateTitle displays a title with animation effects
func AnimateTitle(title string, config AnimationConfig) {
	if !config.Enabled || !IsInteractive() {
		return
	}
	
	// Always use banner animation for the title
	bannerConfig := config
	bannerConfig.Type = AnimationBanner
	
	animateBanner(title, bannerConfig)
}