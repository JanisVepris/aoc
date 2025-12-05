package day12

import (
	"fmt"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/files"
)

var (
	grid          [][]int
	width, height int
	startPos      [2]int
	endPos        [2]int
	neighbors     = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func Setup() {
	lines := files.ReadFile("2022/day12/input.txt")

	grid = make([][]int, len(lines))

	for row, line := range lines {
		grid[row] = make([]int, len(line))

		for col, ch := range line {
			if ch == 'S' {
				startPos = [2]int{row, col}
				grid[row][col] = 0
				ch = 'a'
			}
			if ch == 'E' {
				endPos = [2]int{row, col}
				grid[row][col] = 25
				ch = 'z'
			}

			grid[row][col] = int(ch) - int('a')
		}
	}

	width = len(grid[0])
	height = len(grid)
}

func Part1() {
	result := 0

	visited := make(map[uint64]bool)

	type QItem struct {
		pos   [2]int // row, col
		steps int
	}

	queue := []QItem{}
	array.Push(&queue, QItem{pos: startPos, steps: 0})

	for len(queue) > 0 {
		current := array.ShiftInPlace(&queue)

		// already visited skip
		if _, ok := visited[posToKey(current.pos)]; ok {
			continue
		}

		visited[posToKey(current.pos)] = true

		// found the end
		if current.pos == endPos {
			result = current.steps
			break
		}

		currentHeight := grid[current.pos[0]][current.pos[1]]

		// collect neighbors
		for _, neighbor := range neighbors {
			nr := current.pos[0] + neighbor[0]
			nc := current.pos[1] + neighbor[1]

			// out of bounds
			if nr < 0 || nr >= height || nc < 0 || nc >= width {
				continue
			}

			neighborHeight := grid[nr][nc]

			// too hight to reach
			if currentHeight-neighborHeight < -1 {
				continue
			}

			// add to queue
			array.Push(&queue, QItem{pos: [2]int{nr, nc}, steps: current.steps + 1})
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	fmt.Printf("Part 2: %d\n", result)

	visited := make(map[uint64]bool)

	type QItem struct {
		pos   [2]int // row, col
		steps int
	}

	queue := []QItem{}
	array.Push(&queue, QItem{pos: endPos, steps: 0})

	// do the same but from the end and search for height 0
	for len(queue) > 0 {
		current := array.ShiftInPlace(&queue)

		// already visited skip
		if _, ok := visited[posToKey(current.pos)]; ok {
			continue
		}

		visited[posToKey(current.pos)] = true

		currentHeight := grid[current.pos[0]][current.pos[1]]

		// found it!
		if currentHeight == 0 {
			result = current.steps
			break
		}

		for _, neighbor := range neighbors {
			nr := current.pos[0] + neighbor[0]
			nc := current.pos[1] + neighbor[1]

			// out of bounds
			if nr < 0 || nr >= height || nc < 0 || nc >= width {
				continue
			}

			neighborHeight := grid[nr][nc]

			// neighbor height can only be lower by 1
			if currentHeight-neighborHeight > 1 {
				continue
			}

			array.Push(&queue, QItem{pos: [2]int{nr, nc}, steps: current.steps + 1})
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func posToKey(pos [2]int) uint64 {
	return uint64(pos[0])<<32 | uint64(pos[1])
}
