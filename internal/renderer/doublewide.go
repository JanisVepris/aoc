package renderer

import (
	"os"
	"strings"
)

type DoubleWideRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewDoubleWideRenderer[T comparable](lineWidth int) *DoubleWideRenderer[T] {
	return &DoubleWideRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (d *DoubleWideRenderer[T]) SetData(data []T) {
	d.data = data
	d.cursor = 0
}

func (d *DoubleWideRenderer[T]) SetOutput(output *os.File) {
	d.out = output
}

func (d *DoubleWideRenderer[T]) SetLineWidth(lineWidth int) {
	d.lineWidth = lineWidth
}

func (d *DoubleWideRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	d.onFunc = onFunc
}

func (d *DoubleWideRenderer[T]) SetOnValue(value T) {
	d.onVal = value
}

func (d *DoubleWideRenderer[T]) Clone() *DoubleWideRenderer[T] {
	dataCopy := make([]T, len(d.data))
	copy(dataCopy, d.data)

	return &DoubleWideRenderer[T]{
		out:       d.out,
		cursor:    d.cursor,
		data:      dataCopy,
		onVal:     d.onVal,
		onFunc:    d.onFunc,
		lineWidth: d.lineWidth,
	}
}

func (d *DoubleWideRenderer[T]) Append(values ...T) {
	d.data = append(d.data, values...)
}

func (d *DoubleWideRenderer[T]) RenderAtCursor() string {
	if d.cursor >= len(d.data) {
		return "　"
	}

	data := d.data[d.cursor]
	isOn := data == d.onVal || (d.onFunc != nil && d.onFunc(data))

	if !isOn {
		return "　"
	}

	return "　"
}

func (d *DoubleWideRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := d.cursor
	d.cursor = 0

	for d.cursor < len(d.data) {
		sb.WriteString(d.RenderAtCursor())

		d.Advance()

		if d.cursor%d.lineWidth == 0 && d.cursor > 0 {
			sb.WriteString("\n")
		}
	}

	d.cursor = oldCursor
	return sb.String()
}

func (d *DoubleWideRenderer[T]) WriteSingle() bool {
	d.out.WriteString(d.RenderAtCursor())

	oldCursor := d.cursor
	d.Advance()

	return d.cursor%d.lineWidth == 0 && oldCursor%d.lineWidth != 0
}

func (d *DoubleWideRenderer[T]) WriteLine() bool {
	if d.cursor >= len(d.data) {
		return false
	}

	for d.cursor < len(d.data) {
		d.out.WriteString(d.RenderAtCursor())

		oldCursor := d.cursor
		d.Advance()

		if d.cursor%d.lineWidth == 0 && oldCursor%d.lineWidth != 0 {
			d.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (d *DoubleWideRenderer[T]) WriteAll() {
	d.cursor = 0

	for d.cursor < len(d.data) {
		d.out.WriteString(d.RenderAtCursor())

		d.Advance()

		if d.cursor%d.lineWidth == 0 && d.cursor > 0 {
			d.out.WriteString("\n")
		}
	}
}

func (d *DoubleWideRenderer[T]) Clear() {
	d.cursor = 0
	d.data = []T{}
}

func (d *DoubleWideRenderer[T]) ResetCursor() {
	d.cursor = 0
}

func (d *DoubleWideRenderer[T]) HasMore() bool {
	return d.cursor < len(d.data)
}

func (d *DoubleWideRenderer[T]) Len() int {
	return len(d.data)
}

func (d *DoubleWideRenderer[T]) Progress() (cursor, total int) {
	return d.cursor, len(d.data)
}

func (d *DoubleWideRenderer[T]) Advance() {
	d.cursor++
}
