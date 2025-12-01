package day02

import (
	"fmt"
	"strings"

	"janisvepris/aoc25/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day02/input.txt")
}

func Part1() {
	result := 0

	for _, line := range lines {
		shapes := strings.Split(line, " ")

		switch shapes[1] {
		case "X":
			result += 1
			if shapes[0] == "A" {
				result += 3
			} else if shapes[0] == "B" {
				result += 0
			} else if shapes[0] == "C" {
				result += 6
			}
		case "Y":
			result += 2
			if shapes[0] == "A" {
				result += 6
			} else if shapes[0] == "B" {
				result += 3
			} else if shapes[0] == "C" {
				result += 0
			}
		case "Z":
			result += 3
			if shapes[0] == "A" {
				result += 0
			} else if shapes[0] == "B" {
				result += 6
			} else if shapes[0] == "C" {
				result += 3
			}
		}

	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range lines {
		shapes := strings.Split(line, " ")

		if shapes[1] == "X" {
			switch shapes[0] {
			case "A":
				result += 3
			case "B":
				result += 1
			case "C":
				result += 2
			}
			continue
		}

		if shapes[1] == "Y" {
			result += 3
			switch shapes[0] {
			case "A":
				result += 1
			case "B":
				result += 2
			case "C":
				result += 3
			}
			continue
		}

		if shapes[1] == "Z" {
			result += 6
			switch shapes[0] {
			case "A":
				result += 2
			case "B":
				result += 3
			case "C":
				result += 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
