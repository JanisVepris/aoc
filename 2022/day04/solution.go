package day04

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day04/input.txt")
}

func Part1() {
	result := 0

	for _, line := range lines {
		a1, a2, b1, b2 := getNumbers(line)

		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			result++
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range lines {
		a1, a2, b1, b2 := getNumbers(line)
		if a1 <= b2 && a2 >= b1 {
			result++
		}

	}

	fmt.Printf("Part 2: %d\n", result)
}

func getNumbers(s string) (a1, a2, b1, b2 int) {
	pair1, pair2, _ := strings.Cut(s, ",")
	a1s, a2s, _ := strings.Cut(pair1, "-")
	b1s, b2s, _ := strings.Cut(pair2, "-")
	a1, a2, b1, b2 = conv.StrToInt(a1s), conv.StrToInt(a2s), conv.StrToInt(b1s), conv.StrToInt(b2s)

	return
}
