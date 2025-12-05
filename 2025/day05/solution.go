package day05

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day05/input.txt")
}

func Part1() {
	result := 0

	ranges := make([][2]int, 0)
	idStart := 0

	for i, line := range lines {
		if line == "" {
			idStart = i + 1
			break
		}

		parts := strings.Split(line, "-")
		ranges = append(ranges, [2]int{
			conv.StrToInt(parts[0]),
			conv.StrToInt(parts[1]),
		})
	}

	for i := idStart; i < len(lines); i++ {
		id := conv.StrToInt(lines[i])

		for _, r := range ranges {
			if r[0] <= id && id <= r[1] {
				result++
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	ranges := make([][2]int, 0)

	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		ranges = append(ranges, [2]int{
			conv.StrToInt(parts[0]),
			conv.StrToInt(parts[1]),
		})
	}

	for i := 0; i < len(ranges)-1; i++ {
		rA := ranges[i]

		for j := i + 1; j < len(ranges); j++ {
			rB := ranges[j]

			if rA[0] <= rB[1] && rB[0] <= rA[1] {
				ranges = array.RemoveElement(ranges, i)
				ranges = array.RemoveElement(ranges, j-1)
				ranges = append(ranges, [2]int{
					min(rA[0], rB[0]),
					max(rA[1], rB[1]),
				})

				i -= 1
				break
			}
		}
	}

	for _, r := range ranges {
		result += r[1] - r[0] + 1
	}

	fmt.Printf("Part 2: %d\n", result)
}
