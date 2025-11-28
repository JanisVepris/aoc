package terminal

import "fmt"

type Color string

const (
	Black   Color = "30"
	Red     Color = "31"
	Green   Color = "32"
	Yellow  Color = "33"
	Blue    Color = "34"
	Magenta Color = "35"
	Cyan    Color = "36"
	White   Color = "37"
)

func TermColorize(text string, color Color) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", color, text)
}
