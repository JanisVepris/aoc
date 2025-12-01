package day16

import (
	"fmt"
	"slices"

	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

type Direction int

const (
	UP    Direction = iota
	RIGHT Direction = iota
	DOWN  Direction = iota
	LEFT  Direction = iota
)

type Cell struct {
	X, Y           int
	Type           rune
	Lit            bool
	EmitDirections []Direction
}

type Grid struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func (g *Grid) Get(x, y int) Cell {
	return g.Cells[y][x]
}

func (g *Grid) Set(c Cell) {
	g.Cells[c.Y][c.X] = c
}

func (g *Grid) Copy() *Grid {
	newGrid := Grid{
		Width:  g.Width,
		Height: g.Height,
		Cells:  make([][]Cell, len(g.Cells)),
	}

	for i := 0; i < len(g.Cells); i++ {
		newGrid.Cells[i] = make([]Cell, len(g.Cells[i]))
		copy(newGrid.Cells[i], g.Cells[i])
	}

	return &newGrid
}

func (g *Grid) CountLit() int {
	result := 0

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Get(x, y)
			if cell.Lit {
				result++
			}
		}
	}

	return result
}

const (
	M_SLASH      rune = '/'
	M_BACK       rune = '\\'
	M_HORIZONTAL rune = '-'
	M_VERTICAL   rune = '|'
	SPACE        rune = '.'
)

var (
	lines []string
	grid  *Grid
)

func Setup() {
	lines = files.ReadFile("2023/day16/input.txt")
	grid = parseGrid(lines)
}

func Part1() {
	g := grid.Copy()
	traverse(g, 0, 0, RIGHT)

	result := g.CountLit()

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	// check left side
	for y := 0; y < grid.Height; y++ {
		g := grid.Copy()

		traverse(g, 0, y, RIGHT)

		res := g.CountLit()

		result = maths.MaxInt(result, res)
	}

	// check right side
	for y := 0; y < grid.Height; y++ {
		g := grid.Copy()
		x := grid.Width - 1

		traverse(g, x, y, LEFT)

		res := g.CountLit()

		result = maths.MaxInt(result, res)
	}

	// check top side
	for x := 0; x < grid.Width; x++ {
		g := grid.Copy()

		traverse(g, x, 0, DOWN)

		res := g.CountLit()

		result = maths.MaxInt(result, res)
	}

	// check bottom side
	for x := 0; x < grid.Width; x++ {
		g := grid.Copy()
		y := grid.Height - 1
		traverse(g, x, y, UP)

		res := g.CountLit()

		result = maths.MaxInt(result, res)
	}

	fmt.Printf("Part 2: %d\n", result)
}

func traverse(grid *Grid, x, y int, dir Direction) {
	if x < 0 || x > grid.Width-1 || y < 0 || y > grid.Height-1 {
		return
	}

	cell := grid.Get(x, y)

	cell.Lit = true

	nextDirs := getNextDirs(cell, dir)
	newDirs := make([]Direction, 0)

	for _, nextDir := range nextDirs {
		if slices.Contains(cell.EmitDirections, nextDir) {
			continue
		}

		cell.EmitDirections = append(cell.EmitDirections, nextDir)
		newDirs = append(newDirs, nextDir)
	}

	grid.Set(cell)

	if len(newDirs) == 0 {
		return
	}

	for _, newDir := range newDirs {
		switch newDir {
		case UP:
			traverse(grid, x, y-1, newDir)
		case RIGHT:
			traverse(grid, x+1, y, newDir)
		case DOWN:
			traverse(grid, x, y+1, newDir)
		case LEFT:
			traverse(grid, x-1, y, newDir)
		}
	}
}

func getNextDirs(currentCell Cell, currentDir Direction) []Direction {
	switch currentCell.Type {
	case M_SLASH:
		switch currentDir {
		case UP:
			return []Direction{RIGHT}
		case RIGHT:
			return []Direction{UP}
		case DOWN:
			return []Direction{LEFT}
		case LEFT:
			return []Direction{DOWN}
		}
	case M_BACK:
		switch currentDir {
		case UP:
			return []Direction{LEFT}
		case RIGHT:
			return []Direction{DOWN}
		case DOWN:
			return []Direction{RIGHT}
		case LEFT:
			return []Direction{UP}
		}
	case M_HORIZONTAL:
		switch currentDir {
		case UP:
			return []Direction{LEFT, RIGHT}
		case RIGHT:
			return []Direction{RIGHT}
		case DOWN:
			return []Direction{LEFT, RIGHT}
		case LEFT:
			return []Direction{LEFT}
		}
	case M_VERTICAL:
		switch currentDir {
		case UP:
			return []Direction{UP}
		case RIGHT:
			return []Direction{UP, DOWN}
		case DOWN:
			return []Direction{DOWN}
		case LEFT:
			return []Direction{UP, DOWN}
		}
	case SPACE:
		return []Direction{currentDir}
	}

	panic("Unknown cell type")
}

func parseGrid(lines []string) *Grid {
	gridHeight := len(lines)
	gridWidth := len(lines[0])
	grid := Grid{
		Width:  gridWidth,
		Height: gridHeight,
		Cells:  make([][]Cell, gridHeight),
	}

	for y, row := range lines {
		grid.Cells[y] = make([]Cell, gridWidth)

		for x, char := range row {
			grid.Cells[y][x] = Cell{
				X:    x,
				Y:    y,
				Type: char,
			}
		}
	}

	for y, line := range lines {
		for x, char := range line {
			cell := grid.Cells[y][x]
			cell.Type = char
		}
	}

	return &grid
}
