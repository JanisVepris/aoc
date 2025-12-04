package renderer

import (
	"os"
	"strings"
)

type SextantRenderer[T comparable] struct {
	out       *os.File
	cursor    int
	data      []T
	onVal     T
	onFunc    func(value T) bool
	lineWidth int
}

func NewSextantRenderer[T comparable](lineWidth int) *SextantRenderer[T] {
	return &SextantRenderer[T]{
		lineWidth: lineWidth,
		cursor:    0,
		out:       os.Stdout,
	}
}

func (s *SextantRenderer[T]) SetData(data []T) {
	s.data = data
	s.cursor = 0
}

func (s *SextantRenderer[T]) SetOutput(output *os.File) {
	s.out = output
}

func (s *SextantRenderer[T]) SetLineWidth(lineWidth int) {
	s.lineWidth = lineWidth
}

func (s *SextantRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	s.onFunc = onFunc
}

func (s *SextantRenderer[T]) SetOnValue(value T) {
	s.onVal = value
}

func (s *SextantRenderer[T]) Clone() *SextantRenderer[T] {
	dataCopy := make([]T, len(s.data))
	copy(dataCopy, s.data)

	return &SextantRenderer[T]{
		out:       s.out,
		cursor:    s.cursor,
		data:      dataCopy,
		onVal:     s.onVal,
		onFunc:    s.onFunc,
		lineWidth: s.lineWidth,
	}
}

func (s *SextantRenderer[T]) Append(values ...T) {
	s.data = append(s.data, values...)
}

func (s *SextantRenderer[T]) RenderAtCursor() string {
	indexes := [6]int{
		s.cursor,
		s.cursor + 1,
		s.cursor + s.lineWidth,
		s.cursor + s.lineWidth + 1,
		s.cursor + s.lineWidth*2,
		s.cursor + s.lineWidth*2 + 1,
	}

	bits := 0
	cursorCol := s.cursor % s.lineWidth

	for i, dataIdx := range indexes {
		if dataIdx >= len(s.data) {
			continue
		}

		if i%2 == 1 && cursorCol == s.lineWidth-1 {
			continue
		}

		data := s.data[dataIdx]
		if data == s.onVal || (s.onFunc != nil && s.onFunc(data)) {
			bits |= 1 << i
		}
	}

	return string(rune(0x1FB00 + bits))
}

func (s *SextantRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := s.cursor
	s.cursor = 0

	for s.cursor < len(s.data) {
		sb.WriteString(s.RenderAtCursor())

		oldPos := s.cursor
		s.Advance()

		if s.cursor-oldPos > 2 {
			sb.WriteString("\n")
		}
	}

	s.cursor = oldCursor
	return sb.String()
}

func (s *SextantRenderer[T]) WriteSingle() bool {
	s.out.WriteString(s.RenderAtCursor())

	oldCursor := s.cursor
	s.Advance()

	return s.cursor-oldCursor > 2
}

func (s *SextantRenderer[T]) WriteLine() bool {
	if s.cursor >= len(s.data) {
		return false
	}

	for s.cursor < len(s.data) {
		s.out.WriteString(s.RenderAtCursor())

		oldCursor := s.cursor
		s.Advance()

		if s.cursor-oldCursor > 2 {
			s.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (s *SextantRenderer[T]) WriteAll() {
	s.cursor = 0

	for s.cursor < len(s.data) {
		s.out.WriteString(s.RenderAtCursor())

		oldCursor := s.cursor
		s.Advance()

		if s.cursor-oldCursor > 2 {
			s.out.WriteString("\n")
		}
	}
}

func (s *SextantRenderer[T]) Clear() {
	s.cursor = 0
	s.data = []T{}
}

func (s *SextantRenderer[T]) ResetCursor() {
	s.cursor = 0
}

func (s *SextantRenderer[T]) HasMore() bool {
	return s.cursor < len(s.data)
}

func (s *SextantRenderer[T]) Len() int {
	return len(s.data)
}

func (s *SextantRenderer[T]) Progress() (cursor, total int) {
	return s.cursor, len(s.data)
}

func (s *SextantRenderer[T]) Advance() {
	s.cursor += 2
	switch s.cursor % s.lineWidth {
	case 0:
		s.cursor += s.lineWidth * 2
	case 1:
		s.cursor += (s.lineWidth * 2) - 1
	}
}
