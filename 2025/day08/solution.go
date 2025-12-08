package day08

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/dsu"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/heap"
	"janisvepris/aoc/internal/points"
)

var (
	lines       []string
	pointsSlice []points.Point3DF
)

type Pair struct {
	Idx1, Idx2 int
	Dist       float64
}

func Setup() {
	lines = files.ReadFile("2025/day08/input.txt")

	for _, line := range lines {
		parts := strings.Split(line, ",")
		pointsSlice = append(pointsSlice, points.Point3DF{
			X: conv.StrToFloat64(parts[0]),
			Y: conv.StrToFloat64(parts[1]),
			Z: conv.StrToFloat64(parts[2]),
		})
	}
}

func Part1() {
	heap := heap.NewHeap(func(a, b Pair) bool {
		return a.Dist < b.Dist
	})

	// find shortest dists
	for i, p1 := range pointsSlice {
		for j := i + 1; j < len(pointsSlice); j++ {
			ps := &pointsSlice[j]
			dx := ps.X - p1.X
			dy := ps.Y - p1.Y
			dz := ps.Z - p1.Z
			distSquared := dx*dx + dy*dy + dz*dz
			heap.Push(Pair{Idx1: i, Idx2: j, Dist: distSquared})
		}
	}

	connections := 0
	dsu := dsu.NewDSU(len(pointsSlice))
	for connections < 1000 {
		pair, _ := heap.Pop()
		dsu.Union(pair.Idx1, pair.Idx2)
		connections++
	}

	// get three largest components
	max1, max2, max3 := 0, 0, 0

	for _, root := range dsu.GetRoots() {
		size := dsu.GetSize(root)
		if size > max1 {
			max1, max2, max3 = size, max1, max2
		} else if size > max2 {
			max2, max3 = size, max2
		} else if size > max3 {
			max3 = size
		}
	}

	result := max1 * max2 * max3

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	heap := heap.NewHeap(func(a, b Pair) bool {
		return a.Dist < b.Dist
	})

	// find shortest dists
	for i, p1 := range pointsSlice {
		for j := i + 1; j < len(pointsSlice); j++ {
			ps := pointsSlice[j]
			dx := ps.X - p1.X
			dy := ps.Y - p1.Y
			dz := ps.Z - p1.Z
			distSquared := dx*dx + dy*dy + dz*dz
			heap.Push(Pair{Idx1: i, Idx2: j, Dist: distSquared})
		}
	}

	// form the connections
	var lastX1, lastX2 float64
	dsu := dsu.NewDSU(len(pointsSlice))
	for heap.Len() > 0 {
		pair, _ := heap.Pop()
		dsu.Union(pair.Idx1, pair.Idx2)

		// everything is connected
		if dsu.GetComponentCount() == 1 {
			lastX1 = pointsSlice[pair.Idx1].X
			lastX2 = pointsSlice[pair.Idx2].X
			break
		}
	}

	result := int(lastX1 * lastX2)

	fmt.Printf("Part 2: %d\n", result)
}
