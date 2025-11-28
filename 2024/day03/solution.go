package day03

import (
	"fmt"
	"regexp"
	"strings"

	"janisvepris/aoc25/internal/array"
	"janisvepris/aoc25/internal/conv"
	"janisvepris/aoc25/internal/files"
)

var (
	lines []string
	input string = ""
)

func Setup() {
	lines = files.ReadFile("2024/day03/input.txt")

	for _, line := range lines {
		input += line
	}
}

func Part1() {
	instructions := getInstructions(input)

	result := array.Reduce(instructions, func(carry int, instruction []int) int {
		return carry + (instruction[0] * instruction[1])
	}, 0)
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	cleanString := parseDoInstructions(input)

	instructions := getInstructions(cleanString)

	result := array.Reduce(instructions, func(carry int, instruction []int) int {
		return carry + (instruction[0] * instruction[1])
	}, 0)
	fmt.Printf("Part 2: %d\n", result)
}

func getInstructions(input string) [][]int {
	regex := `mul\((\d{1,3}),(\d{1,3})\)`

	pattern := regexp.MustCompile(regex)

	result := pattern.FindAllStringSubmatch(input, -1)

	return array.Map(result, func(i int, instruction []string) []int {
		return parseInstruction(instruction)
	})
}

func parseInstruction(instruction []string) []int {
	return []int{conv.StrToInt(instruction[1]), conv.StrToInt(instruction[2])}
}

func parseDoInstructions(input string) string {
	remaining := strings.Clone(input)
	result := ""

	regex := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)
	instDo := "do()"
	instDont := "don't()"

	do := true

	for len(remaining) != 0 {
		if strings.HasPrefix(remaining, instDo) {
			do = true
			remaining = strings.TrimPrefix(remaining, instDo)
			continue
		}

		if strings.HasPrefix(remaining, instDont) {
			do = false
			doIndex := strings.Index(remaining, instDo)

			if doIndex == -1 {
				break
			}

			remaining = remaining[doIndex:]

			continue
		}

		if !do {
			remaining = remaining[1:]
			continue
		}

		match := regex.FindString(remaining)

		if match != "" {
			result += match
			remaining = remaining[len(match):]
			continue
		}

		remaining = remaining[1:]
	}

	return result
}
