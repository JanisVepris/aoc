package day03

import (
	"fmt"

	"janisvepris/aoc25/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day03/input.txt")
}

func Part1() {
	result := 0

	for _, line := range lines {
		mid := len(line) / 2
		com1 := line[:mid]
		com2 := line[mid:]

		com1Map := make(map[rune]bool)

		for _, char := range com1 {
			com1Map[char] = true
		}

		for _, char := range com2 {
			if com1Map[char] {
				if char >= 'a' && char <= 'z' {
					result += int(char-'a') + 1
				}

				if char >= 'A' && char <= 'Z' {
					result += int(char-'A') + 27
				}

				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for i := 0; i < len(lines); i += 3 {
		ruck1 := lines[i]
		ruck2 := lines[i+1]
		ruck3 := lines[i+2]

		itemMap := make(map[rune]int)

		for _, char := range ruck1 {
			itemMap[char] = 1
		}

		for _, char := range ruck2 {
			if _, ok := itemMap[char]; ok {
				itemMap[char] = 2
			}
		}

		for _, char := range ruck3 {
			if _, ok := itemMap[char]; ok && itemMap[char] == 2 {
				if char >= 'a' && char <= 'z' {
					result += int(char-'a') + 1
				}

				if char >= 'A' && char <= 'Z' {
					result += int(char-'A') + 27
				}

				break
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
