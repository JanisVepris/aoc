//go:generate go run tools/generate_registry.go

package main

import (
	"flag"
	"fmt"

	"janisvepris/aoc/internal/str"
)

func main() {
	var year string
	var day string
	var downloadInput bool
	var createTemplate bool

	flag.StringVar(&year, "y", "", "Year of the challenge")
	flag.StringVar(&day, "d", "00", "Day of the challenge")
	flag.BoolVar(&downloadInput, "dl", false, "Download input data for the specified day and year")
	flag.BoolVar(&createTemplate, "t", false, "Create template solution.go for the specified day and year")
	flag.Parse()

	if year == "" {
		fmt.Println("Year must be specified.")
		flag.Usage()
		return
	}

	day = str.PadLeft(day, 2, "0")

	if day != "00" && createTemplate {
		GenerateSolution(year, day)
	}

	if day != "00" && downloadInput {
		DownloadInput(year, day)
	}

	if createTemplate || downloadInput {
		return
	}

	yearSolutions, foundYear := solutions[year]
	if !foundYear {
		fmt.Printf("No solutions found for year %s\n", year)
		fmt.Println("Available years:")
		if len(solutions) == 0 {
			fmt.Println("  (none - run 'go generate' to build registry)")
		} else {
			for y := range solutions {
				fmt.Printf("  %s\n", y)
			}
		}
		return
	}

	if day == "00" {
		for d, solution := range yearSolutions {
			fmt.Printf("=== %s-%s ===\n", year, d)
			solution()
		}

		return
	}

	solution, foundDay := yearSolutions[day]
	if !foundDay {
		fmt.Printf("No solution found for year %s, day %s\n", year, day)
		fmt.Printf("Available days for %s:\n", year)
		for d := range yearSolutions {
			fmt.Printf("  %s\n", d)
		}
		return
	}

	solution()
}
