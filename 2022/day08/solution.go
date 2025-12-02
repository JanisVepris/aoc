package day08

import (
	"fmt"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

var (
	lines []string
	grid  [][]int
)

func Setup() {
	lines = files.ReadFile("2022/day08/input.txt")

	grid = make([][]int, len(lines))

	for y := 0; y < len(lines); y++ {
		grid[y] = make([]int, len(lines[y]))

		for x := 0; x < len(lines[y]); x++ {
			grid[y][x] = conv.ToInt(lines[y][x])
		}
	}
}

func Part1() {
	result := 0

	type key struct {
		x, y, dx, dy int
	}
	cache := make(map[key]int)

	var findMaxHeight func(x, y, dx, dy int) int
	findMaxHeight = func(x, y, dx, dy int) int {
		k := key{x, y, dx, dy}

		if val, ok := cache[k]; ok {
			return val
		}

		currX := x + dx
		currY := y + dy

		if currX < 0 || currY < 0 || currX >= len(grid[0]) || currY >= len(grid) {
			cache[k] = -1
			return -1
		}

		maxH := grid[currY][currX]
		nextH := findMaxHeight(currX, currY, dx, dy)

		if nextH > maxH {
			maxH = nextH
		}

		cache[k] = maxH

		return maxH
	}

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			currentHeight := grid[y][x]

			maxRight := findMaxHeight(x, y, 0, 1)
			maxLeft := findMaxHeight(x, y, 0, -1)
			maxDown := findMaxHeight(x, y, 1, 0)
			maxUp := findMaxHeight(x, y, -1, 0)

			if currentHeight > maths.MinInt(maxRight, maxLeft, maxDown, maxUp) {
				result++
			}
		}
	}

	result += len(grid)*2 + (len(grid[0])-2)*2

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			currentHeight := conv.ToInt(lines[y][x])

			viewRight := 0
			for vx := x + 1; vx < len(lines[y]); vx++ {
				viewRight++
				if grid[y][vx] >= currentHeight {
					break
				}
			}

			viewLeft := 0
			for vx := x - 1; vx >= 0; vx-- {
				viewLeft++
				if grid[y][vx] >= currentHeight {
					break
				}
			}

			viewDown := 0
			for vy := y + 1; vy < len(lines); vy++ {
				viewDown++
				if grid[vy][x] >= currentHeight {
					break
				}
			}

			viewUp := 0
			for vy := y - 1; vy >= 0; vy-- {
				viewUp++
				if grid[vy][x] >= currentHeight {
					break
				}
			}

			score := viewRight * viewLeft * viewDown * viewUp
			if score > result {
				result = score
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
