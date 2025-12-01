package day10

import (
	"fmt"
	"slices"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/files"
)

type (
	Direction int
	PipeType  rune
)

const (
	TOP    Direction = iota
	BOTTOM Direction = iota
	LEFT   Direction = iota
	RIGHT  Direction = iota

	NS    PipeType = '|' // ║
	EW    PipeType = '-' // ═
	NE    PipeType = 'L' // ╚
	NW    PipeType = 'J' // ╝
	SW    PipeType = '7' // ╗
	SE    PipeType = 'F' // ╔
	START PipeType = 'S'
	SPACE PipeType = '.'
)

var adjacencyMap = map[PipeType]map[Direction][]PipeType{
	START: {
		TOP:    {NS, SW, SE},
		BOTTOM: {NS, NE, NW},
		LEFT:   {EW, NE, SE},
		RIGHT:  {EW, NW, SW},
	},
	NE: {
		TOP:    {NS, SE, SW, START},
		BOTTOM: {},
		LEFT:   {},
		RIGHT:  {EW, SW, NW, START},
	},
	NS: {
		TOP:    {NS, SE, SW, START},
		BOTTOM: {NS, NE, NW, START},
		LEFT:   {},
		RIGHT:  {},
	},
	SE: {
		TOP:    {},
		BOTTOM: {NS, NE, NW, START},
		LEFT:   {},
		RIGHT:  {EW, NW, SW, START},
	},
	SW: {
		TOP:    {},
		BOTTOM: {NS, NE, NW, START},
		LEFT:   {EW, NE, SE, START},
		RIGHT:  {},
	},
	NW: {
		TOP:    {NS, SE, SW, START},
		BOTTOM: {},
		LEFT:   {EW, NE, SE, START},
		RIGHT:  {},
	},
	EW: {
		TOP:    {},
		BOTTOM: {},
		LEFT:   {EW, NE, SE, START},
		RIGHT:  {EW, NW, SW, START},
	},
}

type Element struct {
	p Point
	t PipeType
}

type Point struct {
	x int
	y int
}

type Polygon struct {
	elements []Element
}

func (polygon *Polygon) EnclosesElement(element Element) bool {
	wn := 0 // winding number

	for i := 0; i < len(polygon.elements); i++ {
		cur := polygon.elements[i]
		next := polygon.elements[(i+1)%len(polygon.elements)]

		if cur.p.y <= element.p.y {
			if next.p.y > element.p.y && isLeft(cur, next, element) > 0 {
				wn++
			}
		} else {
			if next.p.y <= element.p.y && isLeft(cur, next, element) < 0 {
				wn--
			}
		}
	}

	return wn != 0
}

func isLeft(p1, p2, p3 Element) int {
	return (p2.p.x-p1.p.x)*(p3.p.y-p1.p.y) - (p3.p.x-p1.p.x)*(p2.p.y-p1.p.y)
}

type Grid [][]Element

func (g *Grid) Get(p Point) Element {
	return (*g)[p.y][p.x]
}

var (
	lines   []string
	grid    Grid
	polygon Polygon
)

func Setup() {
	lines = files.ReadFile("2023/day10/input.txt")
	grid = buildGrid(lines)
	startingPoint := findStartElement(grid)

	polygon, _ = step(startingPoint, startingPoint, grid, 0)
}

func Part1() {
	result := len(polygon.elements) / 2
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range grid {
		for _, element := range line {
			if slices.Contains(polygon.elements, element) {
				continue
			}

			if polygon.EnclosesElement(element) {
				result++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func step(currentElement Element, previousElement Element, grid Grid, stepIdx int) (Polygon, bool) {
	possibleSteps, _ := findConnectedPipes(currentElement, grid)
	possibleSteps = array.Filter(possibleSteps, func(e Element) bool { return e.p != previousElement.p })

	if len(possibleSteps) == 0 {
		return Polygon{}, false
	}

	startIndex := slices.IndexFunc(possibleSteps, func(e Element) bool { return e.t == START })

	if startIndex != -1 {
		element := possibleSteps[startIndex]

		return Polygon{elements: []Element{element}}, true
	}

	result := Polygon{}

	for _, element := range possibleSteps {
		polygon, isSuccess := step(element, currentElement, grid, stepIdx+1)

		if isSuccess {
			result.elements = polygon.elements
			result.elements = append(result.elements, element)

			return result, true
		}
	}

	panic("Could not find a path")
}

func findStartElement(grid Grid) Element {
	for _, line := range grid {
		for _, element := range line {
			if element.t == START {
				return element
			}
		}
	}

	panic("Could not find starting position")
}

func buildGrid(lines []string) Grid {
	grid := make(Grid, len(lines))

	for y, line := range lines {
		grid[y] = make([]Element, len(line))

		for x, char := range line {
			grid[y][x] = Element{Point{x, y}, PipeType(char)}
		}
	}

	return grid
}

func findConnectedPipes(pipe Element, grid Grid) ([]Element, int) {
	adjacencyMap := adjacencyMap[pipe.t]
	adjacentPipes := make([]Element, 0)

	// top
	if pipe.p.y > 0 {
		top := grid[pipe.p.y-1][pipe.p.x]

		if slices.Contains(adjacencyMap[TOP], top.t) {
			adjacentPipes = append(adjacentPipes, top)
		}
	}

	// bottom
	if pipe.p.y < len(grid)-1 {
		bottom := grid[pipe.p.y+1][pipe.p.x]

		if slices.Contains(adjacencyMap[BOTTOM], bottom.t) {
			adjacentPipes = append(adjacentPipes, bottom)
		}
	}

	// left
	if pipe.p.x > 0 {
		left := grid[pipe.p.y][pipe.p.x-1]

		if slices.Contains(adjacencyMap[LEFT], left.t) {
			adjacentPipes = append(adjacentPipes, left)
		}
	}

	// right
	if pipe.p.x < len(grid[0])-1 {
		right := grid[pipe.p.y][pipe.p.x+1]

		if slices.Contains(adjacencyMap[RIGHT], right.t) {
			adjacentPipes = append(adjacentPipes, right)
		}
	}

	return adjacentPipes, len(adjacentPipes)
}
