package day02

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var ranges []string

func Setup() {
	line := files.ReadFile("2025/day02/input.txt")[0]

	ranges = strings.Split(line, ",")
}

func Part1() {
	result := 0

	for _, r := range ranges {
		IDs := strings.Split(r, "-")
		start := conv.StrToInt(IDs[0])
		end := conv.StrToInt(IDs[1])

		for id := start; id <= end; id++ {
			idStr := conv.ToStr(id)

			if len(idStr)%2 != 0 {
				continue
			}

			firstHalf := idStr[:len(idStr)/2]
			secondHalf := idStr[len(idStr)/2:]

			if firstHalf == secondHalf {
				result += id
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	hasPattern := func(s string) bool {
		sLen := len(s)

		for pLen := 1; pLen <= sLen/2; pLen++ {
			if sLen%pLen != 0 {
				continue
			}

			matches := true
			for i := range sLen {
				if s[i] != s[i%pLen] {
					matches = false
					break
				}
			}

			if matches {
				return matches
			}
		}

		return false
	}

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start := conv.StrToInt(parts[0])
		end := conv.StrToInt(parts[1])

		for id := start; id <= end; id++ {
			if hasPattern(conv.ToStr(id)) {
				result += id
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
