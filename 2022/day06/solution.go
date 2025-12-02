package day06

import (
	"fmt"

	"janisvepris/aoc/internal/files"
)

var line string

func Setup() {
	line = files.ReadFile("2022/day06/input.txt")[0]
}

func Part1() {
	result := 0

	stream := []rune(line)

	interval := 4

	for i := range stream {
		window := stream[i : i+interval]
		unique := make(map[rune]bool)
		for _, char := range window {
			unique[char] = true
		}
		if len(unique) == interval {
			result = i + interval
			break
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	stream := []rune(line)

	interval := 14

	for i := range stream {
		window := stream[i : i+interval]
		unique := make(map[rune]bool)
		for _, char := range window {
			unique[char] = true
		}
		if len(unique) == interval {
			result = i + interval
			break
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
