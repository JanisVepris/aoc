package renderer

import (
	"os"
	"strings"
)

type HalfblockRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewHalfblockRenderer[T comparable](lineWidth int) *HalfblockRenderer[T] {
	return &HalfblockRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (h *HalfblockRenderer[T]) SetData(data []T) {
	h.data = data
	h.cursor = 0
}

func (h *HalfblockRenderer[T]) SetOutput(output *os.File) {
	h.out = output
}

func (h *HalfblockRenderer[T]) SetLineWidth(lineWidth int) {
	h.lineWidth = lineWidth
}

func (h *HalfblockRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	h.onFunc = onFunc
}

func (h *HalfblockRenderer[T]) SetOnValue(value T) {
	h.onVal = value
}

func (h *HalfblockRenderer[T]) Clone() *HalfblockRenderer[T] {
	dataCopy := make([]T, len(h.data))
	copy(dataCopy, h.data)

	return &HalfblockRenderer[T]{
		out:       h.out,
		cursor:    h.cursor,
		data:      dataCopy,
		onVal:     h.onVal,
		onFunc:    h.onFunc,
		lineWidth: h.lineWidth,
	}
}

func (h *HalfblockRenderer[T]) Append(values ...T) {
	h.data = append(h.data, values...)
}

func (h *HalfblockRenderer[T]) RenderAtCursor() string {
	topIdx := h.cursor
	bottomIdx := h.cursor + h.lineWidth

	topOn := false
	bottomOn := false

	if topIdx < len(h.data) {
		data := h.data[topIdx]
		topOn = data == h.onVal || (h.onFunc != nil && h.onFunc(data))
	}

	if bottomIdx < len(h.data) {
		data := h.data[bottomIdx]
		bottomOn = data == h.onVal || (h.onFunc != nil && h.onFunc(data))
	}

	if topOn && bottomOn {
		return "█"
	} else if topOn {
		return "▀"
	} else if bottomOn {
		return "▄"
	}
	return " "
}

func (h *HalfblockRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := h.cursor
	h.cursor = 0

	for h.cursor < len(h.data) {
		sb.WriteString(h.RenderAtCursor())

		oldPos := h.cursor
		h.Advance()

		if h.cursor-oldPos > 1 {
			sb.WriteString("\n")
		}
	}

	h.cursor = oldCursor
	return sb.String()
}

func (h *HalfblockRenderer[T]) WriteSingle() bool {
	h.out.WriteString(h.RenderAtCursor())

	oldCursor := h.cursor
	h.Advance()

	return h.cursor-oldCursor > 1
}

func (h *HalfblockRenderer[T]) WriteLine() bool {
	if h.cursor >= len(h.data) {
		return false
	}

	for h.cursor < len(h.data) {
		h.out.WriteString(h.RenderAtCursor())

		oldCursor := h.cursor
		h.Advance()

		if h.cursor-oldCursor > 1 {
			h.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (h *HalfblockRenderer[T]) WriteAll() {
	h.cursor = 0

	for h.cursor < len(h.data) {
		h.out.WriteString(h.RenderAtCursor())

		oldCursor := h.cursor
		h.Advance()

		if h.cursor-oldCursor > 1 {
			h.out.WriteString("\n")
		}
	}
}

func (h *HalfblockRenderer[T]) Clear() {
	h.cursor = 0
	h.data = []T{}
}

func (h *HalfblockRenderer[T]) ResetCursor() {
	h.cursor = 0
}

func (h *HalfblockRenderer[T]) HasMore() bool {
	return h.cursor < len(h.data)
}

func (h *HalfblockRenderer[T]) Len() int {
	return len(h.data)
}

func (h *HalfblockRenderer[T]) Progress() (cursor, total int) {
	return h.cursor, len(h.data)
}

func (h *HalfblockRenderer[T]) Advance() {
	h.cursor++
	if h.cursor%h.lineWidth == 0 {
		h.cursor += h.lineWidth
	}
}
