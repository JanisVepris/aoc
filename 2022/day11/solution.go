package day11

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

var lines []string

type Monkey struct {
	inspections  int
	items        []int
	condition    int
	trueMonkey   int
	falseMonkey  int
	operandValue int
	useOldValue  bool
	isMultiply   bool
}

func Setup() {
	lines = files.ReadFile("2022/day11/input.txt")
}

func Part1() {
	monkeys := parseMonkeys()

	simulate(monkeys, 20, 3)

	max1, max2 := 0, 0

	for _, monkey := range monkeys {
		if monkey.inspections > max1 {
			max2 = max1
			max1 = monkey.inspections
		} else if monkey.inspections > max2 {
			max2 = monkey.inspections
		}
	}

	monkeyBusiness := max1 * max2

	fmt.Printf("Part 1: %d\n", monkeyBusiness)
}

func Part2() {
	monkeys := parseMonkeys()

	simulate(monkeys, 10000, 1)

	max1, max2 := 0, 0
	for _, monkey := range monkeys {
		if monkey.inspections > max1 {
			max2 = max1
			max1 = monkey.inspections
		} else if monkey.inspections > max2 {
			max2 = monkey.inspections
		}
	}

	monkeyBusiness := max1 * max2

	fmt.Printf("Part 2: %d\n", monkeyBusiness)
}

func parseMonkeys() map[int]*Monkey {
	monkeys := map[int]*Monkey{}
	var condition, trueMonkey, falseMonkey int
	var operation, operand string
	var items []int
	monkeyIdx := 0

	for i, line := range lines {
		line = strings.TrimSpace(line)

		switch true {
		case line == "":
			operandValue := 0
			useOldValue := false
			if operand == "old" {
				useOldValue = true
			} else {
				operandValue = conv.StrToInt(operand)
			}
			isMultiply := operation == "*"

			monkeys[monkeyIdx] = &Monkey{
				inspections:  0,
				items:        items,
				condition:    condition,
				trueMonkey:   trueMonkey,
				falseMonkey:  falseMonkey,
				operandValue: operandValue,
				useOldValue:  useOldValue,
				isMultiply:   isMultiply,
			}
			items = []int{}
			monkeyIdx++
			continue
		case strings.HasPrefix(line, "Starting"):
			parts := strings.Split(line, ": ")
			itemStrings := strings.Split(parts[1], ", ")
			for _, itemStr := range itemStrings {
				items = append(items, conv.StrToInt(itemStr))
			}
		case strings.HasPrefix(line, "Operation"):
			parts := strings.Split(line, " ")
			operation = parts[len(parts)-2]
			operand = parts[len(parts)-1]
		case strings.HasPrefix(line, "Test"):
			parts := strings.Split(line, " ")
			condition = conv.StrToInt(parts[len(parts)-1])
		case strings.HasPrefix(line, "If true"):
			parts := strings.Split(line, " ")
			trueMonkey = conv.StrToInt(parts[len(parts)-1])
		case strings.HasPrefix(line, "If false"):
			parts := strings.Split(line, " ")
			falseMonkey = conv.StrToInt(parts[len(parts)-1])
		}

		if i == len(lines)-1 {
			operandValue := 0
			useOldValue := false
			if operand == "old" {
				useOldValue = true
			} else {
				operandValue = conv.StrToInt(operand)
			}
			isMultiply := operation == "*"

			monkeys[monkeyIdx] = &Monkey{
				inspections:  0,
				items:        items,
				condition:    condition,
				trueMonkey:   trueMonkey,
				falseMonkey:  falseMonkey,
				operandValue: operandValue,
				useOldValue:  useOldValue,
				isMultiply:   isMultiply,
			}
		}
	}

	return monkeys
}

func simulate(monkeys map[int]*Monkey, rounds, worryDivider int) {
	divisors := make([]int, len(monkeys))

	for i := 0; i < len(monkeys); i++ {
		divisors[i] = monkeys[i].condition
	}

	lcm := maths.LCM(divisors...)

	for i := 0; i < len(monkeys); i++ {
		capacity := len(monkeys[i].items) * 20
		if capacity < 100 {
			capacity = 100
		}
		items := make([]int, len(monkeys[i].items), capacity)
		copy(items, monkeys[i].items)
		monkeys[i].items = items
	}

	for range rounds {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]

			if len(monkey.items) == 0 {
				continue
			}

			itemsToProcess := monkey.items
			for _, worryLevel := range itemsToProcess {
				monkey.inspections++
				modifier := 0
				if monkey.useOldValue {
					modifier = worryLevel
				} else {
					modifier = monkey.operandValue
				}

				if monkey.isMultiply {
					worryLevel = worryLevel * modifier
				} else {
					worryLevel = worryLevel + modifier
				}

				worryLevel %= lcm

				worryLevel /= worryDivider

				if worryLevel%monkey.condition == 0 {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, worryLevel)
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, worryLevel)
				}
			}
			monkey.items = monkey.items[:0]
		}
	}
}
