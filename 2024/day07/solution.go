package day07

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2024/day07/input.txt")
}

func Part1() {
	result := 0

	for _, line := range lines {
		target, digits := parseLine(line)

		var dfs func(i, current int) bool
		dfs = func(i, current int) bool {
			if i == len(digits) {
				return current == target
			}

			next := digits[i]

			if dfs(i+1, current+next) {
				return true
			}

			if dfs(i+1, current*next) {
				return true
			}

			return false
		}

		isPossible := dfs(1, digits[0])

		if isPossible {
			result += target
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range lines {
		target, digits := parseLine(line)

		var dfs func(i, current int) bool
		dfs = func(i, current int) bool {
			if i == len(digits) {
				return current == target
			}

			next := digits[i]

			// try addition
			if dfs(i+1, current+next) {
				return true
			}

			// try multiplication
			if dfs(i+1, current*next) {
				return true
			}

			// try concat
			value := conv.StrToInt(fmt.Sprintf("%d%d", current, next))

			if dfs(i+1, value) {
				return true
			}

			return false
		}

		isPossible := dfs(1, digits[0])

		if isPossible {
			result += target
		}
	}
	fmt.Printf("Part 2: %d\n", result)
}

func parseLine(line string) (product int, operands []int) {
	parts := strings.Split(line, ":")

	product = conv.StrToInt(strings.TrimSpace(parts[0]))
	operands = array.Map(
		strings.Split(strings.TrimSpace(parts[1]), " "),
		func(_ int, s string) int {
			return conv.StrToInt(strings.TrimSpace(s))
		},
	)

	return product, operands
}
