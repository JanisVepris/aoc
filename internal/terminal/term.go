// Package terminal provides utilities for terminal operations.
package terminal

import "fmt"

// Clear clears the terminal screen.
func Clear() {
	fmt.Print("\033[2J")
}

func CursorReset() {
	fmt.Print("\033[H")
}

func CursorHide() {
	fmt.Print("\033[?25l")
}

func CursorShow() {
	fmt.Print("\033[?25h")
}
