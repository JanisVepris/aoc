package renderer

import (
	"os"
	"strings"
)

// BoxDrawRenderer renders data using box-drawing characters for connected line art.
//
// Usage example:
//
//	renderer := renderer.NewBoxDrawRenderer[rune](width)
//	renderer.SetOnValue('#')
//	renderer.Append(data...)
//	renderer.WriteAll()
//
// Key features:
//   - Detects adjacency patterns (top, bottom, left, right neighbors)
//   - Automatically selects appropriate connector: ─│┌┐└┘├┤┬┴┼
//   - Creates connected line art from binary data
//   - Perfect for grids, borders, and structural visualizations
//   - 1×1 mapping (one character per data point)
type BoxDrawRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewBoxDrawRenderer[T comparable](lineWidth int) *BoxDrawRenderer[T] {
	return &BoxDrawRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (b *BoxDrawRenderer[T]) SetData(data []T) {
	b.data = data
	b.cursor = 0
}

func (b *BoxDrawRenderer[T]) SetOutput(output *os.File) {
	b.out = output
}

func (b *BoxDrawRenderer[T]) SetLineWidth(lineWidth int) {
	b.lineWidth = lineWidth
}

func (b *BoxDrawRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	b.onFunc = onFunc
}

func (b *BoxDrawRenderer[T]) SetOnValue(value T) {
	b.onVal = value
}

func (b *BoxDrawRenderer[T]) Clone() *BoxDrawRenderer[T] {
	dataCopy := make([]T, len(b.data))
	copy(dataCopy, b.data)

	return &BoxDrawRenderer[T]{
		out:       b.out,
		cursor:    b.cursor,
		data:      dataCopy,
		onVal:     b.onVal,
		onFunc:    b.onFunc,
		lineWidth: b.lineWidth,
	}
}

func (b *BoxDrawRenderer[T]) Append(values ...T) {
	b.data = append(b.data, values...)
}

func (b *BoxDrawRenderer[T]) isOn(idx int) bool {
	if idx < 0 || idx >= len(b.data) {
		return false
	}
	data := b.data[idx]
	return data == b.onVal || (b.onFunc != nil && b.onFunc(data))
}

func (b *BoxDrawRenderer[T]) RenderAtCursor() string {
	if b.cursor >= len(b.data) {
		return " "
	}

	if !b.isOn(b.cursor) {
		return " "
	}

	col := b.cursor % b.lineWidth

	top := b.isOn(b.cursor - b.lineWidth)
	bottom := b.isOn(b.cursor + b.lineWidth)
	left := col > 0 && b.isOn(b.cursor-1)
	right := col < b.lineWidth-1 && b.isOn(b.cursor+1)

	return boxDrawChar(top, bottom, left, right)
}

func boxDrawChar(top, bottom, left, right bool) string {
	switch {
	case top && bottom && left && right:
		return "┼"
	case top && bottom && left && !right:
		return "┤"
	case top && bottom && !left && right:
		return "├"
	case top && bottom && !left && !right:
		return "│"
	case top && !bottom && left && right:
		return "┴"
	case top && !bottom && left && !right:
		return "┘"
	case top && !bottom && !left && right:
		return "└"
	case top && !bottom && !left && !right:
		return "╵"
	case !top && bottom && left && right:
		return "┬"
	case !top && bottom && left && !right:
		return "┐"
	case !top && bottom && !left && right:
		return "┌"
	case !top && bottom && !left && !right:
		return "╷"
	case !top && !bottom && left && right:
		return "─"
	case !top && !bottom && left && !right:
		return "╴"
	case !top && !bottom && !left && right:
		return "╶"
	default:
		return "█"
	}
}

func (b *BoxDrawRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := b.cursor
	b.cursor = 0

	for b.cursor < len(b.data) {
		sb.WriteString(b.RenderAtCursor())

		b.Advance()

		if b.cursor%b.lineWidth == 0 && b.cursor > 0 {
			sb.WriteString("\n")
		}
	}

	b.cursor = oldCursor
	return sb.String()
}

func (b *BoxDrawRenderer[T]) WriteSingle() bool {
	b.out.WriteString(b.RenderAtCursor())

	oldCursor := b.cursor
	b.Advance()

	return b.cursor%b.lineWidth == 0 && oldCursor%b.lineWidth != 0
}

func (b *BoxDrawRenderer[T]) WriteLine() bool {
	if b.cursor >= len(b.data) {
		return false
	}

	for b.cursor < len(b.data) {
		b.out.WriteString(b.RenderAtCursor())

		oldCursor := b.cursor
		b.Advance()

		if b.cursor%b.lineWidth == 0 && oldCursor%b.lineWidth != 0 {
			b.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (b *BoxDrawRenderer[T]) WriteAll() {
	b.cursor = 0

	for b.cursor < len(b.data) {
		b.out.WriteString(b.RenderAtCursor())

		b.Advance()

		if b.cursor%b.lineWidth == 0 && b.cursor > 0 {
			b.out.WriteString("\n")
		}
	}
}

func (b *BoxDrawRenderer[T]) Clear() {
	b.cursor = 0
	b.data = []T{}
}

func (b *BoxDrawRenderer[T]) ResetCursor() {
	b.cursor = 0
}

func (b *BoxDrawRenderer[T]) HasMore() bool {
	return b.cursor < len(b.data)
}

func (b *BoxDrawRenderer[T]) Len() int {
	return len(b.data)
}

func (b *BoxDrawRenderer[T]) Progress() (cursor, total int) {
	return b.cursor, len(b.data)
}

func (b *BoxDrawRenderer[T]) Advance() {
	b.cursor++
}
