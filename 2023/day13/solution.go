package day13

import (
	"fmt"
	"slices"

	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/logic"
)

var (
	lines    []string
	patterns [][]string
)

func Setup() {
	lines = files.ReadFile("2023/day13/input.txt")
	patterns = parsePatterns(lines)
}

func Part1() {
	result := 0
	for _, pattern := range patterns {
		value := find(pattern, 0, false) * 100

		if value != 0 {
			result += value
			continue
		}

		result += find(transpose(pattern), 0, false)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	for _, pattern := range patterns {
		value := find(pattern, 1, false) * 100

		if value != 0 {
			result += value
			continue
		}

		result += find(transpose(pattern), 1, false)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func find(pattern []string, diffReq int, reversed bool) int {
	for i := 2; i <= len(pattern); i += 2 {
		slice := make([]string, 0)

		for x := 0; x < i; x++ {
			slice = append(slice, pattern[x])
		}

		if isMirror(slice, diffReq) {
			return logic.If(reversed, len(pattern)-i/2, i/2)
		}
	}

	if reversed {
		return 0
	}

	reversedPattern := slices.Clone(pattern)
	slices.Reverse(reversedPattern)

	return find(reversedPattern, diffReq, true)
}

func isMirror(pattern []string, diffReq int) bool {
	differences := 0

	i1 := 0
	i2 := len(pattern) - 1

	for i1 < i2 {
		diffCount := compare(pattern[i1], pattern[i2])

		if diffCount+differences > diffReq {
			return false
		}

		differences += diffCount
		i1++
		i2--
	}

	return differences == diffReq
}

func compare(line1, line2 string) (diffCount int) {
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			diffCount++
		}
	}

	return
}

func transpose(pattern []string) []string {
	result := make([]string, 0)

	for i := 0; i < len(pattern[0]); i++ {
		row := ""

		for _, line := range pattern {
			row += string(line[i])
		}

		result = append(result, row)
	}

	return result
}

func parsePatterns(lines []string) [][]string {
	patterns := make([][]string, 0)
	pattern := make([]string, 0)

	for i, line := range lines {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
			continue
		}

		pattern = append(pattern, line)

		if i == len(lines)-1 {
			patterns = append(patterns, pattern)
		}
	}

	return patterns
}
