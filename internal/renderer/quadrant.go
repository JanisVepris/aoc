package renderer

import (
	"os"
	"strings"
)

type QuadrantRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewQuadrantRenderer[T comparable](lineWidth int) *QuadrantRenderer[T] {
	return &QuadrantRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (q *QuadrantRenderer[T]) SetData(data []T) {
	q.data = data
	q.cursor = 0
}

func (q *QuadrantRenderer[T]) SetOutput(output *os.File) {
	q.out = output
}

func (q *QuadrantRenderer[T]) SetLineWidth(lineWidth int) {
	q.lineWidth = lineWidth
}

func (q *QuadrantRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	q.onFunc = onFunc
}

func (q *QuadrantRenderer[T]) SetOnValue(value T) {
	q.onVal = value
}

func (q *QuadrantRenderer[T]) Clone() *QuadrantRenderer[T] {
	dataCopy := make([]T, len(q.data))
	copy(dataCopy, q.data)

	return &QuadrantRenderer[T]{
		out:       q.out,
		cursor:    q.cursor,
		data:      dataCopy,
		onVal:     q.onVal,
		onFunc:    q.onFunc,
		lineWidth: q.lineWidth,
	}
}

func (q *QuadrantRenderer[T]) Append(values ...T) {
	q.data = append(q.data, values...)
}

func (q *QuadrantRenderer[T]) RenderAtCursor() string {
	tlIdx := q.cursor
	trIdx := q.cursor + 1
	blIdx := q.cursor + q.lineWidth
	brIdx := q.cursor + q.lineWidth + 1

	tl := false
	tr := false
	bl := false
	br := false

	cursorCol := q.cursor % q.lineWidth

	if tlIdx < len(q.data) {
		data := q.data[tlIdx]
		tl = data == q.onVal || (q.onFunc != nil && q.onFunc(data))
	}

	if trIdx < len(q.data) && cursorCol < q.lineWidth-1 {
		data := q.data[trIdx]
		tr = data == q.onVal || (q.onFunc != nil && q.onFunc(data))
	}

	if blIdx < len(q.data) {
		data := q.data[blIdx]
		bl = data == q.onVal || (q.onFunc != nil && q.onFunc(data))
	}

	if brIdx < len(q.data) && cursorCol < q.lineWidth-1 {
		data := q.data[brIdx]
		br = data == q.onVal || (q.onFunc != nil && q.onFunc(data))
	}

	return quadrantChar(tl, tr, bl, br)
}

func quadrantChar(tl, tr, bl, br bool) string {
	switch {
	case tl && tr && bl && br:
		return "█"
	case tl && tr && bl && !br:
		return "▛"
	case tl && tr && !bl && br:
		return "▜"
	case tl && tr && !bl && !br:
		return "▀"
	case tl && !tr && bl && br:
		return "▙"
	case tl && !tr && bl && !br:
		return "▌"
	case tl && !tr && !bl && br:
		return "▚"
	case tl && !tr && !bl && !br:
		return "▘"
	case !tl && tr && bl && br:
		return "▟"
	case !tl && tr && bl && !br:
		return "▞"
	case !tl && tr && !bl && br:
		return "▐"
	case !tl && tr && !bl && !br:
		return "▝"
	case !tl && !tr && bl && br:
		return "▄"
	case !tl && !tr && bl && !br:
		return "▖"
	case !tl && !tr && !bl && br:
		return "▗"
	default:
		return " "
	}
}

func (q *QuadrantRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := q.cursor
	q.cursor = 0

	for q.cursor < len(q.data) {
		sb.WriteString(q.RenderAtCursor())

		oldPos := q.cursor
		q.Advance()

		if q.cursor-oldPos > 2 {
			sb.WriteString("\n")
		}
	}

	q.cursor = oldCursor
	return sb.String()
}

func (q *QuadrantRenderer[T]) WriteSingle() bool {
	q.out.WriteString(q.RenderAtCursor())

	oldCursor := q.cursor
	q.Advance()

	return q.cursor-oldCursor > 2
}

func (q *QuadrantRenderer[T]) WriteLine() bool {
	if q.cursor >= len(q.data) {
		return false
	}

	for q.cursor < len(q.data) {
		q.out.WriteString(q.RenderAtCursor())

		oldCursor := q.cursor
		q.Advance()

		if q.cursor-oldCursor > 2 {
			q.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (q *QuadrantRenderer[T]) WriteAll() {
	q.cursor = 0

	for q.cursor < len(q.data) {
		q.out.WriteString(q.RenderAtCursor())

		oldCursor := q.cursor
		q.Advance()

		if q.cursor-oldCursor > 2 {
			q.out.WriteString("\n")
		}
	}
}

func (q *QuadrantRenderer[T]) Clear() {
	q.cursor = 0
	q.data = []T{}
}

func (q *QuadrantRenderer[T]) ResetCursor() {
	q.cursor = 0
}

func (q *QuadrantRenderer[T]) HasMore() bool {
	return q.cursor < len(q.data)
}

func (q *QuadrantRenderer[T]) Len() int {
	return len(q.data)
}

func (q *QuadrantRenderer[T]) Progress() (cursor, total int) {
	return q.cursor, len(q.data)
}

func (q *QuadrantRenderer[T]) Advance() {
	q.cursor += 2
	switch q.cursor % q.lineWidth {
	case 0:
		q.cursor += q.lineWidth
	case 1:
		q.cursor += q.lineWidth - 1
	}
}
