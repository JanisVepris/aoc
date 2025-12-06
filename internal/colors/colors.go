package colors

import (
	"image/color"
	"math"
	"strconv"
)

func HsvToRGBA(h, s, v float64) color.RGBA {
	h = math.Mod(h, 1.0)
	i := int(h * 6)
	f := h*6 - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	var r, g, b float64

	switch i % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}

func HexToRGBA(s string) color.RGBA {
	if len(s) != 7 && len(s) != 9 {
		panic("color must be #RRGGBB or #RRGGBBAA")
	}
	if s[0] != '#' {
		panic("color must start with #")
	}

	r, err := strconv.ParseUint(s[1:3], 16, 8)
	if err != nil {
		return color.RGBA{}
	}
	g, err := strconv.ParseUint(s[3:5], 16, 8)
	if err != nil {
		return color.RGBA{}
	}
	b, err := strconv.ParseUint(s[5:7], 16, 8)
	if err != nil {
		return color.RGBA{}
	}

	a := uint8(255)
	if len(s) == 9 {
		v, err := strconv.ParseUint(s[7:9], 16, 8)
		if err != nil {
			return color.RGBA{}
		}
		a = uint8(v)
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: a,
	}
}
