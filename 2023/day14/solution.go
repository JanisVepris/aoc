package day14

import (
	"fmt"
	"hash/fnv"
	"slices"
	"strings"

	"janisvepris/aoc25/internal/files"
)

type Node struct {
	hash   uint32
	matrix [][]rune
	next   *Node
}

const (
	CUBE  rune = '#'
	ROCK  rune = 'O'
	SPACE rune = '.'
	UP    int  = 1
	LEFT  int  = 2
	DOWN  int  = 3
	RIGHT int  = 4
)

var lines []string

func Setup() {
	lines = files.ReadFile("2023/day14/input.txt")
}

func Part1() {
	result := 0
	matrix := pushRocks(getMatrix(lines), UP)

	result += calcResult(matrix)
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	matrix := getMatrix(lines)

	cycleCount := 1000000000

	cache := map[uint32]*Node{}

	completedCycles := 0
	loopNode := &Node{}

	for c := 1; c <= cycleCount; c++ {
		hash := hashMatrix(matrix)

		node, cacheHit := cache[hash]

		if cacheHit && node.next != nil {
			matrix = node.next.matrix
			continue
		}

		matrix = pushRocks(matrix, UP)
		matrix = pushRocks(matrix, LEFT)
		matrix = pushRocks(matrix, DOWN)
		matrix = pushRocks(matrix, RIGHT)

		newHash := hashMatrix(matrix)

		completedCycles++

		if nextNode, ok := cache[newHash]; ok && cacheHit {
			node.next = nextNode
			loopNode = node

			break
		} else {
			newHash := hashMatrix(matrix)

			cache[newHash] = &Node{hash: newHash, matrix: matrix}

			if prevNode, ok := cache[hash]; ok {
				prevNode.next = cache[newHash]
			}
		}
	}

	loopSize := 0
	startHash := loopNode.hash

	for {
		loopNode = loopNode.next
		loopSize++

		if loopNode.hash == startHash {
			break
		}
	}
	cycleCount = (cycleCount - completedCycles) % loopSize

	for c := 0; c <= cycleCount; c++ {
		loopNode = loopNode.next
	}

	result := calcResult(loopNode.matrix)

	fmt.Printf("Part 2: %d\n", result)
}

func calcResult(matrix [][]rune) int {
	result := 0

	yLen := len(matrix)

	for y, line := range matrix {
		for _, char := range line {
			if char == ROCK {
				result += yLen - y
			}
		}
	}

	return result
}

func pushRocks(matrix [][]rune, dir int) (newMatrix [][]rune) {
	yLen := len(matrix)
	xLen := len(matrix[0])

	for range yLen {
		newMatrix = append(newMatrix, []rune(strings.Repeat(string(SPACE), xLen)))
	}

	lastRockPos := -1

	switch dir {
	case UP:
		for x := range xLen {
			for y := range yLen {
				elem := matrix[y][x]

				if elem == SPACE {
					continue
				}

				if elem == CUBE {
					lastRockPos = y
					newMatrix[y][x] = CUBE
					continue
				}

				lastRockPos = lastRockPos + 1
				newMatrix[lastRockPos][x] = ROCK
			}
			lastRockPos = -1
		}
	case LEFT:
		for y := range yLen {
			for x := range xLen {
				elem := matrix[y][x]

				if elem == SPACE {
					continue
				}

				if elem == CUBE {
					lastRockPos = x
					newMatrix[y][x] = CUBE
					continue
				}

				lastRockPos = lastRockPos + 1
				newMatrix[y][lastRockPos] = ROCK
			}
			lastRockPos = -1
		}
	case DOWN:
		slices.Reverse(matrix)
		newMatrix = pushRocks(matrix, UP)
		slices.Reverse(newMatrix)
	case RIGHT:
		for i := range matrix {
			slices.Reverse(matrix[i])
		}

		newMatrix = pushRocks(matrix, LEFT)

		for i := range newMatrix {
			slices.Reverse(newMatrix[i])
		}
	}

	return
}

func getMatrix(lines []string) [][]rune {
	result := make([][]rune, 0)

	for _, line := range lines {
		result = append(result, []rune(line))
	}

	return result
}

func hashMatrix(arr [][]rune) uint32 {
	h := fnv.New32a()

	for _, row := range arr {
		for _, r := range row {
			h.Write([]byte(string(r)))
		}
	}

	return h.Sum32()
}
