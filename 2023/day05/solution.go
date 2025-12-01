package day05

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

type SrcDestMap struct {
	items []SrcDestMapItem
}

func (m *SrcDestMap) Add(item SrcDestMapItem) {
	m.items = append(m.items, item)
}

func (m *SrcDestMap) GetDest(source Point) Point {
	for _, item := range m.items {
		if source >= item.srcStart && source <= item.srcEnd {
			return Point(item.destStart + (source - item.srcStart))
		}
	}

	return source
}

type SrcDestMapItem struct {
	srcStart  Point
	srcEnd    Point
	destStart Point
	destEnd   Point
}

type Point int

func NewSrcDestItem(destStart int, srcStart int, rangeLength int) SrcDestMapItem {
	srcEnd := srcStart
	destEnd := destStart

	if rangeLength != 0 {
		srcEnd = srcStart + rangeLength - 1
		destEnd = destStart + rangeLength - 1
	}

	return SrcDestMapItem{
		srcStart:  Point(srcStart),
		srcEnd:    Point(srcEnd),
		destStart: Point(destStart),
		destEnd:   Point(destEnd),
	}
}

func buildSeeds(seedNumbers string) (seeds []Point) {
	for _, seedNumber := range strings.Split(seedNumbers, " ") {
		if seedNumber == "" {
			continue
		}

		seeds = append(
			seeds,
			Point(conv.StrToInt(strings.TrimSpace(seedNumber))),
		)
	}

	return seeds
}

func buildMap(mapLines []string) (newMap SrcDestMap) {
	mapLines = array.Shift(mapLines)

	for _, mapLine := range mapLines {
		numbers := strings.Split(mapLine, " ")
		newMap.Add(NewSrcDestItem(
			conv.StrToInt(numbers[0]),
			conv.StrToInt(numbers[1]),
			conv.StrToInt(numbers[2]),
		))
	}

	return
}

func findMinimumDest(seeds []Point, maps []SrcDestMap) Point {
	minimumPoint := Point(-1)

	for _, point := range seeds {
		for mapIdx := range maps {
			point = maps[mapIdx].GetDest(point)
		}

		if minimumPoint == -1 {
			minimumPoint = point
		}

		if point < minimumPoint {
			minimumPoint = point
		}
	}

	return Point(minimumPoint)
}

var (
	lines []string
	seeds []Point
	maps  []SrcDestMap = []SrcDestMap{}
)

func Setup() {
	lines = files.ReadFile("2023/day05/input.txt")
	seedLine, lines := array.ShiftRet(lines)

	seeds = buildSeeds(strings.Split(seedLine, ":")[1])

	lines = array.Shift(lines)

	block := []string{}

	for i, line := range lines {
		if line == "" {
			newMap := buildMap(block)
			maps = append(maps, newMap)

			block = []string{}

			continue
		}

		block = append(block, line)

		if i == len(lines)-1 {
			newMap := buildMap(block)
			maps = append(maps, newMap)
		}
	}
}

func Part1() {
	result := findMinimumDest(seeds, maps)

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	// TODO: this part is way too slow
	var newSeeds []Point

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1] - 1

		for j := start; j <= end; j++ {
			newSeeds = append(newSeeds, Point(j))
		}
	}

	result := findMinimumDest(newSeeds, maps)
	fmt.Printf("Part 2: %d\n", result)
}
