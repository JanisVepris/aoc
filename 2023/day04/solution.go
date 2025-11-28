package day04

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc25/internal/conv"
	"janisvepris/aoc25/internal/files"
)

var (
	lines []string
	cards []Card
)

func Part1() {
	result := 0

	for _, card := range cards {
		result += card.points()
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	fmt.Printf("Part 2: %d\n", result)
}

type Card struct {
	winningNumbers []int
	numbers        []int
}

func Setup() {
	lines = files.ReadFile("2023/day04/input.txt")
	parseCards()
}

func (c *Card) winningCount() int {
	count := 0

	for _, x := range c.winningNumbers {
		if slices.Contains(c.numbers, x) {
			count += 1
		}
	}

	return count
}

func (c *Card) points() int {
	winningCount := c.winningCount()

	if winningCount == 0 {
		return 0
	}

	result := 1

	for i := 1; i <= winningCount-1; i++ {
		result *= 2
	}

	return result
}

func parseCards() {
	for _, line := range lines {
		numbers := strings.Split(line, ":")[1]

		winningString, numberString := splitCardNumbers(numbers)

		cards = append(cards, Card{winningNumbers: splitStringToInts(winningString), numbers: splitStringToInts(numberString)})
	}
}

func splitCardNumbers(numbers string) (string, string) {
	x := strings.Split(numbers, "|")

	return strings.TrimSpace(x[0]), strings.TrimSpace(x[1])
}

func splitStringToInts(numbers string) []int {
	var ints []int

	splits := strings.Split(numbers, " ")

	for _, n := range splits {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}

		ints = append(ints, conv.StrToInt(n))
	}

	return ints
}
