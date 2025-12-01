package day02

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

var (
	lines   []string
	reports [][]int
)

func Setup() {
	lines = files.ReadFile("2024/day02/input.txt")
	reports = make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		for _, part := range parts {
			reports[i] = append(reports[i], conv.StrToInt(part))
		}
	}
}

func Part1() {
	result := 0

	for _, report := range reports {
		if isSafe(report, false) {
			result++
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, report := range reports {
		if isSafe(report, true) {
			result++
		}
	}
	fmt.Printf("Part 2: %d\n", result)
}

func isSafe(report []int, errorTolerance bool) bool {
	i, j := 0, 1
	reportLen := len(report)

	isDecreasing := report[0] > report[1]

	for j < reportLen {
		val1 := report[i]
		val2 := report[j]
		hasError := false

		diff := maths.AbsInt(val1 - val2)

		if isDecreasing && val1 < val2 {
			hasError = true
		}

		if !isDecreasing && val1 > val2 {
			hasError = true
		}

		if diff < 1 || diff > 3 {
			hasError = true
		}

		if errorTolerance && hasError {
			variation1 := array.RemoveElement(report, i)
			if isSafe(variation1, false) {
				return true
			}

			variation2 := array.RemoveElement(report, j)

			if isSafe(variation2, false) {
				return true
			}

			if i != 0 {
				variation3 := array.RemoveElement(report, 0)

				if isSafe(variation3, false) {
					return true
				}
			}
		}

		if hasError {
			return false
		}

		i++
		j++
	}

	return true
}
