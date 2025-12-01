package day01

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

// TODO: redo this horrible solution
var lines []string

func Setup() {
	lines = files.ReadFile("2023/day01/input.txt")
}

func Part1() {
	result := 0

	for _, line := range lines {
		result += getFirstDigit(line, []string{}) + getLastDigit(line, []string{})
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		result += getFirstDigit(line, words) + getLastDigit(line, words)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func getFirstDigit(line string, words []string) int {
	for _, word := range words {
		if strings.HasPrefix(line, word) {
			return wordToDigit(word) * 10
		}
	}

	value := line[0]

	if '0' < value && value <= '9' {
		return conv.StrToInt(string(line[0]) + "0")
	}

	return getFirstDigit(line[1:], words)
}

func getLastDigit(line string, words []string) int {
	for _, word := range words {
		if strings.HasSuffix(line, word) {
			return wordToDigit(word)
		}
	}

	value := line[len(line)-1]

	if '0' < value && value <= '9' {
		return conv.StrToInt(string(value))
	}

	return getLastDigit(line[:len(line)-1], words)
}

func isDigit(num int) bool {
	return 48 <= num && num <= 57
}

func wordToDigit(word string) int {
	switch word {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
