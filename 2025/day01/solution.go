package day01

import (
	"fmt"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day01/input.txt")
}

func Part1() {
	result := 0
	pointer := 50

	for _, line := range lines {
		steps := conv.StrToInt(line[1:])
		if line[0] == 'L' {
			steps = -steps
		}

		pointer = ((pointer+steps)%100 + 100) % 100

		if pointer == 0 {
			result++
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	pointer := 50
	previous := pointer

	for _, line := range lines {
		passes := 0
		steps := conv.StrToInt(line[1:])

		if line[0] == 'L' {
			pointer -= steps

			if pointer < 0 {
				passes = (-pointer-1)/100 + 1
				pointer = pointer % 100
				if pointer < 0 {
					pointer += 100
				}
			}

			if previous == 0 {
				passes--
			}

			if pointer == 0 {
				result++
			}

			result += passes
		} else {
			pointer += steps

			if pointer > 100 {
				passes = (pointer-101)/100 + 1
				pointer = pointer % 100
				if pointer == 0 {
					pointer = 100
				}
			}

			if pointer == 100 {
				pointer = 0
			}

			if pointer == 0 {
				result++
			}

			result += passes
		}

		previous = pointer
	}
	fmt.Printf("Part 2: %d\n", result)
}
