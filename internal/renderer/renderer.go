package renderer

import "os"

type Renderer[T comparable] interface {
	SetData(data []T) *BrailleRenderer[T]
	SetOutput(output *os.File) *BrailleRenderer[T]
	SetLineWidth(lineWidth int) *BrailleRenderer[T]
	SetOnFunc(onFunc func(value T) bool) *BrailleRenderer[T]
	SetOnValue(value T) *BrailleRenderer[T]
	Append(values ...T) *BrailleRenderer[T]
	Clear()
	ResetCursor()

	RenderAtCursor() string
	RenderToString() string

	WriteSingle() bool
	WriteLine() bool
	WriteAll()

	HasMore() bool
	Len() int
	Progress() (cursor, total int)

	Clone() *BrailleRenderer[T]
	Advance()
}
