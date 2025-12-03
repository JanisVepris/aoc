package day03

import (
	"fmt"

	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day03/input.txt")
}

func Part1() {
	result := 0

	var max1, max2 int

	for _, line := range lines {
		max1, max2 = 0, 0
		for i, char := range line {
			value := int(char - '0')

			if value > max1 && i < len(line)-1 {
				max1 = value
				max2 = 0

				continue
			}

			if value > max2 {
				max2 = value
			}
		}

		result += max1*10 + max2
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range lines {
		numbers := make([]int, 0, 12)
		start := 0
		numbersNeeded := 12

		for numbersNeeded > 0 {
			end := len(line) - numbersNeeded + 1

			max := 0
			maxIdx := 0

			for i := start; i < end; i++ {
				value := int(line[i] - '0')

				if value > max {
					max = value
					maxIdx = i - start
				}
			}

			numbers = append(numbers, max)
			start += maxIdx + 1
			numbersNeeded--
		}

		jolts := 0
		for _, num := range numbers {
			jolts = jolts*10 + num
		}

		result += jolts

	}

	fmt.Printf("Part 2: %d\n", result)
}
