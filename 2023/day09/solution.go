package day09

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var (
	lines     []string
	sequences [][]int
)

func Setup() {
	lines = files.ReadFile("2023/day09/input.txt")
	sequences = parseSequences(lines)
}

func Part1() {
	result := 0
	for _, seq := range sequences {
		ex := extrapolate(seq)

		last := ex + seq[len(seq)-1]

		result += last
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	for _, seq := range sequences {
		slices.Reverse(seq)
		ex := extrapolate(seq)

		last := ex + seq[len(seq)-1]

		result += last
	}
	fmt.Printf("Part 2: %d\n", result)
}

func extrapolate(sequence []int) int {
	diffSeq := []int{}
	allZero := true

	last := 0

	for i := 0; i < len(sequence)-1; i++ {
		diff := sequence[i+1] - sequence[i]

		if diff != 0 {
			allZero = false
		}

		diffSeq = append(diffSeq, diff)
		last = diff
	}

	if allZero {
		return 0
	}

	lastInner := extrapolate(diffSeq)

	return last + lastInner
}

func parseSequences(lines []string) (sequences [][]int) {
	for _, line := range lines {
		numbers := strings.Fields(line)

		sequences = append(sequences, array.Map(numbers, func(idx int, number string) int { return conv.StrToInt(number) }))
	}

	return
}
