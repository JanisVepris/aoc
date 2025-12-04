package day04

import (
	"fmt"

	"janisvepris/aoc/internal/files"
)

var (
	grid    [][]rune
	height  int
	width   int
	offsets = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top
		{0, -1}, {0, 1}, // mid
		{1, -1}, {1, 0}, {1, 1}, // bottom
	}
)

func Setup() {
	lines := files.ReadFile("2025/day04/input.txt")
	width = len(lines[0])
	height = len(lines)

	grid = make([][]rune, height)
	for y := range height {
		grid[y] = make([]rune, width)
		for range width {
			grid[y] = []rune(lines[y])
		}
	}
}

func Part1() {
	result := 0

	for y := range height {
		for x := range width {
			if grid[y][x] != '@' {
				continue
			}

			neighbors := 0

			for _, offset := range offsets {
				ny := y + offset[0]
				nx := x + offset[1]

				if ny < 0 || ny >= height || nx < 0 || nx >= width {
					continue
				}

				if grid[ny][nx] == '@' {
					neighbors += 1
				}
			}

			if neighbors < 4 {
				result += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	counts := make([][]int, height)
	for y := range height {
		counts[y] = make([]int, width)
	}

	paperPositions := [][2]int{} // y,x

	for y := range height {
		for x := range width {
			if grid[y][x] != '@' {
				continue
			}

			paperPositions = append(paperPositions, [2]int{y, x})

			for _, offset := range offsets {
				ny := y + offset[0]
				nx := x + offset[1]

				if ny < 0 || ny >= height || nx < 0 || nx >= width {
					continue
				}

				counts[ny][nx] += 1
			}
		}
	}

	result := 0

	for {
		toRemove := [][2]int{}

		for _, pos := range paperPositions {
			if grid[pos[0]][pos[1]] == '@' && counts[pos[0]][pos[1]] < 4 {
				toRemove = append(toRemove, pos)
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, coord := range toRemove {
			y, x := coord[0], coord[1]
			grid[y][x] = 'x'
			result += 1
			for _, offset := range offsets {
				ny := y + offset[0]
				nx := x + offset[1]

				if ny < 0 || ny >= height || nx < 0 || nx >= width {
					continue
				}
				counts[ny][nx] -= 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
