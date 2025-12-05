package day03

import (
	"fmt"
	"slices"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var (
	lines   []string
	offsets = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top
		{0, -1}, {0, 1}, // middle
		{1, -1}, {1, 0}, {1, 1}, // bottom
	}
)

func Setup() {
	lines = files.ReadFile("2023/day03/input.txt")
}

func Part1() {
	result := 0
	number := ""
	isValid := false

	for row := range lines {
		for col, ch := range lines[row] {
			if ch < '0' || ch > '9' {
				if isValid {
					result += conv.StrToInt(number)
					isValid = false
				}
				if number != "" {
					number = ""
				}
				continue
			}

			number += string(ch)

			for _, offset := range offsets {
				nr := row + offset[0]
				nc := col + offset[1]

				if nr < 0 || nr >= len(lines[0]) || nc < 0 || nc >= len(lines) {
					continue
				}

				neighbor := lines[nr][nc]

				if (neighbor < '0' || neighbor > '9') && neighbor != '.' {
					isValid = true
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	number := ""
	isValid := false
	gearRatios := map[uint64][]int{}
	gears := [][2]int{}

	for row := range lines {
		for col, ch := range lines[row] {
			if ch < '0' || ch > '9' {
				if isValid {
					isValid = false

					for _, gear := range gears {
						gearRatios[posToKey(gear)] = append(gearRatios[posToKey(gear)], conv.StrToInt(number))
					}
				}
				if number != "" {
					number = ""
				}

				if len(gears) > 0 {
					gears = [][2]int{}
				}

				continue
			}

			number += string(ch)

			for _, offset := range offsets {
				nr := row + offset[0]
				nc := col + offset[1]

				if nr < 0 || nr >= len(lines[0]) || nc < 0 || nc >= len(lines) {
					continue
				}

				neighbor := lines[nr][nc]

				if (neighbor < '0' || neighbor > '9') && neighbor != '.' {
					isValid = true
				}

				if neighbor == '*' {
					if !slices.Contains(gears, [2]int{nr, nc}) {
						array.Push(&gears, [2]int{nr, nc})
					}
				}
			}
		}
	}

	for _, ratios := range gearRatios {
		if len(ratios) != 2 {
			continue
		}

		result += ratios[0] * ratios[1]
	}

	fmt.Printf("Part 2: %d\n", result)
}

func posToKey(pos [2]int) uint64 {
	return uint64(pos[0])<<32 | uint64(pos[1])
}
