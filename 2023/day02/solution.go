package day02

import (
	"fmt"
	"strings"

	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

type CubeValue struct {
	color string
	count int
}

type CubeCounts struct {
	red   int
	green int
	blue  int
}

var lines []string

func Setup() {
	lines = files.ReadFile("2023/day02/input.txt")
}

func Part1() {
	maxValues := CubeCounts{red: 12, green: 13, blue: 14}
	result := 0

	for _, line := range lines {
		gameID, counts := parseLine(line)

		if isGamePossible(maxValues, counts) {
			result += gameID
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	result := 0

	for _, line := range lines {
		_, counts := parseLine(line)

		result += getPower(counts)

	}
	fmt.Printf("Part 2: %d\n", result)
}

func isGamePossible(maxValues CubeCounts, cubeCounts CubeCounts) bool {
	if maxValues.red < cubeCounts.red {
		return false
	}

	if maxValues.blue < cubeCounts.blue {
		return false
	}

	if maxValues.green < cubeCounts.green {
		return false
	}

	return true
}

func parseLine(line string) (gameID int, counts CubeCounts) {
	lineParts := strings.Split(line, ":")

	gameID = getGameID(lineParts[0])

	counts = parseGames(lineParts[1])

	return
}

func parseGames(gamesString string) CubeCounts {
	games := strings.Split(gamesString, ";")

	currentCounts := CubeCounts{0, 0, 0}

	for _, game := range games {
		cubeValues := getGameValues(game)

		for _, cubeValue := range cubeValues {
			if cubeValue.color == "red" {
				currentCounts.red = maths.MaxInt(currentCounts.red, cubeValue.count)
			}

			if cubeValue.color == "green" {
				currentCounts.green = maths.MaxInt(currentCounts.green, cubeValue.count)
			}

			if cubeValue.color == "blue" {
				currentCounts.blue = maths.MaxInt(currentCounts.blue, cubeValue.count)
			}
		}
	}

	return currentCounts
}

func getGameValues(gameString string) []CubeValue {
	values := strings.Split(gameString, ", ")

	var cubeValues []CubeValue

	for _, value := range values {
		parts := strings.Split(strings.TrimSpace(value), " ")
		cubeValues = append(cubeValues, CubeValue{color: parts[1], count: conv.StrToInt(parts[0])})
	}

	return cubeValues
}

func getGameID(gameIDString string) int {
	parts := strings.Split(gameIDString, " ")

	return conv.StrToInt(parts[1])
}

func getPower(counts CubeCounts) int {
	return counts.red * counts.green * counts.blue
}
