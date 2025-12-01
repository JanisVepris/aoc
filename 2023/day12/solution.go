package day12

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

const (
	WORKING rune = '.'
	DAMAGED rune = '#'
	UNKNOWN rune = '?'
)

var (
	lines   []string
	configs []string
	seqs    [][]int
)

func Setup() {
	lines = files.ReadFile("2023/day12/input.txt")
	configs, seqs = parseLines(lines)
}

func Part1() {
	result := 0
	cache := make(map[string]int)

	for i := 0; i < len(configs); i++ {
		result += calcWays(configs[i], seqs[i], "", cache)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	unfoldFactor := 5
	result := 0
	cache := make(map[string]int)

	for i := 0; i < len(configs); i++ {
		config := configs[i]
		seq := seqs[i]

		for j := 1; j < unfoldFactor; j++ {
			config = config + "?" + configs[i]
			seq = append(seq, seqs[i]...)
		}

		result += calcWays(config, seq, "", cache)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func calcWays(currentConfig string, seq []int, res string, cache map[string]int) int {
	if len(seq) == 0 && strings.Contains(currentConfig, string(DAMAGED)) {
		return 0
	}

	if len(currentConfig) == 0 && len(seq) != 0 {
		return 0
	}

	if len(seq) == 0 {
		return 1
	}

	key := getKey(currentConfig, seq)

	if val, ok := cache[key]; ok {
		return val
	}

	result := 0

	runes := []rune(currentConfig)
	currentValue := runes[0]

	if currentValue == WORKING {
		result += calcWays(currentConfig[1:], slices.Clone(seq), res+string(currentValue), cache)
	}

	if currentValue == DAMAGED {
		newSeq := slices.Clone(seq)
		newSeq[0]--

		sequenceEnded := false
		nextConfig := currentConfig[1:]
		hasNextElement := len(runes) > 1

		if newSeq[0] == 0 {
			newSeq = newSeq[1:]
			sequenceEnded = true
		}

		// next must be either UNKNOWN or WORKING, impossible to progress
		if sequenceEnded && hasNextElement && runes[1] == DAMAGED {
			cache[key] = result
			return result
		}

		// sequence hasn't ended, next must not be WORKING
		if !sequenceEnded && hasNextElement && runes[1] == WORKING {
			cache[key] = result
			return result
		}

		// sequence ended next must be WORKING, replace UNKNOWN with WORKING
		if sequenceEnded && hasNextElement && runes[1] == UNKNOWN {
			nextConfig = string(WORKING) + nextConfig[1:]
		}

		// sequence hasn't ended, next must be DAMAGED, replace UNKNOWN with DAMAGED
		if !sequenceEnded && hasNextElement && runes[1] == UNKNOWN {
			nextConfig = string(DAMAGED) + nextConfig[1:]
		}

		result = calcWays(nextConfig, newSeq, res+string(currentValue), cache)
	}

	if currentValue == UNKNOWN {
		result += calcWays(string(DAMAGED)+currentConfig[1:], slices.Clone(seq), res, cache)
		result += calcWays(string(WORKING)+currentConfig[1:], slices.Clone(seq), res, cache)
	}

	cache[key] = result
	return result
}

func parseLines(lines []string) (configs []string, sequences [][]int) {
	for _, line := range lines {
		parts := strings.Fields(line)

		config := parts[0]
		seq := parts[1]

		configs = append(configs, config)

		numbers := array.Map(
			strings.Split(seq, ","), func(idx int, s string) int { return conv.StrToInt(s) })

		sequences = append(sequences, numbers)
	}

	return
}

func getKey(config string, seq []int) string {
	return config + conv.SliceToStr(seq, ",")
}
