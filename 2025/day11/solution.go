package day11

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/files"
)

var lines []string

type Device struct {
	outs []string
}

func Setup() {
	lines = files.ReadFile("2025/day11/input.txt")
}

func Part1() {
	result := 0

	devices := parseDevices()
	result = dfs(devices, "you", "out", make(map[string]int, len(devices)))

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	devices := parseDevices()

	part1 := dfs(devices, "svr", "fft", make(map[string]int, len(devices)))
	part2 := dfs(devices, "fft", "dac", make(map[string]int, len(devices)))
	part3 := dfs(devices, "dac", "out", make(map[string]int, len(devices)))

	result = part1 * part2 * part3

	fmt.Printf("Part 2: %d\n", result)
}

func dfs(devices map[string]*Device, start, end string, cache map[string]int) int {
	if start == end {
		return 1
	}

	if start == "out" {
		return 0
	}

	if val, ok := cache[start]; ok {
		return val
	}

	total := 0

	for _, out := range devices[start].outs {
		total += dfs(devices, out, end, cache)
	}

	cache[start] = total

	return total
}

func parseDevices() map[string]*Device {
	devices := make(map[string]*Device)
	for _, line := range lines {
		parts := strings.Split(line, ":")

		id := parts[0]
		outs := strings.Fields(parts[1])

		devices[id] = &Device{outs: outs}
	}

	return devices
}
