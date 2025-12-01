package day06

import (
	"fmt"

	"janisvepris/aoc/internal/files"
)

var lines []string

type Point struct {
	x          int
	y          int
	isObstacle bool
	char       rune
	visited    bool
}

type Dir struct {
	x int
	y int
}

var (
	LEFT  = Dir{-1, 0}
	RIGHT = Dir{1, 0}
	UP    = Dir{0, -1}
	DOWN  = Dir{0, 1}
)

func Setup() {
	lines = files.ReadFile("2024/day06/input.txt")
}

func Part1() {
	grid := getGrid(lines)
	guardPos, guardDir := getGuardPosition(grid)

	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1

	steps := 0

	for {
		grid[guardPos.y][guardPos.x].visited = true

		nextX := guardPos.x + guardDir.x
		nextY := guardPos.y + guardDir.y

		if nextX < 0 || nextX > maxX || nextY < 0 || nextY > maxY {
			break
		}

		if grid[nextY][nextX].isObstacle {
			guardDir = turn(guardDir)
			continue
		}

		guardPos.x = nextX
		guardPos.y = nextY

		steps++
	}

	visited := 0

	for _, row := range grid {
		for _, point := range row {
			if point.visited {
				visited++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", visited)
}

func Part2() {
	result := 0
	fmt.Printf("Part 2: %d\n", result)
}

func isLoop(grid [][]Point, guardPos Point, guardDir Dir) bool {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	states := make(map[string]bool)

	getState := func(pos Point, dir Dir) string {
		return fmt.Sprintf("%d,%d,%d,%d", pos.x, pos.y, dir.x, dir.y)
	}

	for {
		state := getState(guardPos, guardDir)

		if _, ok := states[state]; ok {
			return true
		}

		states[state] = true

		nextX := guardPos.x + guardDir.x
		nextY := guardPos.y + guardDir.y

		if nextX < 0 || nextX > maxX || nextY < 0 || nextY > maxY {
			return false
		}

		if grid[nextY][nextX].isObstacle {
			guardDir = turn(guardDir)
			continue
		}

		guardPos.x = nextX
		guardPos.y = nextY
	}
}

func getGrid(lines []string) [][]Point {
	grid := make([][]Point, len(lines))

	for y, line := range lines {
		for x, char := range line {
			grid[y] = append(grid[y], Point{x, y, char == '#', char, false})
		}
	}

	return grid
}

func getGuardPosition(grid [][]Point) (Point, Dir) {
	var position Point
	var direction Dir

	for y, row := range grid {
		for x, point := range row {

			if point.char == '.' || point.char == '#' {
				continue
			}

			position = Point{x, y, false, '@', false}

			switch point.char {
			case '^':
				direction = UP
			case 'v':
				direction = DOWN
			case '<':
				direction = LEFT
			case '>':
				direction = RIGHT
			}
		}
	}

	return position, direction
}

func turn(currentDir Dir) Dir {
	var newDir Dir
	switch currentDir {
	case UP:
		newDir = RIGHT
	case RIGHT:
		newDir = DOWN
	case DOWN:
		newDir = LEFT
	case LEFT:
		newDir = UP
	}

	return newDir
}
