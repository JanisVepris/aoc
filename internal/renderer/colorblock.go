package renderer

import (
	"fmt"
	"os"
	"strings"
)

type ColorBlockRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
	colorFunc func(T) int
}

func NewColorBlockRenderer[T comparable](lineWidth int) *ColorBlockRenderer[T] {
	return &ColorBlockRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
		colorFunc: defaultColorBlockFunc[T](),
	}
}

func defaultColorBlockFunc[T comparable]() func(T) int {
	return func(val T) int {
		return 15
	}
}

func (c *ColorBlockRenderer[T]) SetData(data []T) {
	c.data = data
	c.cursor = 0
}

func (c *ColorBlockRenderer[T]) SetOutput(output *os.File) {
	c.out = output
}

func (c *ColorBlockRenderer[T]) SetLineWidth(lineWidth int) {
	c.lineWidth = lineWidth
}

func (c *ColorBlockRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	c.onFunc = onFunc
}

func (c *ColorBlockRenderer[T]) SetOnValue(value T) {
	c.onVal = value
}

func (c *ColorBlockRenderer[T]) SetColorFunc(colorFunc func(T) int) {
	c.colorFunc = colorFunc
}

func (c *ColorBlockRenderer[T]) Clone() *ColorBlockRenderer[T] {
	dataCopy := make([]T, len(c.data))
	copy(dataCopy, c.data)

	return &ColorBlockRenderer[T]{
		out:       c.out,
		cursor:    c.cursor,
		data:      dataCopy,
		onVal:     c.onVal,
		onFunc:    c.onFunc,
		lineWidth: c.lineWidth,
		colorFunc: c.colorFunc,
	}
}

func (c *ColorBlockRenderer[T]) Append(values ...T) {
	c.data = append(c.data, values...)
}

func (c *ColorBlockRenderer[T]) RenderAtCursor() string {
	tlIdx := c.cursor
	trIdx := c.cursor + 1
	blIdx := c.cursor + c.lineWidth
	brIdx := c.cursor + c.lineWidth + 1

	tl := false
	tr := false
	bl := false
	br := false
	color := 15

	cursorCol := c.cursor % c.lineWidth

	if tlIdx < len(c.data) {
		data := c.data[tlIdx]
		tl = data == c.onVal || (c.onFunc != nil && c.onFunc(data))
		if tl {
			color = c.colorFunc(data)
		}
	}

	if trIdx < len(c.data) && cursorCol < c.lineWidth-1 {
		data := c.data[trIdx]
		tr = data == c.onVal || (c.onFunc != nil && c.onFunc(data))
	}

	if blIdx < len(c.data) {
		data := c.data[blIdx]
		bl = data == c.onVal || (c.onFunc != nil && c.onFunc(data))
	}

	if brIdx < len(c.data) && cursorCol < c.lineWidth-1 {
		data := c.data[brIdx]
		br = data == c.onVal || (c.onFunc != nil && c.onFunc(data))
	}

	char := quadrantChar(tl, tr, bl, br)
	if char == " " {
		return " "
	}

	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, char)
}

func (c *ColorBlockRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := c.cursor
	c.cursor = 0

	for c.cursor < len(c.data) {
		sb.WriteString(c.RenderAtCursor())

		oldPos := c.cursor
		c.Advance()

		if c.cursor-oldPos > 2 {
			sb.WriteString("\n")
		}
	}

	c.cursor = oldCursor
	return sb.String()
}

func (c *ColorBlockRenderer[T]) WriteSingle() bool {
	c.out.WriteString(c.RenderAtCursor())

	oldCursor := c.cursor
	c.Advance()

	return c.cursor-oldCursor > 2
}

func (c *ColorBlockRenderer[T]) WriteLine() bool {
	if c.cursor >= len(c.data) {
		return false
	}

	for c.cursor < len(c.data) {
		c.out.WriteString(c.RenderAtCursor())

		oldCursor := c.cursor
		c.Advance()

		if c.cursor-oldCursor > 2 {
			c.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (c *ColorBlockRenderer[T]) WriteAll() {
	c.cursor = 0

	for c.cursor < len(c.data) {
		c.out.WriteString(c.RenderAtCursor())

		oldCursor := c.cursor
		c.Advance()

		if c.cursor-oldCursor > 2 {
			c.out.WriteString("\n")
		}
	}
}

func (c *ColorBlockRenderer[T]) Clear() {
	c.cursor = 0
	c.data = []T{}
}

func (c *ColorBlockRenderer[T]) ResetCursor() {
	c.cursor = 0
}

func (c *ColorBlockRenderer[T]) HasMore() bool {
	return c.cursor < len(c.data)
}

func (c *ColorBlockRenderer[T]) Len() int {
	return len(c.data)
}

func (c *ColorBlockRenderer[T]) Progress() (cursor, total int) {
	return c.cursor, len(c.data)
}

func (c *ColorBlockRenderer[T]) Advance() {
	c.cursor += 2
	switch c.cursor % c.lineWidth {
	case 0:
		c.cursor += c.lineWidth
	case 1:
		c.cursor += c.lineWidth - 1
	}
}
