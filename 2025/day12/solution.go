package day12

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2025/day12/input.txt")
}

func Part1() {
	shapeArea := []int{}

	i := 0
	// parse shapes
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		// region lines look like WxH:
		if isRegion(line) {
			break
		}
		// shape header "N:"
		if line[len(line)-1] == ':' {
			idx := conv.StrToInt(line[:len(line)-1])
			for len(shapeArea) <= idx {
				shapeArea = append(shapeArea, 0)
			}
			i++
			for i < len(lines) {
				row := strings.TrimSpace(lines[i])
				if row == "" || row[len(row)-1] == ':' || isRegion(row) {
					break
				}
				for j := 0; j < len(row); j++ {
					if row[j] == '#' {
						shapeArea[idx]++
					}
				}
				i++
			}
			continue
		}
		i++
	}

	ok := 0

	// parse regions
	for ; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || !isRegion(line) {
			continue
		}

		// parse "WxH: ..."
		p := strings.IndexByte(line, ':')
		wh := line[:p]
		counts := strings.TrimSpace(line[p+1:])

		x := strings.IndexByte(wh, 'x')
		w := conv.StrToInt(wh[:x])
		h := conv.StrToInt(wh[x+1:])

		need := 0
		shape := 0
		n := 0
		inNum := false

		for j := 0; j <= len(counts); j++ {
			if j == len(counts) || counts[j] == ' ' || counts[j] == '\t' {
				if inNum {
					if shape < len(shapeArea) {
						need += n * shapeArea[shape]
					}
					shape++
					n = 0
					inNum = false
				}
				continue
			}
			// digit
			inNum = true
			n = n*10 + int(counts[j]-'0')
		}

		if need <= w*h {
			ok++
		}
	}

	fmt.Printf("Part 1: %d\n", ok)
}

func Part2() {
	fmt.Println("Part 2: N/A")
}

func isRegion(s string) bool {
	x := strings.IndexByte(s, 'x')
	c := strings.IndexByte(s, ':')
	return x > 0 && c > x
}
