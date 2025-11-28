// Package terminal provides utilities for terminal operations.
package terminal

import (
	"os"
	"os/exec"
)

// TermClear clears the terminal screen.
func TermClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
