package day08

import (
	"fmt"
	"regexp"

	"janisvepris/aoc25/internal/files"
)

type Point struct {
	X, Y int
}

var (
	grid          [][]rune
	width, height int
	antennaRegex                  = regexp.MustCompile(`[A-Za-z0-9]`)
	tata          map[string]bool = map[string]bool{}
)

func Setup() {
	lines := files.ReadFile("2024/day08/input.txt")

	grid = make([][]rune, len(lines))

	for y, line := range lines {
		grid[y] = []rune(line)
	}

	height, width = len(grid), len(grid[0])
}

func Part1() {
	antennas := map[rune][]Point{}

	for y := range grid {
		for x, cell := range grid[y] {
			if !antennaRegex.MatchString(string(cell)) {
				continue
			}

			if _, ok := antennas[cell]; !ok {
				antennas[cell] = []Point{{X: x, Y: y}}
				continue
			}

			antennas[cell] = append(antennas[cell], Point{X: x, Y: y})
		}
	}

	antinodes := map[string]bool{}

	for antenna := range antennas {
		found := findAntinodes1(antennas[antenna])

		for _, antinode := range found {
			antinodes[antinode] = true
		}
	}

	result := len(antinodes)

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	antennas := map[rune][]Point{}

	for y := range grid {
		for x, cell := range grid[y] {
			if !antennaRegex.MatchString(string(cell)) {
				continue
			}

			if _, ok := antennas[cell]; !ok {
				antennas[cell] = []Point{{X: x, Y: y}}
				continue
			}

			antennas[cell] = append(antennas[cell], Point{X: x, Y: y})
		}
	}

	antinodes := map[string]bool{}

	for antenna := range antennas {
		found := findAntinodes2(antennas[antenna])

		for _, antinode := range found {
			antinodes[antinode] = true
			tata[antinode] = true
		}
	}

	result := len(antinodes)

	fmt.Printf("Part 2: %d\n", result)
}

func findAntinodes1(positions []Point) (antinodes []string) {
	for i, p1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			p2 := positions[j]

			dx := p2.X - p1.X
			dy := p2.Y - p1.Y

			nx1 := p1.X - dx
			ny1 := p1.Y - dy

			nx2 := p2.X + dx
			ny2 := p2.Y + dy

			if isInBounds(nx1, ny1) {
				antinodes = append(antinodes, fmt.Sprintf("%d,%d", nx1, ny1))
			}
			if isInBounds(nx2, ny2) {
				antinodes = append(antinodes, fmt.Sprintf("%d,%d", nx2, ny2))
			}
		}
	}

	return antinodes
}

func findAntinodes2(positions []Point) (antinodes []string) {
	for i, p1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			p2 := positions[j]

			dx := p2.X - p1.X
			dy := p2.Y - p1.Y

			// find start of the pattern
			startX, startY := p1.X, p1.Y
			k := 1

			for {
				nx := p1.X - (dx * k)
				ny := p1.Y - (dy * k)

				if !isInBounds(nx, ny) {
					break
				}

				startX, startY = nx, ny
				k++
			}

			// continue the pattern from the start
			k = 0
			for {
				nx := startX + (dx * k)
				ny := startY + (dy * k)

				if !isInBounds(nx, ny) {
					break
				}

				antinodes = append(antinodes, fmt.Sprintf("%d,%d", nx, ny))
				k++
			}
		}
	}

	return antinodes
}

func isInBounds(x, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}
