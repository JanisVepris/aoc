package day18

import (
	"fmt"
	"strconv"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

var (
	gw, gh int
	lines  []string
)

var dirs = map[string][2]int{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

type Instruction struct {
	dir   [2]int
	steps int
}

func Setup() {
	lines = files.ReadFile("2023/day18/input.txt")
}

func Part1() {
	instructions := parseInstructions()
	vertices := make([][2]int, len(instructions))

	current := [2]int{0, 0}
	for i, instr := range instructions {
		dir := instr.dir
		current[0] += dir[1] * instr.steps
		current[1] += dir[0] * instr.steps
		vertices[i] = current
	}

	result := calcArea(vertices)

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	instructions := parseHexInstructions()

	vertices := make([][2]int, len(instructions))

	current := [2]int{0, 0}
	for i, instr := range instructions {
		dir := instr.dir
		current[0] += dir[1] * instr.steps
		current[1] += dir[0] * instr.steps
		vertices[i] = current
	}

	result := calcArea(vertices)

	fmt.Printf("Part 2: %d\n", result)
}

func parseInstructions() []*Instruction {
	instructions := make([]*Instruction, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		instructions[i] = &Instruction{
			dir:   dirs[parts[0]],
			steps: conv.StrToInt(parts[1]),
		}
	}

	return instructions
}

func parseHexInstructions() []*Instruction {
	instructions := make([]*Instruction, len(lines))

	for i, line := range lines {
		hex := strings.Fields(line)[2][2:8]

		steps, _ := strconv.ParseInt(hex[0:5], 16, 0)
		var dir [2]int

		switch hex[5] {
		case '0':
			dir = dirs["R"]
		case '1':
			dir = dirs["D"]
		case '2':
			dir = dirs["L"]
		case '3':
			dir = dirs["U"]
		}

		instructions[i] = &Instruction{
			dir:   dir,
			steps: int(steps),
		}
	}

	return instructions
}

func calcArea(vertices [][2]int) int {
	A2 := 0
	B := 0
	nVert := len(vertices)

	for i := range nVert {
		j := (i + 1) % nVert

		A2 += vertices[i][1]*vertices[j][0] - vertices[j][1]*vertices[i][0]

		dx := maths.AbsInt(vertices[j][1] - vertices[i][1])
		dy := maths.AbsInt(vertices[j][0] - vertices[i][0])
		B += dx + dy
	}

	if A2 < 0 {
		A2 = -A2
	}

	return (A2+B)/2 + 1
}
