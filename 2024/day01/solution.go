package day01

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc25/internal/array"
	"janisvepris/aoc25/internal/conv"
	"janisvepris/aoc25/internal/files"
	"janisvepris/aoc25/internal/maths"
)

var (
	lines        []string
	list1, list2 []int
)

func Setup() {
	lines = files.ReadFile("2024/day01/input.txt")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, "   ")
		list1 = append(list1, conv.StrToInt(parts[0]))
		list2 = append(list2, conv.StrToInt(parts[1]))
	}
}

func Part1() {
	distances := []int{}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range list1 {
		distances = append(distances, maths.AbsInt(list1[i]-list2[i]))
	}

	sumFunc := func(value, carry int) int {
		return value + carry
	}

	result := array.Reduce(distances, sumFunc, 0)
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	occurences := map[int]int{}

	for _, value := range list1 {
		count := 0
		if _, ok := occurences[value]; !ok {
			for _, value2 := range list2 {
				if value == value2 {
					count++
				}
			}

			occurences[value] = count
		}

		result += value * occurences[value]
	}
	fmt.Printf("Part 2: %d\n", result)
}
