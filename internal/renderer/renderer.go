package renderer

import "os"

type Renderer[T comparable] interface {
	SetData(data []T)
	SetOutput(output *os.File)
	SetLineWidth(lineWidth int)
	SetOnFunc(onFunc func(value T) bool)
	SetOnValue(value T)
	Append(values ...T)
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

	Advance()
}
