package day05

import (
	"fmt"
	"slices"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var (
	lines          []string
	invalidUpdates [][]string          = [][]string{}
	rules          map[string][]string = map[string][]string{}
	updates        [][]string          = [][]string{}
)

func Setup() {
	lines = files.ReadFile("2024/day05/input.txt")

	isRules := true

	for _, line := range lines {
		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			parts := strings.Split(line, "|")
			if rule, ok := rules[parts[0]]; ok {
				rules[parts[0]] = append(rule, parts[1])
			} else {
				rules[parts[0]] = []string{parts[1]}
			}
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
}

func Part1() {
	result := 0

	for _, update := range updates {
		if isValid(rules, update) {
			updateLen := len(update)
			result += conv.StrToInt(update[updateLen/2])
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

// TODO: This is wrong, the result is incorrect
func Part2() {
	result := 0

	for _, update := range updates {
		for !isValid(rules, update) {
			update = swapInvalidPages(rules, update)
		}

		result += conv.StrToInt(update[len(update)/2])
	}
	fmt.Printf("Part 2: %d\n", result)
}

func swapInvalidPages(rules map[string][]string, update []string) []string {
	for idx, page := range update {
		pageRules, ok := rules[page]

		if !ok {
			continue
		}

		for _, rule := range pageRules {
			ruleIdx := slices.Index(update, rule)

			if ruleIdx == -1 {
				continue
			}

			if ruleIdx < idx {
				update = swapUpdates(update, idx, ruleIdx)
			}
		}
	}

	return update
}

func isValid(rules map[string][]string, update []string) bool {
	for idx, page := range update {
		pageRules, ok := rules[page]

		if !ok {
			continue
		}

		for _, rule := range pageRules {
			ruleIdx := slices.Index(update, rule)

			if ruleIdx == -1 {
				continue
			}

			if ruleIdx < idx {
				return false
			}
		}
	}

	return true
}

func swapUpdates(update []string, idx1, idx2 int) []string {
	update[idx1], update[idx2] = update[idx2], update[idx1]
	return update
}
