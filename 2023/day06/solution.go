package day06

import (
	"fmt"
	"strings"

	"janisvepris/aoc25/internal/array"
	"janisvepris/aoc25/internal/conv"
	"janisvepris/aoc25/internal/files"
)

type Race struct {
	idx      int
	time     int
	distance int
}

func (race *Race) WaysToWin() int {
	ways := 0

	for i := 0; i < race.time; i++ {
		if TravelDistance(i, *race) > race.distance {
			ways++
		}
	}

	return ways
}

var (
	lines []string
	races []Race
)

func Setup() {
	lines = files.ReadFile("2023/day06/input.txt")
	races = BuildRaces(lines)
}

func Part1() {
	result := 1

	array.Each(races, func(i int, race Race) {
		result *= race.WaysToWin()
	})
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	time := ""
	distance := ""

	for _, race := range races {
		time += fmt.Sprint(race.time)
		distance += fmt.Sprint(race.distance)
	}

	bigRace := Race{idx: 0, time: conv.StrToInt(time), distance: conv.StrToInt(distance)}

	result := bigRace.WaysToWin()
	fmt.Printf("Part 2: %d\n", result)
}

func TravelDistance(holdDuration int, race Race) int {
	return holdDuration * (race.time - holdDuration)
}

func BuildRaces(lines []string) (races []Race) {
	filterFn := func(s string) bool {
		return s != ""
	}

	mapFn := func(idx int, s string) int {
		return conv.StrToInt(s)
	}

	times := array.Map(
		array.Shift(
			array.Filter(
				strings.Split(lines[0], " "),
				filterFn,
			),
		),
		mapFn,
	)

	distances := array.Map(
		array.Shift(
			array.Filter(
				strings.Split(lines[1], " "),
				filterFn,
			),
		),
		mapFn,
	)

	for i, time := range times {
		races = append(races, Race{idx: i, time: time, distance: distances[i]})
	}

	return
}
