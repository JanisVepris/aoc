package day09

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/geom/rect"
)

var points [][2]int

type Edge struct {
	x1, y1 int
	x2, y2 int
}
type Poly struct {
	Vertices [][2]int
	cache    map[uint64]bool
	hEdges   []Edge
	vEdges   []Edge
}

func (p *Poly) InitEdges() {
	if p.hEdges != nil || p.vEdges != nil {
		return
	}

	n := len(p.Vertices)
	p.hEdges = make([]Edge, 0, n)
	p.vEdges = make([]Edge, 0, n)

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		x1, y1 := p.Vertices[j][0], p.Vertices[j][1]
		x2, y2 := p.Vertices[i][0], p.Vertices[i][1]

		if y1 == y2 { // horizontal
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			p.hEdges = append(p.hEdges, Edge{x1: x1, y1: y1, x2: x2, y2: y2})
		} else if x1 == x2 { // vertical
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			p.vEdges = append(p.vEdges, Edge{x1: x1, y1: y1, x2: x2, y2: y2})
		}
	}
}

func (p *Poly) RectEdgesCrossPolygon(minX, maxX, minY, maxY int) bool {
	topY := minY
	bottomY := maxY
	leftX := minX
	rightX := maxX

	// horizontal rect edges vs vertical polygon edges
	for _, e := range p.vEdges {
		px := e.x1
		py1 := e.y1
		py2 := e.y2

		if px > minX && px < maxX && py1 < topY && py2 > topY {
			return true
		}

		if px > minX && px < maxX && py1 < bottomY && py2 > bottomY {
			return true
		}
	}

	// vertical rect edges vs horizontal polygon edges
	for _, e := range p.hEdges {
		py := e.y1
		px1 := e.x1
		px2 := e.x2

		// left
		if py > minY && py < maxY && px1 < leftX && px2 > leftX {
			return true
		}

		// right
		if py > minY && py < maxY && px1 < rightX && px2 > rightX {
			return true
		}
	}

	return false
}

func (p *Poly) ContainsPointInclusive(x, y int) bool {
	key := posToKey(x, y)

	if result, ok := p.cache[posToKey(x, y)]; ok {
		return result
	}

	inside := false

	n := len(p.Vertices)

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := p.Vertices[i][0], p.Vertices[i][1]
		xj, yj := p.Vertices[j][0], p.Vertices[j][1]

		// boundary check
		if (y-yi)*(xj-xi) == (x-xi)*(yj-yi) &&
			((xi <= x && x <= xj) || (xj <= x && x <= xi)) &&
			((yi <= y && y <= yj) || (yj <= y && y <= yi)) {
			p.cache[key] = true
			return true
		}

		intersect := ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi)+xi)
		if intersect {
			inside = !inside
		}
	}

	p.cache[key] = inside
	return inside
}

func Setup() {
	lines := files.ReadFile("2025/day09/input.txt")

	points = make([][2]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")

		points[i] = [2]int{
			conv.StrToInt(parts[0]),
			conv.StrToInt(parts[1]),
		}
	}
}

func Part1() {
	maxArea := 0
	area := 0
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			rectangle := rect.NewRect(a[0], a[1], b[0], b[1])
			area = rectangle.AreaGrid()

			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Part 1: %d\n", maxArea)
}

func Part2() {
	poly := &Poly{
		Vertices: points,
		cache:    make(map[uint64]bool),
	}
	poly.InitEdges()

	maxArea := 0

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			rectangle := rect.NewRect(a[0], a[1], b[0], b[1])

			minX, maxX, minY, maxY := rectangle.GetExtremes()

			// check corners first
			if !poly.ContainsPointInclusive(minX, minY) ||
				!poly.ContainsPointInclusive(maxX, minY) ||
				!poly.ContainsPointInclusive(maxX, maxY) ||
				!poly.ContainsPointInclusive(minX, maxY) {
				continue
			}

			if poly.RectEdgesCrossPolygon(minX, maxX, minY, maxY) {
				continue
			}

			area := rectangle.AreaGrid()
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Part 2: %d\n", maxArea)
}

func posToKey(x, y int) uint64 {
	return uint64(x)<<32 | uint64(y)
}
