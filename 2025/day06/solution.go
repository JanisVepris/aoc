package day06

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day06/input.txt")
}

func Part1() {
	result := 0

	ops := strings.Fields(lines[len(lines)-1])

	totals := make([]int, len(ops))
	for i, op := range ops {
		if op == "*" {
			totals[i] = 1
		}
	}

	for i := 0; i < len(lines)-1; i++ {
		fields := strings.Fields(lines[i])

		for j, field := range fields {
			n := conv.StrToInt(field)

			if ops[j] == "+" {
				totals[j] += n
			} else {
				totals[j] *= n
			}
		}
	}

	for _, total := range totals {
		result += total
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	opLineIdx := len(lines) - 1

	start := 0
	for start < len(lines[opLineIdx]) {
		op := lines[opLineIdx][start]
		if op == ' ' {
			start++
			continue
		}

		search := lines[opLineIdx][start+1:]
		found := strings.IndexAny(search, "+*")
		end := 0

		if found == -1 {
			end = len(lines[opLineIdx])
		} else {
			end = start + found
		}

		nums := []int{}
		numba := ""
		for i := start; i < end; i++ {
			for j := range opLineIdx {
				char := lines[j][i]
				if char == ' ' {
					continue
				}
				numba += string(lines[j][i])
			}

			if numba != "" {
				nums = append(nums, conv.StrToInt(numba))
			}

			numba = ""
		}

		total := 0

		if op == '*' {
			total = 1
		}

		for _, n := range nums {
			if op == '+' {
				total += n
			} else {
				total *= n
			}
		}

		result += total

		start = end
	}

	fmt.Printf("Part 2: %d\n", result)
}
