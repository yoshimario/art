package functions

import (
	"fmt"
)

// Define color codes
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorReset  = "\033[0m"
)

// PrintRed prints text in red color
func PrintRed(text string) {
	fmt.Println(ColorRed + text + ColorReset)
}

// PrintGreen prints text in green color
func PrintGreen(text string) {
	fmt.Println(ColorGreen + text + ColorReset)
}

// PrintYellow prints text in yellow color
func PrintYellow(text string) {
	fmt.Println(ColorYellow + text + ColorReset)
}

// PrintBlue prints text in blue color
func PrintBlue(text string) {
	fmt.Println(ColorBlue + text + ColorReset)
}

// PrintPurple prints text in purple color
func PrintPurple(text string) {
	fmt.Println(ColorPurple + text + ColorReset)
}

// PrintCyan prints text in cyan color
func PrintCyan(text string) {
	fmt.Println(ColorCyan + text + ColorReset)
}

// PrintWhite prints text in white color
func PrintWhite(text string) {
	fmt.Println(ColorWhite + text + ColorReset)
}

// Colorize returns a colored string
func Colorize(colorCode, text string) string {
	return colorCode + text + ColorReset
}