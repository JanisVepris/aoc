package renderer

import (
	"fmt"
	"os"
	"strings"
)

// ColorRenderer renders data using ANSI 256-color codes for beautiful heatmap-style visualizations.
//
// Usage example:
//
//	renderer := renderer.NewColorRenderer[rune](width)
//	renderer.SetOnValue('@')
//	renderer.SetColorFunc(func(val rune) int {
//	    if val == '@' {
//	        return 196  // Red
//	    }
//	    return 21  // Blue
//	})
//	renderer.Append(data...)
//	renderer.WriteAll()
//
// Key features:
//   - SetColorFunc(func(T) int) - Map data values to ANSI 256 colors (0-255)
//   - Uses full block █ with colors
//   - 1×1 mapping (one character per data point)
//   - Excellent terminal support
type ColorRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
	colorFunc func(T) int
}

func NewColorRenderer[T comparable](lineWidth int) *ColorRenderer[T] {
	return &ColorRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
		colorFunc: defaultColorFunc[T](),
	}
}

func defaultColorFunc[T comparable]() func(T) int {
	return func(val T) int {
		return 196
	}
}

func (c *ColorRenderer[T]) SetData(data []T) {
	c.data = data
	c.cursor = 0
}

func (c *ColorRenderer[T]) SetOutput(output *os.File) {
	c.out = output
}

func (c *ColorRenderer[T]) SetLineWidth(lineWidth int) {
	c.lineWidth = lineWidth
}

func (c *ColorRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	c.onFunc = onFunc
}

func (c *ColorRenderer[T]) SetOnValue(value T) {
	c.onVal = value
}

func (c *ColorRenderer[T]) SetColorFunc(colorFunc func(T) int) {
	c.colorFunc = colorFunc
}

func (c *ColorRenderer[T]) Clone() *ColorRenderer[T] {
	dataCopy := make([]T, len(c.data))
	copy(dataCopy, c.data)

	return &ColorRenderer[T]{
		out:       c.out,
		cursor:    c.cursor,
		data:      dataCopy,
		onVal:     c.onVal,
		onFunc:    c.onFunc,
		lineWidth: c.lineWidth,
		colorFunc: c.colorFunc,
	}
}

func (c *ColorRenderer[T]) Append(values ...T) {
	c.data = append(c.data, values...)
}

func (c *ColorRenderer[T]) RenderAtCursor() string {
	if c.cursor >= len(c.data) {
		return " "
	}

	data := c.data[c.cursor]
	isOn := data == c.onVal || (c.onFunc != nil && c.onFunc(data))

	if !isOn {
		return " "
	}

	color := c.colorFunc(data)
	return fmt.Sprintf("\033[38;5;%dm█\033[0m", color)
}

func (c *ColorRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := c.cursor
	c.cursor = 0

	for c.cursor < len(c.data) {
		sb.WriteString(c.RenderAtCursor())

		c.Advance()

		if c.cursor%c.lineWidth == 0 && c.cursor > 0 {
			sb.WriteString("\n")
		}
	}

	c.cursor = oldCursor
	return sb.String()
}

func (c *ColorRenderer[T]) WriteSingle() bool {
	c.out.WriteString(c.RenderAtCursor())

	oldCursor := c.cursor
	c.Advance()

	return c.cursor%c.lineWidth == 0 && oldCursor%c.lineWidth != 0
}

func (c *ColorRenderer[T]) WriteLine() bool {
	if c.cursor >= len(c.data) {
		return false
	}

	for c.cursor < len(c.data) {
		c.out.WriteString(c.RenderAtCursor())

		oldCursor := c.cursor
		c.Advance()

		if c.cursor%c.lineWidth == 0 && oldCursor%c.lineWidth != 0 {
			c.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (c *ColorRenderer[T]) WriteAll() {
	c.cursor = 0

	for c.cursor < len(c.data) {
		c.out.WriteString(c.RenderAtCursor())

		c.Advance()

		if c.cursor%c.lineWidth == 0 && c.cursor > 0 {
			c.out.WriteString("\n")
		}
	}
}

func (c *ColorRenderer[T]) Clear() {
	c.cursor = 0
	c.data = []T{}
}

func (c *ColorRenderer[T]) ResetCursor() {
	c.cursor = 0
}

func (c *ColorRenderer[T]) HasMore() bool {
	return c.cursor < len(c.data)
}

func (c *ColorRenderer[T]) Len() int {
	return len(c.data)
}

func (c *ColorRenderer[T]) Progress() (cursor, total int) {
	return c.cursor, len(c.data)
}

func (c *ColorRenderer[T]) Advance() {
	c.cursor++
}
