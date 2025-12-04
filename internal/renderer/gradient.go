package renderer

import (
	"fmt"
	"os"
	"strings"
)

type GradientRenderer[T comparable] struct {
	out          *os.File
	cursor       int
	data         []T
	onVal        T
	onFunc       func(value T) bool
	lineWidth    int
	gradientFunc func(T) (r, g, b int)
}

func NewGradientRenderer[T comparable](lineWidth int) *GradientRenderer[T] {
	return &GradientRenderer[T]{
		lineWidth:    lineWidth,
		cursor:       0,
		out:          os.Stdout,
		gradientFunc: defaultGradientFunc[T](),
	}
}

func defaultGradientFunc[T comparable]() func(T) (r, g, b int) {
	return func(val T) (r, g, b int) {
		return 255, 0, 0
	}
}

func (g *GradientRenderer[T]) SetData(data []T) {
	g.data = data
	g.cursor = 0
}

func (g *GradientRenderer[T]) SetOutput(output *os.File) {
	g.out = output
}

func (g *GradientRenderer[T]) SetLineWidth(lineWidth int) {
	g.lineWidth = lineWidth
}

func (g *GradientRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	g.onFunc = onFunc
}

func (g *GradientRenderer[T]) SetOnValue(value T) {
	g.onVal = value
}

func (g *GradientRenderer[T]) SetGradientFunc(gradientFunc func(T) (r, g, b int)) {
	g.gradientFunc = gradientFunc
}

func (g *GradientRenderer[T]) Clone() *GradientRenderer[T] {
	dataCopy := make([]T, len(g.data))
	copy(dataCopy, g.data)

	return &GradientRenderer[T]{
		out:          g.out,
		cursor:       g.cursor,
		data:         dataCopy,
		onVal:        g.onVal,
		onFunc:       g.onFunc,
		lineWidth:    g.lineWidth,
		gradientFunc: g.gradientFunc,
	}
}

func (g *GradientRenderer[T]) Append(values ...T) {
	g.data = append(g.data, values...)
}

func (g *GradientRenderer[T]) RenderAtCursor() string {
	if g.cursor >= len(g.data) {
		return " "
	}

	data := g.data[g.cursor]
	isOn := data == g.onVal || (g.onFunc != nil && g.onFunc(data))

	if !isOn {
		return " "
	}

	r, gr, b := g.gradientFunc(data)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm█\033[0m", r, gr, b)
}

func (g *GradientRenderer[T]) RenderToString() string {
	var sb strings.Builder
	oldCursor := g.cursor
	g.cursor = 0

	for g.cursor < len(g.data) {
		sb.WriteString(g.RenderAtCursor())

		g.Advance()

		if g.cursor%g.lineWidth == 0 && g.cursor > 0 {
			sb.WriteString("\n")
		}
	}

	g.cursor = oldCursor
	return sb.String()
}

func (g *GradientRenderer[T]) WriteSingle() bool {
	g.out.WriteString(g.RenderAtCursor())

	oldCursor := g.cursor
	g.Advance()

	return g.cursor%g.lineWidth == 0 && oldCursor%g.lineWidth != 0
}

func (g *GradientRenderer[T]) WriteLine() bool {
	if g.cursor >= len(g.data) {
		return false
	}

	for g.cursor < len(g.data) {
		g.out.WriteString(g.RenderAtCursor())

		oldCursor := g.cursor
		g.Advance()

		if g.cursor%g.lineWidth == 0 && oldCursor%g.lineWidth != 0 {
			g.out.WriteString("\n")
			return true
		}
	}

	return false
}

func (g *GradientRenderer[T]) WriteAll() {
	g.cursor = 0

	for g.cursor < len(g.data) {
		g.out.WriteString(g.RenderAtCursor())

		g.Advance()

		if g.cursor%g.lineWidth == 0 && g.cursor > 0 {
			g.out.WriteString("\n")
		}
	}
}

func (g *GradientRenderer[T]) Clear() {
	g.cursor = 0
	g.data = []T{}
}

func (g *GradientRenderer[T]) ResetCursor() {
	g.cursor = 0
}

func (g *GradientRenderer[T]) HasMore() bool {
	return g.cursor < len(g.data)
}

func (g *GradientRenderer[T]) Len() int {
	return len(g.data)
}

func (g *GradientRenderer[T]) Progress() (cursor, total int) {
	return g.cursor, len(g.data)
}

func (g *GradientRenderer[T]) Advance() {
	g.cursor++
}
