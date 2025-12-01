package day01

import (
	"fmt"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day01/input.txt")
}

func Part1() {
	max := 0
	current := 0

	for _, line := range lines {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		current += conv.StrToInt(line)
	}

	fmt.Printf("Part 1: %d\n", max)
}

func Part2() {
	result := 0

	top1, top2, top3 := 0, 0, 0

	current := 0

	for _, line := range lines {
		if line == "" {
			if current > top1 {
				top1, top2, top3 = current, top1, top2
			} else if current > top2 {
				top2, top3 = current, top2
			} else if current > top3 {
				top3 = current
			}

			current = 0
			continue
		}

		current += conv.StrToInt(line)
	}

	result = top1 + top2 + top3

	fmt.Printf("Part 2: %d\n", result)
}
