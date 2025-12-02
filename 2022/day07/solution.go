package day07

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
)

var lines []string

func Setup() {
	lines = files.ReadFile("2022/day07/input.txt")
}

func Part1() {
	result := 0
	sizes := calculatePathSizes()

	for _, v := range sizes {
		if v <= 100000 {
			result += v
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	sizes := calculatePathSizes()

	totalSpace := 70000000
	neededSpace := 30000000
	usedSpace := sizes["root"]

	minToDelete := neededSpace - (totalSpace - usedSpace)

	result := sizes["root"]

	for _, size := range sizes {
		if size >= minToDelete && size < result {
			result = size
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func calculatePathSizes() (sizes map[string]int) {
	sizes = make(map[string]int)
	path := []string{}
	pathStrings := []string{} // store current and parent paths as strings here
	for i := 0; i < len(lines); i++ {
		switch lines[i] {
		case "$ cd /":
			path = []string{"root"}
			pathStrings = []string{"root"}
			continue
		case "$ cd ..":
			if len(path) > 0 {
				path = path[:len(path)-1]
				pathStrings = pathStrings[:len(pathStrings)-1]
			}
			continue
		case "$ ls":
			dirSize := 0
			for j := i + 1; j < len(lines); j++ {
				if lines[j][0] != '$' && lines[j][:3] != "dir" {
					fileSize := conv.ToInt(strings.Split(lines[j], " ")[0])
					dirSize += fileSize

					for _, pathStr := range pathStrings {
						sizes[pathStr] += fileSize
					}
				}

				if lines[j][0] == '$' || j == len(lines)-1 {
					i = j - 1 // skip these lines in outer loop
					break
				}
			}
		}

		if strings.HasPrefix(lines[i], "$ cd ") {
			dirName := lines[i][5:]
			path = append(path, dirName)

			if len(pathStrings) == 0 {
				pathStrings = append(pathStrings, dirName)
			} else {
				newPath := pathStrings[len(pathStrings)-1] + "/" + dirName
				pathStrings = append(pathStrings, newPath)
			}
		}
	}
	return
}
