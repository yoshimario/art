package functions

import (
	"fmt"
	"os/exec"
	"runtime"
)

// SoundType represents the type of sound effect to play
type SoundType string

const (
	// Sound effect types
	SoundBeep    SoundType = "beep"
	SoundSuccess SoundType = "success"
	SoundError   SoundType = "error"
	SoundTyping  SoundType = "typing"
	SoundLoading SoundType = "loading"
	SoundBanner  SoundType = "banner"
	SoundNone    SoundType = "none"
)

// SoundConfig stores configuration for sound effects
type SoundConfig struct {
	Enabled      bool
	Type         SoundType
	Volume       int     // 0-100
	Pitch        float64 // Frequency multiplier: 0.5 = lower, 2.0 = higher
	BpmSpeed     int     // For timing repetitive sounds (beats per minute)
	Duration     int     // Duration in milliseconds
	RepeatCount  int     // Number of times to repeat the sound
	RepeatDelay  int     // Delay between repeats in milliseconds
	WaitForSound bool    // Whether to wait for sound to finish before continuing
}

// NewDefaultSoundConfig creates a default sound configuration
func NewDefaultSoundConfig() SoundConfig {
	return SoundConfig{
		Enabled:      false,
		Type:         SoundNone,
		Volume:       50,
		Pitch:        1.0,
		BpmSpeed:     120,
		Duration:     100,
		RepeatCount:  1,
		RepeatDelay:  500,
		WaitForSound: false,
	}
}

// PlaySound plays the specified sound effect
func PlaySound(config SoundConfig) error {
	if !config.Enabled || !IsInteractive() {
		return nil
	}

	switch config.Type {
	case SoundBeep:
		return playBeepSound(config)
	case SoundSuccess:
		return playSuccessSound(config)
	case SoundError:
		return playErrorSound(config)
	case SoundTyping:
		return playTypingSound(config)
	case SoundLoading:
		return playLoadingSound(config)
	case SoundBanner:
		return playBannerSound(config)
	default:
		return nil
	}
}

// playBeepSound plays a simple beep sound using SoX
func playBeepSound(config SoundConfig) error {
	frequency := int(440 * config.Pitch) // Default A4 (440Hz) adjusted by pitch
	duration := float64(config.Duration) / 1000.0 // Convert milliseconds to seconds

	cmd := exec.Command("sox", "-n", "-t", "wav", "-", "synth", fmt.Sprintf("%0.3f", duration), "sine", fmt.Sprintf("%d", frequency), "vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100))
	return runSoXCommand(cmd, config)
}

// playSuccessSound plays a tornado warble sound for success
func playSuccessSound(config SoundConfig) error {
	// Make the duration longer - triple the configured duration
	duration := float64(config.Duration * 3) / 1000.0 // Convert milliseconds to seconds
	
	// Modify the command to use "play" directly instead of "sox"
	// Enhance the tornado warble effect with more pronounced tremolo parameters
	cmd := exec.Command("play", "-n", "-t", "wav", "synth", 
											fmt.Sprintf("%0.3f", duration), 
											"sine", "440", 
											"tremolo", "12", "60",  // Increased tremolo depth and rate
											"vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100))
	
	// Run the command directly
	if config.WaitForSound {
			return cmd.Run()
	} else {
			return cmd.Start()
	}
}

// playErrorSound plays a falling tone sequence for error
func playErrorSound(config SoundConfig) error {
	duration := float64(config.Duration) / 1000.0 // Convert milliseconds to seconds

	// Generate a falling tone from 880Hz to 440Hz
	cmd := exec.Command("sox", "-n", "-t", "wav", "-", "synth", fmt.Sprintf("%0.3f", duration), "sine", "880-440", "vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100))
	return runSoXCommand(cmd, config)
}

// playTypingSound plays a typing sound effect
func playTypingSound(config SoundConfig) error {
	duration := float64(config.Duration) / 1000.0 // Convert milliseconds to seconds

	// Generate a short, sharp click sound
	cmd := exec.Command("sox", "-n", "-t", "wav", "-", "synth", fmt.Sprintf("%0.3f", duration), "square", "1000", "vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100))
	return runSoXCommand(cmd, config)
}

// playLoadingSound plays a loading progress sound
func playLoadingSound(config SoundConfig) error {
	duration := float64(config.Duration) / 1000.0 // Convert milliseconds to seconds

	// Generate a pulsing sound
	cmd := exec.Command("sox", "-n", "-t", "wav", "-", "synth", fmt.Sprintf("%0.3f", duration), "sine", "440", "vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100), "fade", "q", "0.1", "0", "0.1")
	return runSoXCommand(cmd, config)
}

// playBannerSound plays a banner scrolling sound
func playBannerSound(config SoundConfig) error {
	duration := float64(config.Duration) / 1000.0 // Convert milliseconds to seconds

	// Generate a smooth, sweeping sound
	cmd := exec.Command("sox", "-n", "-t", "wav", "-", "synth", fmt.Sprintf("%0.3f", duration), "sine", "440-660", "vol", fmt.Sprintf("%0.2f", float64(config.Volume)/100))
	return runSoXCommand(cmd, config)
}

// runSoXCommand runs the SoX command and handles waiting or starting the process
func runSoXCommand(cmd *exec.Cmd, config SoundConfig) error {
	// Create a pipe to play the sound
	playCmd := exec.Command("play", "-")
	
	// Connect the output of sox to the input of play
	playCmd.Stdin, _ = cmd.StdoutPipe()
	
	// Start the play command first
	if err := playCmd.Start(); err != nil {
			return err
	}
	
	// Run the sox command
	if err := cmd.Run(); err != nil {
			return err
	}
	
	// Wait for play to finish if configured
	if config.WaitForSound {
			return playCmd.Wait()
	}
	
	return nil
}

// DetectSoundCapability checks if the system can play sounds
func DetectSoundCapability() bool {
	switch runtime.GOOS {
	case "linux":
		// Check for common sound utilities
		for _, cmd := range []string{"sox", "beep", "paplay", "aplay", "play"} {
			if _, err := exec.LookPath(cmd); err == nil {
				return true
			}
		}
	case "darwin":
		// macOS should have afplay
		if _, err := exec.LookPath("afplay"); err == nil {
			return true
		}
	case "windows":
		// Windows can use PowerShell for beeps
		if _, err := exec.LookPath("powershell.exe"); err == nil {
			return true
		}
	}
	return false
}
