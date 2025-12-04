package renderer

import (
	"os"
	"strings"
)

// ASCIIDitherRenderer renders data using ASCII characters weighted by visual density
// for classic ASCII art style output.
//
// Usage example:
//
//	renderer := renderer.NewASCIIDitherRenderer[rune](width)
//	renderer.SetOnValue('@')
//	renderer.Append(data...)
//	renderer.WriteAll()
//
// Key features:
//   - Uses density characters: ' .':;~-"^+*!#%@
//   - Off pixels = lightest character (space)
//   - On pixels = heaviest character (@)
//   - Classic retro ASCII art aesthetic
//   - Universal terminal support
//   - 1×1 mapping (one character per data point)
type ASCIIDitherRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

var asciiDensityChars = []rune{' ', '.', '\'', ':', ';', '~', '-', '"', '^', '+', '*', '!', '#', '%', '@'}

func NewASCIIDitherRenderer[T comparable](lineWidth int) *ASCIIDitherRenderer[T] {
	return &ASCIIDitherRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (a *ASCIIDitherRenderer[T]) SetData(data []T) {
	a.data = data
	a.cursor = 0
}

func (a *ASCIIDitherRenderer[T]) SetOutput(output *os.File) {
	a.out = output
}

func (a *ASCIIDitherRenderer[T]) SetLineWidth(lineWidth int) {
	a.lineWidth = lineWidth
}

func (a *ASCIIDitherRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	a.onFunc = onFunc
}

func (a *ASCIIDitherRenderer[T]) SetOnValue(value T) {
	a.onVal = value
}

func (a *ASCIIDitherRenderer[T]) Clone() *ASCIIDitherRenderer[T] {
	dataCopy := make([]T, len(a.data))
	copy(dataCopy, a.data)

	return &ASCIIDitherRenderer[T]{
		out:       a.out,
		cursor:    a.cursor,
		data:      dataCopy,
		onVal:     a.onVal,
		onFunc:    a.onFunc,
		lineWidth: a.lineWidth,
	}
}

func (a *ASCIIDitherRenderer[T]) Append(values ...T) {
	a.data = append(a.data, values...)
}

func (a *ASCIIDitherRenderer[T]) RenderAtCursor() string {
	if a.cursor >= len(a.data) {
		return " "
	}

	data := a.data[a.cursor]
	isOn := data == a.onVal || (a.onFunc != nil && a.onFunc(data))

	if !isOn {
		return string(asciiDensityChars[0])
	}

	charIdx := len(asciiDensityChars) - 1
	return string(asciiDensityChars[charIdx])
}

func (a *ASCIIDitherRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := a.cursor
	a.cursor = 0

	for a.cursor < len(a.data) {
		sb.WriteString(a.RenderAtCursor())

		a.Advance()

		if a.cursor%a.lineWidth == 0 && a.cursor > 0 {
			sb.WriteString("\n")
		}
	}

	a.cursor = oldCursor
	return sb.String()
}

func (a *ASCIIDitherRenderer[T]) WriteSingle() bool {
	a.out.WriteString(a.RenderAtCursor())

	oldCursor := a.cursor
	a.Advance()

	return a.cursor%a.lineWidth == 0 && oldCursor%a.lineWidth != 0
}

func (a *ASCIIDitherRenderer[T]) WriteLine() bool {
	if a.cursor >= len(a.data) {
		return false
	}

	for a.cursor < len(a.data) {
		a.out.WriteString(a.RenderAtCursor())

		oldCursor := a.cursor
		a.Advance()

		if a.cursor%a.lineWidth == 0 && oldCursor%a.lineWidth != 0 {
			a.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (a *ASCIIDitherRenderer[T]) WriteAll() {
	a.cursor = 0

	for a.cursor < len(a.data) {
		a.out.WriteString(a.RenderAtCursor())

		a.Advance()

		if a.cursor%a.lineWidth == 0 && a.cursor > 0 {
			a.out.WriteString("\n")
		}
	}
}

func (a *ASCIIDitherRenderer[T]) Clear() {
	a.cursor = 0
	a.data = []T{}
}

func (a *ASCIIDitherRenderer[T]) ResetCursor() {
	a.cursor = 0
}

func (a *ASCIIDitherRenderer[T]) HasMore() bool {
	return a.cursor < len(a.data)
}

func (a *ASCIIDitherRenderer[T]) Len() int {
	return len(a.data)
}

func (a *ASCIIDitherRenderer[T]) Progress() (cursor, total int) {
	return a.cursor, len(a.data)
}

func (a *ASCIIDitherRenderer[T]) Advance() {
	a.cursor++
}
