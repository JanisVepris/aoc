package day15

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	orderedmap "janisvepris/aoc/internal/ordered_map"
)

type Box struct {
	*orderedmap.OrderedMap[string, int]
}

func NewBox() *Box {
	return &Box{orderedmap.NewOrderedMap[string, int]()}
}

var (
	lines []string
	steps []string
)

func Setup() {
	lines = files.ReadFile("2023/day15/input.txt")
	steps = strings.Split(lines[0], ",")
}

func Part1() {
	result := 0
	for _, step := range steps {
		result += hash(step)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	boxes := make([]*Box, 256)
	for i := range boxes {
		boxes[i] = NewBox()
	}

	for _, step := range steps {
		boxNum, label, action := parseStep(step)

		if action == "-" {
			boxes[boxNum].Delete(label)
			continue
		}

		boxes[boxNum].Set(label, conv.StrToInt(action))
	}

	for boxNum, box := range boxes {
		if box.Len() == 0 {
			continue
		}

		result += calcFocusingPower(boxNum, box)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func hash(step string) (result int) {
	for _, byte := range []byte(step) {
		result += int(byte)
		result *= 17
		result %= 256
	}

	return
}

func calcFocusingPower(boxNum int, box *Box) int {
	result := 0

	for lensNum, label := range box.Keys() {
		focalLength, _ := box.Get(label)

		result += (boxNum + 1) * (lensNum + 1) * focalLength
	}

	return result
}

func parseStep(step string) (int, string, string) {
	if strings.Contains(step, "=") {
		parts := strings.Split(step, "=")

		return hash(parts[0]), parts[0], parts[1]
	}

	label := step[0 : len(step)-1]

	return hash(label), label, "-"
}
