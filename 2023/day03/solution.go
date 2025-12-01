package day03

import (
	"fmt"
	"regexp"
	"unicode"

	"janisvepris/aoc/internal/array"
	"janisvepris/aoc/internal/conv"
	"janisvepris/aoc/internal/files"
	"janisvepris/aoc/internal/maths"
)

type Range struct {
	start int
	end   int
}

func (a *Range) overlap(b AdjacencyRange) bool {
	return b.from <= a.end && a.start <= b.until
}

type AdjacencyRange struct {
	from  int
	until int
}

type Gear struct {
	position int
	iLine    int
}

func (g *Gear) adjacencyRange() AdjacencyRange {
	return AdjacencyRange{from: maths.MaxInt(0, g.position-1), until: maths.MinInt(lineLen, g.position+2)}
}

type EnginePart struct {
	value  int
	iStart int
	iEnd   int
	iLine  int
}

func (p *EnginePart) getRange() Range {
	return Range{p.iStart, p.iEnd}
}

func (p *EnginePart) adjacencyRange() AdjacencyRange {
	return AdjacencyRange{from: maths.MaxInt(0, p.iStart-1), until: maths.MinInt(lineLen, p.iEnd+2)}
}

func (p *EnginePart) isValidPart() bool {
	adjacencyRange := p.adjacencyRange()

	part := lines[p.iLine][adjacencyRange.from:adjacencyRange.until]

	if symbolsRe.MatchString(part) {
		return true
	}

	if p.iLine > 0 {
		part = lines[p.iLine-1][adjacencyRange.from:adjacencyRange.until]

		if symbolsRe.MatchString(part) {
			return true
		}
	}

	if p.iLine < lineCount-1 {
		part = lines[p.iLine+1][adjacencyRange.from:adjacencyRange.until]

		if symbolsRe.MatchString(part) {
			return true
		}
	}

	return false
}

func createEnginePart(value, iStart, iEnd, iLine int) EnginePart {
	return EnginePart{value, iStart, iEnd, iLine}
}

var (
	lines       []string
	lineLen     int
	lineCount   int
	engineParts []EnginePart
	gears       []Gear
)

var (
	numbersRe *regexp.Regexp = regexp.MustCompile(`^[0-9]$`)
	symbolsRe *regexp.Regexp = regexp.MustCompile(`[^0-9\.]`)
)

func Setup() {
	lines = files.ReadFile("2023/day03/input.txt")
	lineCount = len(lines)
	lineLen = len(lines[0])
	engineParts, gears = findParts()
}

func Part1() {
	result := 0

	for _, enginePart := range engineParts {
		if enginePart.isValidPart() {
			result += enginePart.value
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func Part2() {
	// TODO: Implement Part 2
	result := 0
	fmt.Printf("Part 2: %d\n", result)
}

func findParts() (engineParts []EnginePart, gears []Gear) {
	numberScanning := false
	numberStart := -1
	numberEnd := -1
	number := ""

	for iLine, line := range lines {
		for iChar, char := range line {
			if unicode.IsDigit(char) {

				if !numberScanning {
					numberScanning = true
					numberStart = iChar
				}

				number += string(char)
				numberEnd = iChar

				if iChar == lineLen-1 {
					numberScanning = false
					engineParts = append(engineParts, createEnginePart(conv.StrToInt(number), numberStart, numberEnd, iLine))

					numberStart = -1
					numberEnd = -1
					number = ""
				}
			} else if numberScanning {
				numberScanning = false
				engineParts = append(engineParts, createEnginePart(conv.StrToInt(number), numberStart, numberEnd, iLine))

				numberStart = -1
				numberEnd = -1
				number = ""
			}

			if char == '*' {
				gears = append(gears, Gear{position: iChar, iLine: iLine})
			}

		}
	}

	return
}

func getPartsInLine(iLine int) []EnginePart {
	return array.Filter(
		engineParts,
		func(p EnginePart) bool { return p.iLine == iLine && p.isValidPart() },
	)
}
