package day09

import (
	"fmt"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var (
	lines   []string
	offsets = map[string][2]int{
		"U": {0, 1}, // [x, y]
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}
)

func Setup() {
	lines = files.ReadFile("2022/day09/input.txt")
}

func Part1() {
	visited := map[int64]bool{}
	hX, hY, tX, tY := 0, 0, 0, 0
	visited[coordsToKey(tX, tY)] = true

	for _, line := range lines {
		dir, steps := parseLine(line)
		offset := offsets[dir]

		for range steps {
			hX, hY = hX+offset[0], hY+offset[1]

			// adjacency check
			if (hX >= tX-1 && hX <= tX+1) && (hY >= tY-1 && hY <= tY+1) {
				continue
			}

			dx := hX - tX
			dy := hY - tY

			if dx > 0 {
				tX++
			} else if dx < 0 {
				tX--
			}

			if dy > 0 {
				tY++
			} else if dy < 0 {
				tY--
			}

			visited[coordsToKey(tX, tY)] = true
		}
	}

	fmt.Printf("Part 1: %d\n", len(visited))
}

func Part2() {
	visited := map[int64]bool{}
	knots := make([][2]int, 10) // [][x, y]

	visited[coordsToKey(knots[9][0], knots[9][1])] = true

	for _, line := range lines {
		dir, steps := parseLine(line)
		offset := offsets[dir]

		for range steps {
			// move head
			knots[0][0] += offset[0]
			knots[0][1] += offset[1]

			// move rest of the snake
			for i := 1; i < len(knots); i++ {
				hX, hY := knots[i-1][0], knots[i-1][1]
				tX, tY := knots[i][0], knots[i][1]

				if (hX >= tX-1 && hX <= tX+1) && (hY >= tY-1 && hY <= tY+1) {
					continue
				}

				dx := hX - tX
				dy := hY - tY

				if dx > 0 {
					tX++
				} else if dx < 0 {
					tX--
				}

				if dy > 0 {
					tY++
				} else if dy < 0 {
					tY--
				}

				knots[i][0], knots[i][1] = tX, tY
			}

			visited[coordsToKey(knots[9][0], knots[9][1])] = true
		}

	}

	fmt.Printf("Part 2: %d\n", len(visited))
}

func parseLine(line string) (string, int) {
	dir := line[0:1]
	stepCount := conv.StrToInt(line[2:])

	return dir, stepCount
}

func coordsToKey(x, y int) int64 {
	return int64(x)<<32 | int64(uint32(y))
}
