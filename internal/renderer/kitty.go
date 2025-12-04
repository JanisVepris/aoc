package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type KittyRenderer[T comparable] struct {
	out        *os.File
	cursor     int
	data       []T
	onVal      T
	onFunc     func(value T) bool
	lineWidth  int
	pixelScale int
	colorFunc  func(value T, isOn bool, w, h, x, y int) color.Color
	bgColor    color.Color
}

func NewKittyRenderer[T comparable](lineWidth int) *KittyRenderer[T] {
	return &KittyRenderer[T]{
		lineWidth:  lineWidth,
		cursor:     0,
		out:        os.Stdout,
		pixelScale: 4,
		colorFunc:  defaultKittyColorFunc[T](),
		bgColor:    color.RGBA{0, 0, 0, 255},
	}
}

func defaultKittyColorFunc[T comparable]() func(value T, isOn bool, w, h, x, y int) color.Color {
	return func(value T, isOn bool, w, h, x, y int) color.Color {
		if isOn {
			return color.RGBA{255, 255, 255, 255}
		}

		return color.RGBA{0, 0, 0, 255}
	}
}

func (k *KittyRenderer[T]) SetData(data []T) {
	k.data = data
	k.cursor = 0
}

func (k *KittyRenderer[T]) SetOutput(output *os.File) {
	k.out = output
}

func (k *KittyRenderer[T]) SetLineWidth(lineWidth int) {
	k.lineWidth = lineWidth
}

func (k *KittyRenderer[T]) SetOnFunc(onFunc func(value T) bool) {
	k.onFunc = onFunc
}

func (k *KittyRenderer[T]) SetOnValue(value T) {
	k.onVal = value
}

func (k *KittyRenderer[T]) SetPixelScale(scale int) {
	k.pixelScale = scale
}

func (k *KittyRenderer[T]) SetColorFunc(colorFunc func(value T, isOn bool, w, h, x, y int) color.Color) {
	k.colorFunc = colorFunc
}

func (k *KittyRenderer[T]) SetBackgroundColor(bgColor color.Color) {
	k.bgColor = bgColor
}

func (k *KittyRenderer[T]) Clone() *KittyRenderer[T] {
	dataCopy := make([]T, len(k.data))
	copy(dataCopy, k.data)

	return &KittyRenderer[T]{
		out:        k.out,
		cursor:     k.cursor,
		data:       dataCopy,
		onVal:      k.onVal,
		onFunc:     k.onFunc,
		lineWidth:  k.lineWidth,
		pixelScale: k.pixelScale,
		colorFunc:  k.colorFunc,
		bgColor:    k.bgColor,
	}
}

func (k *KittyRenderer[T]) Append(values ...T) {
	k.data = append(k.data, values...)
}

func (k *KittyRenderer[T]) RenderAtCursor() string {
	return ""
}

func (k *KittyRenderer[T]) createImage() *image.RGBA {
	if len(k.data) == 0 {
		return image.NewRGBA(image.Rect(0, 0, 1, 1))
	}

	height := (len(k.data) + k.lineWidth - 1) / k.lineWidth
	width := k.lineWidth

	imgWidth := width * k.pixelScale
	imgHeight := height * k.pixelScale

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for y := range height {
		for x := range width {
			idx := y*k.lineWidth + x
			if idx >= len(k.data) {
				continue
			}

			data := k.data[idx]
			isOn := data == k.onVal || (k.onFunc != nil && k.onFunc(data))

			pixelColor := k.colorFunc(data, isOn, width, height, x, y)

			for py := 0; py < k.pixelScale; py++ {
				for px := 0; px < k.pixelScale; px++ {
					img.Set(x*k.pixelScale+px, y*k.pixelScale+py, pixelColor)
				}
			}
		}
	}

	return img
}

func (k *KittyRenderer[T]) encodeKittyGraphics(img *image.RGBA) string {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return ""
	}

	pngData := buf.Bytes()
	encoded := base64.StdEncoding.EncodeToString(pngData)

	const chunkSize = 4096
	var sb strings.Builder

	for i := 0; i < len(encoded); i += chunkSize {
		end := min(i+chunkSize, len(encoded))

		chunk := encoded[i:end]
		more := 0
		if end < len(encoded) {
			more = 1
		}

		if i == 0 {
			sb.WriteString(fmt.Sprintf("\033_Ga=T,f=100,m=%d;%s\033\\", more, chunk))
		} else {
			sb.WriteString(fmt.Sprintf("\033_Gm=%d;%s\033\\", more, chunk))
		}
	}

	return sb.String()
}

func (k *KittyRenderer[T]) RenderToString() string {
	img := k.createImage()
	return k.encodeKittyGraphics(img)
}

func (k *KittyRenderer[T]) WriteSingle() bool {
	return false
}

func (k *KittyRenderer[T]) WriteLine() bool {
	return false
}

func (k *KittyRenderer[T]) WriteAll() {
	img := k.createImage()
	kittyData := k.encodeKittyGraphics(img)
	k.out.WriteString(kittyData)
	k.out.WriteString("\n")
}

func (k *KittyRenderer[T]) Clear() {
	k.cursor = 0
	k.data = []T{}
}

func (k *KittyRenderer[T]) ResetCursor() {
	k.cursor = 0
}

func (k *KittyRenderer[T]) HasMore() bool {
	return false
}

func (k *KittyRenderer[T]) Len() int {
	return len(k.data)
}

func (k *KittyRenderer[T]) Progress() (cursor, total int) {
	return k.cursor, len(k.data)
}

func (k *KittyRenderer[T]) Advance() {
	k.cursor++
}
