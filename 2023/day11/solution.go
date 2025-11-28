package day11

import (
	"fmt"
	"slices"

	"janisvepris/aoc25/internal/files"
)

const (
	GALAXY rune = '#'
	SPACE  rune = '.'
)

type Element struct {
	x, y int
}

type Pair struct {
	e1, e2 Element
}

func (p Pair) MinDistance(galaxyRows, galaxyCols []int, expansionFactor int) int {
	var x1, x2, y1, y2 int

	if p.e1.x < p.e2.x {
		x1, x2 = p.e1.x, p.e2.x
	} else {
		x1, x2 = p.e2.x, p.e1.x
	}

	if p.e1.y < p.e2.y {
		y1, y2 = p.e1.y, p.e2.y
	} else {
		y1, y2 = p.e2.y, p.e1.y
	}

	emptyColsInBetween, emptyRowsInBetween := 0, 0

	for x := x1 + 1; x < x2; x++ {
		if !slices.Contains(galaxyCols, x) {
			emptyColsInBetween++
		}
	}

	for y := y1 + 1; y < y2; y++ {
		if !slices.Contains(galaxyRows, y) {
			emptyRowsInBetween++
		}
	}

	x2 += emptyColsInBetween * (expansionFactor - 1)
	y2 += emptyRowsInBetween * (expansionFactor - 1)

	return (x2 - x1) + (y2 - y1)
}

type Galaxies []Element

var (
	lines                  []string
	pairs                  []Pair
	galaxyCols, galaxyRows []int
)

func Setup() {
	lines = files.ReadFile("2023/day11/input.txt")
	var galaxies Galaxies
	galaxies, galaxyRows, galaxyCols = parseElements(lines)
	pairs = getPairs(galaxies)
}

func Part1() {
	result := 0
	for _, pair := range pairs {
		result += pair.MinDistance(galaxyRows, galaxyCols, 2)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0
	for _, pair := range pairs {
		result += pair.MinDistance(galaxyRows, galaxyCols, 1000000)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func parseElements(lines []string) (galaxies []Element, galaxyRows, galaxyCols []int) {
	for y, line := range lines {
		for x, char := range line {
			if char == GALAXY {
				galaxies = append(galaxies, Element{x, y})

				if !slices.Contains(galaxyRows, y) {
					galaxyRows = append(galaxyRows, y)
				}

				if !slices.Contains(galaxyCols, x) {
					galaxyCols = append(galaxyCols, x)
				}
			}
		}
	}

	return
}

func getPairs(galaxies Galaxies) (pairs []Pair) {
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, Pair{galaxies[i], galaxies[j]})
		}
	}

	return
}
