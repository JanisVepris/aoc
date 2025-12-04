package renderer

import (
	"os"
	"strings"
)

type BrailleRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewBrailleRenderer[T comparable](lineWidth int) *BrailleRenderer[T] {
	return &BrailleRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (b *BrailleRenderer[T]) SetData(data []T) {
	b.data = data
	b.cursor = 0
}

func (b *BrailleRenderer[T]) SetOutput(output *os.File) {
	b.out = output
}

// SetLineWidth sets the width of each line in the braille rendering.
func (b *BrailleRenderer[T]) SetLineWidth(lineWidth int) {
	b.lineWidth = lineWidth
}

// SetOnFunc sets the function that determines whether a value is considered "on" in the braille rendering.
func (b *BrailleRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	b.onFunc = onFunc
}

// SetOnValue sets the value that will be considered "on" in the braille rendering. This takes precedence over any existing onFunc.
func (b *BrailleRenderer[T]) SetOnValue(value T) {
	b.onVal = value
}

// Clone creates a copy of the current BrailleRenderer with the same settings and data.
func (b *BrailleRenderer[T]) Clone() *BrailleRenderer[T] {
	dataCopy := make([]T, len(b.data))
	copy(dataCopy, b.data)

	return &BrailleRenderer[T]{
		out:       b.out,
		cursor:    b.cursor,
		data:      dataCopy,
		onVal:     b.onVal,
		onFunc:    b.onFunc,
		lineWidth: b.lineWidth,
	}
}

// Append adds values to the braille renderer's data.
func (b *BrailleRenderer[T]) Append(values ...T) {
	b.data = append(b.data, values...)
}

// RenderAtCursor renders a braille character at the current cursor position to a string
func (b *BrailleRenderer[T]) RenderAtCursor() string {
	indexes := [8]int{
		b.cursor,
		b.cursor + b.lineWidth,
		b.cursor + b.lineWidth*2,
		b.cursor + 1,
		b.cursor + b.lineWidth + 1,
		b.cursor + b.lineWidth*2 + 1,
		b.cursor + b.lineWidth*3,
		b.cursor + b.lineWidth*3 + 1,
	}

	dots := 0
	cursorCol := b.cursor % b.lineWidth

	for dotIdx, dataIdx := range indexes {
		if dataIdx >= len(b.data) {
			continue
		}

		if dotIdx >= 3 {
			dataCol := dataIdx % b.lineWidth
			if dataCol == 0 && cursorCol == b.lineWidth-1 {
				continue
			}
		}

		data := b.data[dataIdx]
		if data == b.onVal || (b.onFunc != nil && b.onFunc(data)) {
			dots |= 1 << dotIdx
		}
	}

	return string(rune(0x2800 + dots))
}

// RenderToString renders all data to a string
func (b *BrailleRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := b.cursor
	b.cursor = 0

	for b.cursor < len(b.data) {
		sb.WriteString(b.RenderAtCursor())

		oldPos := b.cursor
		b.Advance()

		if b.cursor-oldPos > 2 {
			sb.WriteString("\n")
		}
	}

	b.cursor = oldCursor
	return sb.String()
}

// WriteSingle renders a single braille character into output, returns true if a full line was completed
func (b *BrailleRenderer[T]) WriteSingle() bool {
	b.out.WriteString(b.RenderAtCursor())

	oldCursor := b.cursor
	b.Advance()

	return b.cursor-oldCursor > 2
}

// WriteLine renders a single line into output
func (b *BrailleRenderer[T]) WriteLine() bool {
	if b.cursor >= len(b.data) {
		return false
	}

	for b.cursor < len(b.data) {
		b.out.WriteString(b.RenderAtCursor())

		oldCursor := b.cursor
		b.Advance()

		if b.cursor-oldCursor > 2 {
			b.out.WriteString("\n")
			return true
		}
	}

	return false
}

// WriteAll renders all into output
func (b *BrailleRenderer[T]) WriteAll() {
	b.cursor = 0

	for b.cursor < len(b.data) {
		b.out.WriteString(b.RenderAtCursor())

		oldCursor := b.cursor
		b.Advance()

		if b.cursor-oldCursor > 2 {
			b.out.WriteString("\n")
		}
	}
}

func (b *BrailleRenderer[T]) Clear() {
	b.cursor = 0
	b.data = []T{}
}

func (b *BrailleRenderer[T]) ResetCursor() {
	b.cursor = 0
}

func (b *BrailleRenderer[T]) HasMore() bool {
	return b.cursor < len(b.data)
}

func (b *BrailleRenderer[T]) Len() int {
	return len(b.data)
}

func (b *BrailleRenderer[T]) Progress() (cursor, total int) {
	return b.cursor, len(b.data)
}

// Advance moves the cursor to the next braille character position
func (b *BrailleRenderer[T]) Advance() {
	b.cursor += 2
	switch b.cursor % b.lineWidth {
	case 0:
		b.cursor += b.lineWidth * 3
	case 1:
		b.cursor += (b.lineWidth * 3) - 1
	}
}
