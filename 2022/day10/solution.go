package day10

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day10/input.txt")
}

func Part1() {
	result := 0

	cycle := 0
	x := 1

	for _, line := range lines {
		increaseBy := 0
		if line == "noop" {
			increaseBy = 1
		} else {
			increaseBy = 2
		}

		for i := 0; i < increaseBy; i++ {
			cycle++
			if (cycle-20)%40 == 0 {
				result += cycle * x
			}
		}

		if line != "noop" {
			v := conv.ToInt(line[5:])
			x += v
		}

	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	cycle := 0
	x := 1
	var output strings.Builder
	output.Grow(40)
	for _, line := range lines {
		drawFrom := x - 1
		drawTo := x + 1

		increaseBy := 0
		if line == "noop" {
			increaseBy = 1
		} else {
			increaseBy = 2
		}

		for i := 0; i < increaseBy; i++ {
			horizontalPos := cycle % 40

			if horizontalPos == 0 {
				output.WriteByte('\n')
			}

			if horizontalPos >= drawFrom && horizontalPos <= drawTo {
				output.WriteString("█")
			} else {
				output.WriteByte(' ')
			}
			cycle++
		}

		if line != "noop" {
			v := conv.ToInt(line[5:])
			x += v
		}
	}
	fmt.Println("Part 2:", output.String())
}
