package day05

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day05/input.txt")
}

func Part1() {
	result := ""

	stacks, instructions := parseInput()

	for _, instr := range instructions {
		source := stacks[instr[1]]
		dest := stacks[instr[2]]

		crates := source[len(source)-instr[0]:]
		slices.Reverse(crates)

		stacks[instr[2]] = append(dest, crates...)
		stacks[instr[1]] = source[:len(source)-instr[0]]
	}

	for i := 1; i <= len(stacks); i++ {
		result += conv.ToStr(stacks[i][len(stacks[i])-1])
	}

	fmt.Printf("Part 1: %s\n", result)
}

func Part2() {
	result := ""

	stacks, instructions := parseInput()

	for _, instr := range instructions {
		source := stacks[instr[1]]
		dest := stacks[instr[2]]

		crates := source[len(source)-instr[0]:]

		stacks[instr[2]] = append(dest, crates...)
		stacks[instr[1]] = source[:len(source)-instr[0]]
	}

	for i := 1; i <= len(stacks); i++ {
		result += conv.ToStr(stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("Part 2: %s\n", result)
}

func parseInput() (stacks map[int][]rune, instructions [][]int) {
	stacks = make(map[int][]rune)
	instructions = [][]int{}
	parseInstructions := false

	for j, line := range lines {
		if line == "" {
			parseInstructions = true
			continue
		}

		if parseInstructions {
			parts := strings.Split(line, " ")

			instructions = append(instructions, []int{conv.ToInt(parts[1]), conv.ToInt(parts[3]), conv.ToInt(parts[5])})
			continue
		}

		for i := 1; i < len(line); i += 4 {
			if line[i] == ' ' || lines[j+1] == "" {
				continue
			}

			stackN := (i)/4 + 1

			if _, ok := stacks[stackN]; !ok {
				stacks[stackN] = []rune{}
			}

			stacks[stackN] = array.Unshift(stacks[stackN], rune(line[i]))
		}
	}

	return stacks, instructions
}
