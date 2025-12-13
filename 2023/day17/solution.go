package day17

import (
	"fmt"

	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/heap"
)

var (
	lines []string
	grid  [][]int
)

func Setup() {
	lines = files.ReadFile("2023/day17/input.txt")
}

func Part1() {
	parseGrid()
	lastNode := dijkstra(1, 3)
	fmt.Printf("Part 1: %d\n", lastNode.cost)
}

func Part2() {
	parseGrid()
	lastNode := dijkstra(4, 10)
	fmt.Printf("Part 2: %d\n", lastNode.cost)
}

func parseGrid() {
	grid = make([][]int, len(lines))

	for row, line := range lines {
		grid[row] = make([]int, len(line))
		for col, ch := range line {
			grid[row][col] = int(ch - '0')
		}
	}
}

type Node struct {
	row, col int
	cost     int
	prev     *Node
	dir      dir
}

func (n *Node) Key() string {
	return fmt.Sprintf("%d_%d_%v", n.row, n.col, n.dir)
}

type dir [2]int

var (
	down  = dir{1, 0}
	up    = dir{-1, 0}
	left  = dir{0, -1}
	right = dir{0, 1}
	turns = map[dir][2]dir{
		down:  {left, right},
		up:    {right, left},
		left:  {up, down},
		right: {down, up},
	}
)

func dijkstra(minLength, maxLength int) *Node {
	queue := heap.NewHeap(func(a, b *Node) bool {
		return a.cost < b.cost
	})

	visited := make(map[string]*Node)

	queue.Push(&Node{row: 0, col: 0, cost: 0, dir: right})
	queue.Push(&Node{row: 0, col: 0, cost: 0, dir: down})

	for queue.Len() > 0 {
		current, _ := queue.Pop()

		if _, ok := visited[current.Key()]; ok {
			continue
		}
		visited[current.Key()] = current

		if current.row == len(grid)-1 && current.col == len(grid[0])-1 {
			return current
		}

		nDirs := turns[current.dir]

		// turn and traverse in both directions
		for _, nDir := range nDirs {
			nCost := current.cost

			// accumulate cost
			for step := 1; step < minLength; step++ {
				nRow := current.row + nDir[0]*step
				nCol := current.col + nDir[1]*step

				if nRow < 0 || nRow >= len(grid) || nCol < 0 || nCol >= len(grid[0]) {
					continue
				}
				nCost += grid[nRow][nCol]
			}

			for step := minLength; step <= maxLength; step++ {
				nRow := current.row + nDir[0]*step
				nCol := current.col + nDir[1]*step

				if nRow < 0 || nRow >= len(grid) || nCol < 0 || nCol >= len(grid[0]) {
					continue
				}

				nCost += grid[nRow][nCol]
				queue.Push(&Node{row: nRow, col: nCol, cost: nCost, prev: current, dir: nDir})
			}
		}
	}

	return nil
}
