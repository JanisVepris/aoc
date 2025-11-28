package day08

import (
	"fmt"
	"strings"

	"janisvepris/aoc25/internal/array"
	"janisvepris/aoc25/internal/files"
	"janisvepris/aoc25/internal/maths"
)

var (
	lines        []string
	instructions map[string]Instruction
	locations    []Instruction
	steps        []Step
)

func Setup() {
	lines = files.ReadFile("2023/day08/input.txt")

	stepSeq, lines := array.ShiftRet(lines)
	lines = array.Shift(lines)

	steps = make([]Step, len(stepSeq))

	for i, step := range stepSeq {
		dir := R

		if step == 'L' {
			dir = L
		}

		steps[i] = Step{dir}
	}

	instructions, locations = buildInstructionSet(lines)
}

func Part1() {
	result := 0
	location := instructions["AAA"]

	stepCount := len(steps)

	stepIdx := 0
	for location.Code != "ZZZ" {
		if stepIdx == stepCount {
			stepIdx = 0
		}

		step := steps[stepIdx]

		if step.dir == L {
			location = instructions[location.L]
		} else {
			location = instructions[location.R]
		}

		result++
		stepIdx++
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := make([]int, 0)

	stepCount := len(steps)

	for _, location := range locations {
		stepIdx := 0
		singleResult := 0

		for !strings.HasSuffix(location.Code, "Z") {
			if stepIdx == stepCount {
				stepIdx = 0
			}

			step := steps[stepIdx]

			if step.dir == L {
				location = instructions[location.L]
			} else {
				location = instructions[location.R]
			}

			singleResult++
			stepIdx++
		}

		result = append(result, singleResult)
	}

	lcm := maths.LCM(result...)
	fmt.Printf("Part 2: %d\n", lcm)
}

type Direction int

const (
	L Direction = iota
	R Direction = iota
)

type Step struct {
	dir Direction
}

type Instruction struct {
	Code string
	L    string
	R    string
}

func buildInstructionSet(lines []string) (map[string]Instruction, []Instruction) {
	instructions := make(map[string]Instruction)
	startingLocations := make([]Instruction, 0)

	for _, line := range lines {
		code, L, R := parseInstruction(line)

		instructions[code] = Instruction{code, L, R}

		if strings.HasSuffix(code, "A") {
			startingLocations = append(startingLocations, instructions[code])
		}
	}

	return instructions, startingLocations
}

func parseInstruction(line string) (code string, L string, R string) {
	parts := array.Filter(strings.Split(line, " "), func(s string) bool { return s != "" && s != "=" })

	return parts[0], parts[1][1 : len(parts[1])-1], parts[2][0 : len(parts[2])-1]
}
