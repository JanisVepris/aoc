package day07

import (
	"fmt"

	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day07/input.txt")
}

func Part1() {
	result := 0

	beams := map[int]bool{}

	for col, ch := range lines[0] {
		if ch == 'S' {
			beams[col] = true
		}
	}

	for row := 1; row < len(lines); row++ {
		for col := range beams {
			ch := lines[row][col]

			if ch == '^' {
				result += 1
				delete(beams, col)
				beams[col-1] = true
				beams[col+1] = true
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	cache := map[uint64]int{}
	start := 0

	for col, ch := range lines[0] {
		if ch == 'S' {
			start = col
		}
	}

	var dfs func(rStart, cStart int) int

	dfs = func(rStart, cStart int) int {
		for i := rStart; i < len(lines); i++ {
			if val, ok := cache[posToKey(i, cStart)]; ok {
				return val
			}

			ch := lines[i][cStart]

			total := 0
			if ch == '^' {
				total += dfs(i, cStart-1)
				total += dfs(i, cStart+1)
				cache[posToKey(i, cStart)] = total

				return total
			}
		}

		return 1
	}

	result = dfs(1, start)

	fmt.Printf("Part 2: %d\n", result)
}

func posToKey(row, col int) uint64 {
	return uint64(row)<<32 | uint64(col)
}
